package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/qingconglaixueit/abing_logger"
)

// Configuration 项目配置
type Configuration struct {
	BaseURL string `json:"base_url"`
	// gpt apikey
	ApiKey string `json:"api_key"`
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`
	// 会话超时时间
	SessionTimeout time.Duration `json:"session_timeout"`
	// GPT请求最大字符数
	MaxTokens uint `json:"max_tokens"`
	// GPT模型
	Model string `json:"model"`
	// 热度
	Temperature float64 `json:"temperature"`
	// 回复前缀
	ReplyPrefix string `json:"reply_prefix"`
	// 清空会话口令
	SessionClearToken string `json:"session_clear_token"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 给配置赋默认值
		config = &Configuration{
			BaseURL:           "https://api.openai.com/v1/",
			AutoPass:          false,
			SessionTimeout:    600,
			MaxTokens:         1024,
			Model:             "gpt-3.5-turbo",
			Temperature:       0.5,
			SessionClearToken: "会话清空",
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				log.Fatalf("open config err: %v", err)
				return
			}
			defer f.Close()
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				log.Fatalf("decode config err: %v", err)
				return
			}
		}
		// 有环境变量使用环境变量
		BaseURL := os.Getenv("BASE_URL")
		ApiKey := os.Getenv("APIKEY")

		if BaseURL != "" {
			config.BaseURL = BaseURL
		}
		if ApiKey != "" {
			config.ApiKey = ApiKey
		}

	})
	if config.ApiKey == "" {
		abing_logger.SugarLogger.Error("config err: api key required")
	}

	return config
}
