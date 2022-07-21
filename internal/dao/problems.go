// Package dao
/*
@Coding : utf-8
@time : 2022/7/3 11:04
@Author : yizhigopher
@Software : GoLand
*/
package dao

import (
	"golangOnlineJudge/internal/globals/database"
	"golangOnlineJudge/internal/models/mysqlModel"
)

type Problems struct {
	mysqlModel.Problems
}

func (m *Problems) Get() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Take(m).Error
}

func (m *Problems) Add() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Create(&m).Error
}

func (m *Problems) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *Problems) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (m *Problems) GetAll() ([]Problems, error) {
	mysqlManager := database.GetMysqlClient()
	problems := []Problems{}
	return problems, mysqlManager.Model(&m).Where(m).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Find(&problems).Error
}
