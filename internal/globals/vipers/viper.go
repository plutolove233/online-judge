package vipers

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"golangOnlineJudge/internal/utils/logs"
	"sync"
)

var (
	databaseViper     *viper.Viper
	databaseViperOnce sync.Once
)

var log = logs.GetLogger()

func GetDatabaseViper() *viper.Viper {
	databaseViperOnce.Do(func() {
		databaseViper = viper.New()
		databaseViper.SetConfigName("database")
		databaseViper.AddConfigPath("./configs")
		databaseViper.SetConfigType("yaml")

		err := databaseViper.ReadInConfig()
		if err != nil {
			log.Errorf("Fatal error config file: %s \n", err)
			panic(err)
		}

		databaseViper.WatchConfig()

		databaseViper.OnConfigChange(func(e fsnotify.Event) {
			log.Warnf("Config file:%s Op:%s\n", e.Name, e.Op)
		})
	})
	return databaseViper
}
