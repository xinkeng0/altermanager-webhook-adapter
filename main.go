package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"webhook-adapter/adapter"
)

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(responseWriter http.ResponseWriter, request *http.Request) {
	// 解析发送参数
	request.ParseForm()
	formValue := request.FormValue
	all, _ := io.ReadAll(request.Body)
	var tempMap map[string]interface{}
	err := json.Unmarshal(all, &tempMap)

	var alertManagerMessage adapter.AlertManagerMessage
	err = json.Unmarshal(all, &alertManagerMessage)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	// 目前默认只支持企业微信
	adapter.SendMessage(alertManagerMessage, formValue("url"))
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(responseWriter, "{\"message\":\"successful receive alert notification message!\"}")
}

func main() {
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
