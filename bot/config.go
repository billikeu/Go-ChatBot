package bot

type Config struct {
	Proxy   string      `json:"proxy"` // socks5://22222; http://dddd
	ChatGPT ChatGPTConf `json:"chatgpt"`
}

type ChatGPTConf struct {
	SecretKey string `json:"secret_key"`
	BaseURL   string `json:"base_url"`
}
