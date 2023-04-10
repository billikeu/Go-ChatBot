package chatgptunofficial

import (
	"sync"

	"github.com/billikeu/Go-ChatBot/bot/params"
)

type Request struct {
	chatMsg *params.MsgQueue
	sync.RWMutex
}

func NewRequest() *Request {
	req := &Request{
		chatMsg: params.NewMsgQueue(),
	}
	return req
}
