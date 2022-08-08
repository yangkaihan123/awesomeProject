package main

import (
	"net/http"
	"strings"
)

func main() {
	SendDingMsg("巴林不要挑战我go机器人")
}

func SendDingMsg(msg string) {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=c8a727a84bd8237e6bec23e09a3c0a4c3584594e3a53597d574db60648a2a9ae`
	content := `{"msgtype": "text",
        "text": {"content": "` + "123" + msg + `"}
    }`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		// handle error
	}
	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		// handle error
	}
}
