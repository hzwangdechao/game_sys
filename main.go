package main

import (
	"game_sys/common"
	"game_sys/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// 程序主入口

func main() {
	InitConfig()
	common.InitClient()
	common.InitLog()
	db := common.InitDB()
	defer db.Close()
	logger := common.GetLog()
	logger.Info("hello")

	engine := gin.Default()

	engine = routers.CollectRouter(engine)
	engine.Run(":" + viper.GetString(`server.port`))

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
