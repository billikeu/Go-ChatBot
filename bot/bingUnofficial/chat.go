package bingunofficial

import (
	"context"
	"fmt"
	"time"

	"github.com/billikeu/Go-ChatBot/bot/params"
	"github.com/billikeu/Go-EdgeGPT/edgegpt"
)

type BingChatUnofficial struct {
	client  *edgegpt.ChatBot
	requst  *Request
	bconfig *BingConfig
	busy    chan struct{}
}

func NewBingChatUnofficial(bconfig *BingConfig) *BingChatUnofficial {
	chat := &BingChatUnofficial{
		bconfig: bconfig,
		requst:  NewRequest(),
		busy:    make(chan struct{}, 1),
	}
	return chat
}

func (chat *BingChatUnofficial) Init() error {
	chat.client = edgegpt.NewChatBot(chat.bconfig.CookiePath, chat.bconfig.Cookies, chat.bconfig.Proxy)
	err := chat.client.Init()
	if err != nil {
		return err
	}
	return nil
}

func (chat *BingChatUnofficial) SetProxy(proxy string) error {
	if chat.bconfig.Proxy == proxy {
		return nil
	}
	chat.client = edgegpt.NewChatBot(chat.bconfig.CookiePath, chat.bconfig.Cookies, proxy)
	err := chat.client.Init()
	if err != nil {
		return err
	}
	chat.bconfig.Proxy = proxy
	return nil
}

func (chat *BingChatUnofficial) SetBaseURL(baseURL string) {

}

func (chat *BingChatUnofficial) SetSystemMsg(content string) {

}

func (chat *BingChatUnofficial) Ask(ctx context.Context, prompt string, callback func(params *params.CallParams, err error)) (err error) {
	// timeout: 3 second for busy
	wait := true
	for wait {
		select {
		case <-time.After(time.Second * 3):
			return fmt.Errorf("busy")
		case chat.busy <- struct{}{}:
			wait = false
		}
	}
	defer func() {
		<-chat.busy
	}()

	msgId, parentId := chat.requst.PutUserMsg(&ChatMsg{
		MsgId: edgegpt.GetUuidV4(),
		Req: &ReqMsg{
			Content: prompt,
		},
	})
	defer func() {
		if err != nil {
			chat.requst.PopMsg()
		}
	}()
	err = chat.client.Ask(prompt, edgegpt.Creative, func(answer *edgegpt.Answer) {
		if answer == nil {
			return
		}
		if callback != nil {
			callback(&params.CallParams{
				MsgId:    msgId,
				ParentId: parentId,
				Text:     answer.Text(),
				Model:    params.NewBingUnofficial,
				Done:     answer.IsDone(),
			}, nil)
		}
		if answer.IsDone() {
			chat.requst.SetRes(&ResMsg{
				Content: answer.Text(),
			})
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func (chat *BingChatUnofficial) RefreshProxy(proxy string) error {
	if chat.bconfig.Proxy == proxy {
		return nil
	}
	chat.client = edgegpt.NewChatBot(chat.bconfig.CookiePath, chat.bconfig.Cookies, proxy)
	err := chat.client.Init()
	if err != nil {
		return err
	}
	chat.bconfig.Proxy = proxy
	return nil
}

func (chat *BingChatUnofficial) RefreshSecretKey(secretKey string) error {
	return nil
}
