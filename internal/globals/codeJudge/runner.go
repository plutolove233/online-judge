// Package codeJudge
/*
@Coding : utf-8
@time : 2022/7/4 17:24
@Author : yizhigopher
@Software : GoLand
*/
package codeJudge

import (
	"bytes"
	"fmt"
	"golangOnlineJudge/internal/services"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

func calcMemory(mem string) uint64 {
	storage := 0
	for i := len(mem) - 1; i >= 0; i-- {
		if mem[i] >= '0' && mem[i] <= '9' {
			storage = storage*10 + int(mem[i])
		}
	}
	return uint64(storage)
}

func runTaskList(pid int) (string, error) {
	cmd := exec.Command("cmd", "/c", "tasklist|findstr", strconv.Itoa(pid))

	stdOut := bytes.Buffer{}
	stdErr := bytes.Buffer{}
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil {
		return stdErr.String(), err
	}
	res := stdOut.String()
	data := strings.Fields(res)
	return data[len(data)-2], nil
}

type RunnerParser struct {
	Status  int    //程序运行结果状态 1: pass	0: Failed
	Message string //程序运行信息
}

func (runner *RunnerParser) CodeJudge(path string, problemID string, submitID string) {
	passCount := 0
	var lock sync.Mutex
	var err1 error
	var allocate uint64
	var cost time.Duration

	problem := services.ProblemsService{}
	if err1 = problem.Get(); err1 != nil {
		runner.Status = 0
		runner.Message = err1.Error()
		return
	}

	codename := fmt.Sprintf("%s/%s.cpp", path, submitID)
	outPATH := fmt.Sprintf("%s/%s", path, submitID)
	cmd := exec.Command("gcc", codename, "-o", outPATH)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err1 = cmd.Run(); err1 != nil {
		runner.Status = 5 // 返回编译错误信息
		runner.Message = fmt.Sprintf("%s:%s", err1.Error(), stderr.String())
		return
	}

	WA := make(chan int)  //错误答案的channel
	MLE := make(chan int) // 超内存的channel
	AC := make(chan int)  // 答案正确的channel
	RE := make(chan int)  // 运行时错误的channel
	TLE := make(chan int) // 超时间的channel

	for i := 0; i < problem.TestNum; i++ {
		inputPath := fmt.Sprintf("./problems/%s/%d.in", problemID, i)
		outputPath := fmt.Sprintf("./problems/%s/%d.out", problemID, i)
		input, err1 := ioutil.ReadFile(inputPath)
		if err1 != nil {
			runner.Status = 0
			runner.Message = err1.Error()
			return
		}
		output, err1 := ioutil.ReadFile(outputPath)
		if err1 != nil {
			runner.Status = 0
			runner.Message = err1.Error()
			return
		}
		go func(i int, input []byte, output []byte) {
			cmd := exec.Command(outPATH)
			var out, stderr bytes.Buffer
			cmd.Stderr = &stderr
			cmd.Stdout = &out
			stdinPipe, err := cmd.StdinPipe()
			if err != nil {
				log.Fatalln(err)
			}
			io.WriteString(stdinPipe, string(input))

			startTime := time.Now()
			if err = cmd.Start(); err != nil {
				runner.Message = fmt.Sprintf("%s:%s", err.Error(), stderr.String())
				RE <- 1 // 运行时错误
				return
			}
			pid := cmd.Process.Pid
			msg, err2 := runTaskList(pid)
			if err2 != nil {
				runner.Message = msg
				RE <- 1
				return
			}
			allocate = calcMemory(msg)

			io.WriteString(stdinPipe, "~\n") // 输入a表示进程结束

			if err = cmd.Wait(); err != nil {
				runner.Message = fmt.Sprintf("%s:%s", err.Error(), stderr.String())
				RE <- 1
				return
			}
			cost = time.Since(startTime)
			// 答案错误
			if string(output) != out.String() {
				runner.Message = fmt.Sprintf("except is: %s; read is: %s", output, out.String())
				WA <- 1
				return
			}

			// 运行超时
			if cost > time.Millisecond*time.Duration(problem.TimeLimit) {
				TLE <- 1
				return
			}
			//运行超内存
			if (allocate / 1024) > uint64(problem.MemoryLimit) {
				MLE <- 1
				allocate = allocate / 1024
				return
			}

			lock.Lock()
			passCount++
			if passCount == problem.TestNum {
				AC <- 1
			}
			lock.Unlock()

		}(i, input, output)
	}
	runner.Status = 0
	select {
	case <-WA:
		//msg = "答案错误"
		runner.Status = 2
	case <-TLE:
		runner.Status = 3
		runner.Message = "time limit error"
	case <-MLE:
		//msg = "运行超内存"
		runner.Status = 4
		runner.Message = fmt.Sprintf("memory limit error, allocate %dKb", allocate)
	case <-RE:
		runner.Status = 6
	case <-AC:
		runner.Status = 1
	}
	return
}
