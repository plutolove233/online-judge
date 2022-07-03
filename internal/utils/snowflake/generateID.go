// Package snowflake
/*
@Coding : utf-8
@time : 2022/7/3 15:33
@Author : yizhigopher
@Software : GoLand
*/
package snowflake

import (
	"github.com/spf13/viper"
	"golang-online-judge/internal/utils/logs"
	"strconv"
	"sync"
)

var (
	worker     *Worker
	workerOnce sync.Once
	log        = logs.GetLogger()
)

func GetSnowFlakeID() string {
	workerOnce.Do(func() {
		workerID := viper.GetInt64("system.WorkerID")
		dataCenterID := viper.GetInt64("system.DataCenterID")
		worker = NewWorker(workerID, dataCenterID)
		log.Infoln("雪花生成算法初始化完成")
	})
	id, err := worker.NextID()
	if err != nil {
		log.Errorln("生成id错误：", err)
	}
	return strconv.FormatUint(id, 10)
}
