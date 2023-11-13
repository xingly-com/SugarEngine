package main

import (
	"SugarEngine/utils"
	"log"
)

func main() {
	log.Println("Hello SugarEngine is starting...")

	accessToken, _ := utils.GetAppAccessToken("102071307", "IbibHkzzkHYaNvEI")
	log.Println(accessToken)
}
