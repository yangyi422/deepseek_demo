package send

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Message struct {
	Role             string `json:"role"`
	Content          string `json:"content"`
	ReasoningContent string `json:"reasoning_content,omitempty"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type ResponseBody struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index        int     `json:"index"`
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

// 请求deepseek的AI
func SendMessage(requestBody RequestBody) (respose ResponseBody) {
	url := "https://api.deepseek.com/chat/completions"
	apiKey := "" // 替换为你的实际API密钥

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("JSON序列化错误: %v\n", err)
		return
	}
	fmt.Printf("请求体: %s\n", string(jsonData))

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("创建请求错误: %v\n", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 创建HTTP客户端并设置超时
	client := &http.Client{
		Timeout: time.Second * 30, // 30秒超时
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求错误: %v\n", err)
		return
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应错误: %v\n", err)
		return
	}

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API返回错误状态码: %d\n", resp.StatusCode)
		fmt.Printf("响应内容: %s\n", string(body))
		return
	}

	var bodyStr = string(body)
	fmt.Printf("响应体: %s\n", bodyStr)

	err = json.Unmarshal(body, &respose)
	if err != nil {
		fmt.Printf("JSON反序列化错误: %v\n", err)
		return
	}

	return
}
