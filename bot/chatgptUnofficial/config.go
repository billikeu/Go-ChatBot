package chatgptunofficial

import "github.com/billikeu/go-chatgpt/chatgptuno"

type Config struct {
	Proxy       string
	AccessToken string
	ChatGPTUno  *chatgptuno.ChatGPTUnoConfig
}
