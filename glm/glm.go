package glm

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"webchat/config"
)

var BASEURL string = config.LoadConfig().GLMBackend

// ChatGLMResponseBody 请求体
type ChatGLMResponseBody struct {
	Response string     `json:"response"`
	History  [][]string `json:"history"`
	Status   int        `json:"status"`
	Time     string     `json:"time"`
}

// ChatGLMRequestBody 响应体
type ChatGLMRequestBody struct {
	Prompt  string     `json:"prompt"`
	History [][]string `json:"history"`
}

func Completions_with_history(msg string, history_stack *History_stack) (string, error) {

	// 校验是否已超出轮数
	history_stack.check_rounds()
	/*
		err := history_stack.check_rounds()
		if err != nil {
			return "", err
		}
	*/
	requestBody := ChatGLMRequestBody{
		Prompt:  msg,
		History: *history_stack.History,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request glm json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	glmResponseBody := &ChatGLMResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, glmResponseBody)
	if err != nil {
		return "", err
	}

	*history_stack.History = glmResponseBody.History

	if len(glmResponseBody.History) > 0 {
		for _, v := range glmResponseBody.History {
			log.Printf("glm response history: 问%s ,答%s \n", v[0], v[1])
			break
		}
	}

	reply := glmResponseBody.Response
	log.Printf("glm response text: %s \n", reply)
	return reply, nil
}

func Completions(sender string, msg string) (string, error) {
	// 读取存储的历史记录
	history, err := GetHistoryStack(sender)
	if err != nil {
		return "", err
	}
	reply, err := Completions_with_history(msg, history)
	return reply, err
}
