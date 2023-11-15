package sendMessage

import "log"

func ModuleLogInfo(text string) {
	log.Println("[module:sendMessage:info] >>> " + text)
}

func ModuleLogError(text string) {
	log.Println("[module:sendMessage:error] >>> " + text)
}

func ModuleLogWarn(text string) {
	log.Println("[module:sendMessage:warn] >>> " + text)
}
