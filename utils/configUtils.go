package utils

import (
	"SugarEngine/common"
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

func BotBasePost(url, payload string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "QQBot "+common.AccessToken)
	req.Header.Set("X-Union-Appid", common.AppId)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
