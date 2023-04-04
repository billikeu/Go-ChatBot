package bot

import bingunofficial "github.com/billikeu/Go-ChatBot/bot/bingUnofficial"

type Config struct {
	Proxy                string                     `json:"proxy"` // socks5://22222; http://dddd
	ChatGPT              ChatGPTConf                `json:"chatgpt"`
	BingUnofficialConfig *bingunofficial.BingConfig `json:"bingunofficial_config"`
}

type ChatGPTConf struct {
	SecretKey string `json:"secret_key"`
	BaseURL   string `json:"base_url"`
}
