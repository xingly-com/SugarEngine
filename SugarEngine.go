package main

import (
	"SugarEngine/common"
	"SugarEngine/module/sendMessage"
	"SugarEngine/utils"
	"log"
)

func main() {
	log.Println("Hello SugarEngine is starting...")

	// 模拟登陆
	common.AccessToken, _ = utils.GetAppAccessToken(common.AppId, common.AppSecret)
	if common.AccessToken == "" {
		log.Println("[main:error] >>> GetAppAccessToken failed")
		return
	}
	log.Println("[main:info] >>> GetAppAccessToken success")
	log.Println("[main:info] >>> ", common.AccessToken)

	// 事件监听

	// 功能区
	sendMessage.SendNormalMessage(1, "820886097", "方糖引擎启动成功", 0)

}
