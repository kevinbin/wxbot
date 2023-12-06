package gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/qingconglaixueit/abing_logger"
	"github.com/qingconglaixueit/wechatbot/config"
)

// const BASEURL = "https://api.openai.com/v1/"

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceItem           `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}

type ChoiceItem struct {
	Index        int         `json:"index"`
	Message      MessageItem `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type MessageItem struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model string `json:"model"`
	// message 是一个数组，数组中的每个元素都是一个对象，对象中有两个属性，role 和 content，role 的值是 user，content 的值是用户输入的文本。

	Message          []MessageItem `json:"messages"`
	MaxTokens        uint          `json:"max_tokens"`
	Temperature      float64       `json:"temperature"`
	TopP             int           `json:"top_p"`
	FrequencyPenalty int           `json:"frequency_penalty"`
	PresencePenalty  int           `json:"presence_penalty"`
}

// Completions gpt文本模型回复
// curl https://api.openai.com/v1/chat/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer your chatGPT key"
// -d '{"model": "gpt-3.5-turbo-1106", "messages": [{"role": "user", "content": "Who won the world series in 2020?"}], "temperature": 0, "max_tokens": 7}'
func Completions(msg string) (string, error) {
	cfg := config.LoadConfig()
	requestBody := ChatGPTRequestBody{
		Model:            cfg.Model,
		Message:          []MessageItem{{Role: "user", Content: msg}},
		MaxTokens:        cfg.MaxTokens,
		Temperature:      cfg.Temperature,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	base_url := config.LoadConfig().BaseURL
	abing_logger.SugarLogger.Info(fmt.Sprintf("gpt request: %v", string(requestData)))
	req, err := http.NewRequest("POST", base_url+"chat/completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.LoadConfig().ApiKey

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{Timeout: 30 * time.Second}
	// abing_logger.SugarLogger.Info(req.URL)
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return "", errors.New(fmt.Sprintf("gpt api status code is %d ,details:  %v ", response.StatusCode, string(body)))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	abing_logger.SugarLogger.Info(fmt.Sprintf("gpt response: %v", string(body)))

	gptResponseBody := &ChatGPTResponseBody{}
	// log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		reply = gptResponseBody.Choices[0].Message.Content
	}
	abing_logger.SugarLogger.Info(fmt.Sprintf("gpt response: %s ", reply))
	return reply, nil
}
