// Package dao
/*
@Coding : utf-8
@time : 2022/7/3 11:04
@Author : yizhigopher
@Software : GoLand
*/
package dao

import (
	"golang-online-judge/internal/globals/database"
	"golang-online-judge/internal/models/mysqlModel"
)

type Users struct {
	mysqlModel.Users
}

func (m *Users) Get() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Take(m).Error
}

func (m *Users) Add() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Create(&m).Error
}

func (m *Users) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *Users) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (m *Users) GetAll(id string) ([]Users, error) {
	mysqlManager := database.GetMysqlClient()
	users := []Users{}
	return users, mysqlManager.Model(&m).Where(map[string]interface{}{
		"UserID":    id,
		"IsDeleted": 0,
	}).Find(&users).Error
}
