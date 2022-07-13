// Package submits
/*
@Coding : utf-8
@time : 2022/7/12 17:19
@Author : yizhigopher
@Software : GoLand
*/
package submits

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-online-judge/internal/globals/codeJudge"
	"golang-online-judge/internal/globals/responseParser"
	"golang-online-judge/internal/models/ginModels"
	"golang-online-judge/internal/services"
	"golang-online-judge/internal/utils/code"
	"golang-online-judge/internal/utils/snowflake"
	"time"
)

type SubmitApi struct{}

type SolveProblemsRequestParser struct {
	ProblemID   string `json:"ProblemID" form:"ProblemID" binding:"required"`
	CodeContext string `json:"CodeContext" form:"CodeContext" binding:"required"`
}

func (u *SubmitApi) SubmitProblems(c *gin.Context) {
	var parser SolveProblemsRequestParser
	var err error
	err = c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取请求参数失败", err)
		return
	}

	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}

	user := temp.(ginModels.UserModel)

	submit := services.SubmitsService{}
	submit.ProblemID = parser.ProblemID
	submit.UserID = user.UserID
	submit.SubmitID = snowflake.GetSnowFlakeID()
	submit.SubmitStatus = "Waiting"

	_, err = code.SaveCodeContext(parser.CodeContext, user.UserID, submit.SubmitID)
	if err != nil {
		responseParser.JsonInternalError(c, "代码上传失败", err)
		return
	}

	err = submit.Add()
	if err != nil {
		responseParser.JsonDBError(c, "上传提交记录失败", err)
		return
	}

	responseParser.JsonOK(c, "代码上传成功", submit.SubmitID)
}

type JudgeSubmitRequestParser struct {
	SubmitID string `json:"SubmitID" form:"SubmitID" binding:"required"`
}

func (u *SubmitApi) JudgeSubmit(c *gin.Context) {
	parser := JudgeSubmitRequestParser{}
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取请求参数失败", err)
		return
	}

	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}
	user := temp.(ginModels.UserModel)

	submit := services.SubmitsService{}
	submit.SubmitID = parser.SubmitID
	err = submit.Get()
	if err != nil {
		responseParser.JsonDBError(c, "获取提交信息失败", err)
		return
	}
	path := fmt.Sprintf("./codeArea/%s", user.UserID)
	runner := &codeJudge.RunnerParser{}
	runner.CodeJudge(path, parser.SubmitID)
	if runner.Status == 0 { //处理系统错误
		err = errors.New(runner.Message)
		responseParser.JsonInternalError(c, "判定程序出错", err)
		return
	}

	res := codeJudge.StatusMsgMap[runner.Status]

	err = submit.Update(map[string]interface{}{
		"SubmitStatus": res,
	})
	if err != nil {
		responseParser.JsonDBError(c, "更新提交记录信息失败", err)
		return
	}

	responseParser.JsonOK(c, "评测成功", map[string]interface{}{
		"SubmitStatus": res,
		"Message":      runner.Message,
	})
}

type GetSubmitResponseParser struct {
	SubmitID		string		`json:"SubmitID"`
	ProblemID		string		`json:"ProblemID"`
	SubmitStatus	string		`json:"SubmitStatus"`
	CreateTime		time.Time	`json:"CreateTime"`
}
type GetSubmitRequestParser struct {
	ProblemID	string	`json:"ProblemID" form:"ProblemID" binding:"required"`
}
func (u *SubmitApi) GetSubmitRecord(c *gin.Context) {
	temp,_ := c.Get("user")
	user := temp.(ginModels.UserModel)
	parser := GetSubmitRequestParser{}
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取请求参数失败", nil)
		return
	}
	submitService := services.SubmitsService{}
	submitService.UserID = user.UserID
	submitService.ProblemID = parser.ProblemID
	submitList, err := submitService.GetAll()
	if err != nil {
		if err.Error() == "record not found" {
			responseParser.JsonNotData(c, "提交记录不存在", err)
			return
		}
		responseParser.JsonDBError(c, "数据库错误", err)
		return
	}
	response := []GetSubmitResponseParser{}
	for _, submit := range submitList{
		each := GetSubmitResponseParser{
			ProblemID: submit.ProblemID,
			SubmitID: submit.SubmitID,
			SubmitStatus: submit.SubmitStatus,
			CreateTime: submit.CreateTime,
		}
		response = append(response, each)
	}

	responseParser.JsonOK(c, "获取提交列表成功", response)
	return
}