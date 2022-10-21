package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Message struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func SendMessage(msg AlertManagerMessage, url string) {
	var m Message
	m.MsgType = "text"
	// parse template
	files, err := template.ParseFiles("adapter/wework.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	buf := bytes.NewBufferString("")
	err = files.Execute(buf, msg)
	if err != nil {
		fmt.Println("execute template failed, err:", err)
		return
	}
	m.Text.Content = buf.String()
	// send messages to wework
	jsons, err := json.Marshal(m)
	if err != nil {
		log.Fatal("SendMessage Marshal failed!", err)
		return
	}
	resp := string(jsons)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(resp))
	if err != nil {
		log.Fatal("SendMessage http NewRequest failed!", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		log.Fatal("SendMessage client Do failed", err)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("SendMessage ReadAll Body failed", err)
		return
	}
	log.Printf("SendMessage success！Body: %s", string(body))
}
