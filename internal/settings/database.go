package settings

import (
	"github.com/spf13/viper"
	"golangOnlineJudge/internal/globals"
	"golangOnlineJudge/internal/globals/database"
)

func InitDatabase() (err error) {
	var log = globals.GetLogger()
	if viper.GetBool("system.UseMysql") {
		err = database.InitMysqlClient()
		if err != nil {
			log.Errorln("mysql初始化错误:", err)
			return
		}
	}
	return
}
