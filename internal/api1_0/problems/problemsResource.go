// Package problems
/*
@Coding : utf-8
@time : 2022/7/3 22:27
@Author : yizhigopher
@Software : GoLand
*/
package problems

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golangOnlineJudge/internal/globals/codes"
	"golangOnlineJudge/internal/globals/responseParser"
	"golangOnlineJudge/internal/models/ginModels"
	"golangOnlineJudge/internal/services"
	"golangOnlineJudge/internal/utils/snowflake"
	"io/ioutil"
	"os"
)

type ProblemApi struct{}

type UploadNewProblemParser struct {
	Title        string `json:"Title" form:"Title" binding:"required"`
	Content      string `json:"Content" form:"Content" binding:"required"`
	InputLayout  string `json:"InputLayout" form:"InputLayout" binding:"required"`
	OutputLayout string `json:"OutputLayout" form:"OutputLayout" binding:"required"`
	TimeLimit    int    `json:"TimeLimit" form:"TimeLimit" binding:"required"`
	MemoryLimit  int    `json:"MemoryLimit" form:"MemoryLimit" binding:"required"`
}

func (p *ProblemApi) UploadNewProblem(c *gin.Context) {
	parser := UploadNewProblemParser{}
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取请求参数失败", err)
		return
	}

	problemService := services.ProblemsService{}
	problemService.ProblemID = snowflake.GetSnowFlakeID()
	problemService.TimeLimit = parser.TimeLimit
	problemService.MemoryLimit = parser.MemoryLimit
	problemService.InputLayout = parser.InputLayout
	problemService.OutputLayout = parser.OutputLayout
	problemService.Title = parser.Title
	problemService.Content = parser.Content

	err = os.MkdirAll("./problems/"+problemService.ProblemID, os.ModePerm)
	if err != nil {
		responseParser.JsonInternalError(c, "创建问题分区失败", err)
		return
	}

	err = problemService.Add()
	if err != nil {
		responseParser.JsonDBError(c, "问题上传失败", err)
		return
	}

	c.JSON(200, gin.H{
		"code":    codes.OK,
		"message": "题目上传成功",
	})
	return
}

type UploadTestCasesParser struct {
	ProblemID string `json:"ProblemID" form:"ProblemID" binding:"required"`
}

func (p *ProblemApi) UploadProblemTestCases(c *gin.Context) {
	parser := UploadTestCasesParser{}
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取请求参数失败", err)
		return
	}
	problemService := services.ProblemsService{}
	problemService.ProblemID = parser.ProblemID
	if err = problemService.Get(); err != nil {
		if err.Error() == "record not found" {
			responseParser.JsonNotData(c, "该问题未创建", err)
			return
		}
		responseParser.JsonDBError(c, "数据库错误", err)
		return
	}

	form, _ := c.MultipartForm()
	testCases := form.File["testCases"]

	flag := false
	for _, testCase := range testCases {
		flag = true
		dist := fmt.Sprintf("./problems/%s/%s", parser.ProblemID, testCase.Filename)
		err = c.SaveUploadedFile(testCase, dist)
		if err != nil {
			responseParser.JsonInternalError(c, "上传问题测试数据失败", err)
			return
		}
	}
	if !flag {
		err1 := errors.New("get test cases failed")
		responseParser.JsonParameterIllegal(c, "获取测试文件失败", err1)
		return
	}

	err = problemService.Update(map[string]interface{}{
		"TestNum": len(testCases) / 2,
		"Status":  1,
	})
	if err != nil {
		responseParser.JsonDBError(c, "问题状态更新出错", err)
		return
	}

	responseParser.JsonOK(c, "问题测试数据上传成功", nil)
	return
}

type ProblemListResponseParser struct {
	ProblemID     string `json:"ProblemID"`
	Title         string `json:"Title"`
	ProblemStatus bool   `json:"ProblemStatus"`
}

func (p *ProblemApi) GetProblemList(c *gin.Context) {
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}
	user := temp.(ginModels.UserModel)
	problemService := services.ProblemsService{}

	problemList, err := problemService.GetAll()
	if err != nil {
		responseParser.JsonDBError(c, "获取问题列表失败", err)
		return
	}

	data := []ProblemListResponseParser{}
	for _, problem := range problemList {
		submitService := services.SubmitsService{}
		submitService.UserID = user.UserID
		submitService.ProblemID = problem.ProblemID
		submitService.SubmitStatus = "AC"
		err := submitService.Get()
		flag := true
		if err != nil {
			if err.Error() == "record not found" {
				flag = false
			} else {
				responseParser.JsonDBError(c, "数据库错误", err)
				return
			}
		}
		each := ProblemListResponseParser{
			ProblemID:     problem.ProblemID,
			Title:         problem.Title,
			ProblemStatus: flag,
		}
		data = append(data, each)
	}

	responseParser.JsonOK(c, "获取问题列表成功", data)
	return
}

type GetProblemDescriptionParser struct {
	ProblemID string `json:"ProblemID" form:"ProblemID" binding:"required"`
}
type ProblemDescriptionResponseParser struct {
	UploadNewProblemParser
	ExampleIn  string `json:"ExampleIn"`
	ExampleOut string `json:"ExampleOut"`
}

func (p *ProblemApi) GetProblemDescription(c *gin.Context) {
	parser := GetProblemDescriptionParser{}
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取问题id失败", err)
		return
	}

	problemService := services.ProblemsService{}
	problemService.ProblemID = parser.ProblemID
	err = problemService.Get()
	if err != nil {
		if err.Error() == "record not found" {
			responseParser.JsonNotData(c, "问题id对应数据不存在", err)
			return
		}
		responseParser.JsonDBError(c, "数据库错误", err)
		return
	}

	data := ProblemDescriptionResponseParser{}
	data.Title = problemService.Title
	data.TimeLimit = problemService.TimeLimit
	data.Content = problemService.Content
	data.OutputLayout = problemService.OutputLayout
	data.InputLayout = problemService.InputLayout
	data.MemoryLimit = problemService.MemoryLimit

	exampleInPath := fmt.Sprintf("./problems/%s/0.in", problemService.ProblemID)
	exampleOutPath := fmt.Sprintf("./problems/%s/0.out", problemService.ProblemID)
	exampleIn, err := ioutil.ReadFile(exampleInPath)
	if err != nil {
		responseParser.JsonInternalError(c, "读取样例文件失败", err)
		return
	}
	exampleOut, err := ioutil.ReadFile(exampleOutPath)
	if err != nil {
		responseParser.JsonInternalError(c, "读取样例文件失败", err)
		return
	}
	data.ExampleIn = string(exampleIn)
	data.ExampleOut = string(exampleOut)

	responseParser.JsonOK(c, "获取问题描述信息成功", data)
	return
}
