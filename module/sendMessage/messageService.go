package sendMessage

import (
	"SugarEngine/common"
	"SugarEngine/utils"
	"log"
	"strconv"
	"time"
)

func SendNormalMessage(sendType int, openId string, content string, msgType int) {
	var route string
	switch sendType {
	case 0:
		route = "/v2/users/" + openId + "/messages"
	case 1:
		route = "/v2/groups/" + openId + "/messages"
	case 2:
		route = "/channels/" + openId + "/messages"
	case 3:
		route = "/dms/" + openId + "/messages"
	}

	payload := "{\"content\":\"" + content + "\"," +
		"\"msg_type\":" + strconv.Itoa(msgType) + "," +
		"\"timestamp\":" + strconv.FormatInt(time.Now().Unix(), 10) + "}"

	res, err := utils.BotBasePost(common.BaseUrl+route, payload)
	if err != nil {
		ModuleLogError(err.Error())
	}
	log.Println(common.BaseUrl+route, payload)
	ModuleLogInfo(string(res))
}
