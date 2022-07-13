// Package dao
/*
@Coding : utf-8
@time : 2022/7/12 16:49
@Author : yizhigopher
@Software : GoLand
*/
package dao

import (
	"golang-online-judge/internal/globals/database"
	"golang-online-judge/internal/models/mysqlModel"
)

type Submits struct {
	mysqlModel.Submits
}

func (s *Submits) Get() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(s).Take(s).Error
}

func (s *Submits) Add() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Create(&s).Error
}

func (s *Submits) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := s.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&s).Updates(args).Error
}

func (s *Submits) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := s.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&s).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (s *Submits) GetAll() ([]Submits, error) {
	mysqlManager := database.GetMysqlClient()
	submits := []Submits{}
	return submits, mysqlManager.Model(&s).Where(s).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Find(&submits).Error
}
