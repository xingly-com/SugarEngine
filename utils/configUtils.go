package utils

import (
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

func GetAppAccessToken(appId, clientSecret string) (string, error) {
	post, err := http.Post("https://bots.qq.com/app/getAppAccessToken", "application/json",
		strings.NewReader("{\"appId\": \""+appId+"\",\"clientSecret\": \""+clientSecret+"\"}"))
	if err != nil {
		return "", err
	}
	defer post.Body.Close()

	res, err := io.ReadAll(post.Body)
	if err != nil {
		return "", err
	}
	return gjson.Get(string(res), "access_token").String(), nil
}

func BotGet() {

}
