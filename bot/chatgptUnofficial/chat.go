package chatgptunofficial

import (
	"context"
	"errors"

	"github.com/billikeu/Go-ChatBot/bot/params"
)

type ChatGPTUnofficial struct {
}

func NewChatGPTUnofficial() *ChatGPTUnofficial {
	chat := &ChatGPTUnofficial{}
	return chat
}

func (chat *ChatGPTUnofficial) Init() error {
	return errors.New("interface for implementation")
}

func (chat *ChatGPTUnofficial) SetProxy(proxy string) error {
	return nil
}

func (chat *ChatGPTUnofficial) SetBaseURL(baseURL string) {

}

func (chat *ChatGPTUnofficial) SetSystemMsg(content string) {

}

func (chat *ChatGPTUnofficial) Ask(ctx context.Context, prompt string, callback func(params *params.CallParams, err error)) (err error) {
	return nil
}

func (chat *ChatGPTUnofficial) RefreshProxy(proxy string) error {
	return nil
}

func (chat *ChatGPTUnofficial) RefreshSecretKey(secretKey string) error {
	return nil
}
