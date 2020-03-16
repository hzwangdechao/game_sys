package controller

import (
	"game_sys/common"
	"game_sys/response"
	"game_sys/utils"
	"github.com/gin-gonic/gin"
)


func GetInfo(c *gin.Context) {
	response.Success(c, gin.H{"text": "hello"}, "ok")
}

func UseUtil(c *gin.Context) {
	log := common.GetLog()
	str := utils.RandomString(10)

	log.Info("qwasdasdae")

	response.Fail(c, gin.H{"su": "qwe"}, str)

}
