package bot

import (
	"context"
	"errors"

	"github.com/billikeu/Go-ChatBot/bot/chatgpt"
	"github.com/billikeu/Go-ChatBot/bot/params"
)

type Conversation interface {
	Init() error
	SetProxy(proxy string) error
	SetBaseURL(baseURL string)
	SetSystemMsg(content string)
	Ask(ctx context.Context, prompt string, callback func(params *params.CallParams, err error)) (err error)
	RefreshProxy(proxy string) error
	RefreshSecretKey(secretKey string) error
}

type Bot struct {
	chatgptConversations map[string]Conversation
	config               *Config
}

func NewBot(conf *Config) *Bot {
	bot := &Bot{
		chatgptConversations: make(map[string]Conversation, 0),
		config:               conf,
	}
	return bot
}

func (bot *Bot) GetConversation(askParams *params.AskParams) (Conversation, error) {
	if len(askParams.ConversationId) < 16 {
		return nil, errors.New("session must more 16 letters")
	}
	if askParams.BotType == params.BotTypeChatGPT {
		conversation := bot.chatgptConversations[askParams.ConversationId]
		if conversation != nil {
			return conversation, nil
		}
		conversation = chatgpt.NewChatGPTConversion(bot.config.ChatGPT.SecretKey)
		err := conversation.SetProxy(bot.config.Proxy)
		if err != nil {
			return nil, err
		}
		conversation.SetBaseURL(bot.config.ChatGPT.BaseURL)
		err = conversation.Init()
		if err != nil {
			return nil, err
		}
		if askParams.SystemRoleMessage != "" {
			conversation.SetSystemMsg(askParams.SystemRoleMessage)
		}
		bot.chatgptConversations[askParams.ConversationId] = conversation
		return conversation, nil
	} else {
		return nil, errors.New("unimplemented interface")
	}
}

func (bot *Bot) Ask(ctx context.Context, askParams *params.AskParams) error {
	conversation, err := bot.GetConversation(askParams)
	if err != nil {
		return err

	}
	if askParams.RefreshProxy {
		if err = conversation.RefreshProxy(askParams.Proxy); err != nil {
			return err
		}
	}
	if askParams.RefreshSecretKey {
		if err = conversation.RefreshSecretKey(askParams.SecretKey); err != nil {
			return err
		}
	}
	err = conversation.Ask(ctx, askParams.Prompt, askParams.Callback)
	if err != nil {
		return err
	}
	return nil
}
