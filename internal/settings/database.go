package settings

import (
	"github.com/spf13/viper"
	"golang-online-judge/internal/globals/database"
	"golang-online-judge/internal/utils/logs"
)

func InitDatabase() (err error) {
	var log = logs.GetLogger()
	if viper.GetBool("system.UseMysql") {
		err = database.InitMysqlClient()
		if err != nil {
			log.Errorln("mysql初始化错误:", err)
			return
		}
	}
	return
}
