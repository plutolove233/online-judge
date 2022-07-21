// Package code
/*
@Coding : utf-8
@time : 2022/7/4 17:23
@Author : yizhigopher
@Software : GoLand
*/
package code

import (
	"fmt"
	"os"
)

func SaveCodeContext(code string, userID string, submitID string) (string, error) {
	dirName := "./codeArea/" + userID
	path := fmt.Sprintf("%s/%s.cpp", dirName, submitID)
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		return "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	_, err = f.Write([]byte(code))
	if err != nil {
		return "", err
	}
	defer f.Close()
	return path, nil
}
