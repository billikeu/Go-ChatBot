package bot

import (
	bingunofficial "github.com/billikeu/Go-ChatBot/bot/bingUnofficial"
	chatgptunofficial "github.com/billikeu/Go-ChatBot/bot/chatgptUnofficial"
)

type Config struct {
	Proxy             string                     `json:"proxy"` // socks5://22222; http://dddd
	ChatGPT           ChatGPTConf                `json:"chatgpt"`
	BingUnofficial    *bingunofficial.BingConfig `json:"bing_unofficial"`
	ChatGPTUnofficial *chatgptunofficial.Config  `json:"chatgpt_unofficial"`
}

type ChatGPTConf struct {
	SecretKey string `json:"secret_key"`
	BaseURL   string `json:"base_url"`
}
