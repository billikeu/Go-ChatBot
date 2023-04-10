package chatgptunofficial

import (
	"context"

	"github.com/billikeu/Go-ChatBot/bot/params"
	"github.com/billikeu/go-chatgpt/chatgptuno"
)

type ChatGPTUnofficial struct {
	chatuno        *chatgptuno.ChatGPTUnoBot
	cfg            *Config
	conversationId string
	parentId       string
}

func NewChatGPTUnofficial(cfg *Config) *ChatGPTUnofficial {
	chat := &ChatGPTUnofficial{
		cfg:     cfg,
		chatuno: chatgptuno.NewChatGPTUnoBot(cfg.ChatGPTUno),
	}
	return chat
}

func (chat *ChatGPTUnofficial) Init() error {
	return nil
}

func (chat *ChatGPTUnofficial) SetProxy(proxy string) error {
	chat.cfg.ChatGPTUno.Proxy = proxy
	return nil
}

func (chat *ChatGPTUnofficial) SetBaseURL(baseURL string) {
	chat.cfg.ChatGPTUno.BaseUrl = baseURL
}

func (chat *ChatGPTUnofficial) SetSystemMsg(content string) {

}

func (chat *ChatGPTUnofficial) Ask(ctx context.Context, prompt string, callback func(params *params.CallParams, err error)) (err error) {
	chat.chatuno.Ask(prompt, chat.conversationId, chat.parentId, "", 300, func(chatRes *chatgptuno.Response, err error) {
		if callback != nil {
			callback(&params.CallParams{
				ConversationId: chatRes.ConversationID,
				MsgId:          chatRes.Message.ID,
				ParentId:       "",
				Text:           chatRes.Message.Content.Parts[0],
				Done:           chatRes.Message.Metadata.FinishDetails.Stop != "",
			}, err)
		}
	})
	return nil
}

func (chat *ChatGPTUnofficial) RefreshProxy(proxy string) error {
	chat.cfg.ChatGPTUno.Proxy = proxy
	return nil
}

func (chat *ChatGPTUnofficial) RefreshSecretKey(secretKey string) error {
	chat.cfg.ChatGPTUno.AccessToken = secretKey
	return nil
}

func (chat *ChatGPTUnofficial) Engine() string {
	return params.ChatGPTUnofficial
}

func (chat *ChatGPTUnofficial) GetConvId() string {
	return chat.conversationId
}

func (chat *ChatGPTUnofficial) SetConvId(convId string) {
	chat.conversationId = convId
}

func (chat *ChatGPTUnofficial) GetParentId() string {
	return chat.parentId
}

func (chat *ChatGPTUnofficial) SetParentId(parentId string) {
	chat.parentId = parentId
}
