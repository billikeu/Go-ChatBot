package bingunofficial

import (
	"log"
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

// return msg_id
func (req *Request) PutSystemMsg(msg *ChatMsg) string {
	req.Lock()
	defer req.Unlock()

	req.chatMsg.Enqueue(msg)
	return msg.ID()
}

// renturn msg_id, parent_id
func (req *Request) PutUserMsg(msg *ChatMsg) (string, string) {
	req.Lock()
	defer req.Unlock()
	var parentId string
	if req.chatMsg.Len() > 0 {
		parentMsg, _ := req.chatMsg.Dequeue()
		parentId = parentMsg.ID()
		req.chatMsg.Enqueue(parentMsg)
	}
	req.chatMsg.Enqueue(msg)
	return msg.ID(), parentId
}

func (req *Request) PopMsg() *ChatMsg {
	req.Lock()
	defer req.Unlock()

	msg, _ := req.chatMsg.Dequeue()
	return msg.(*ChatMsg)

}

func (req *Request) SetRes(response *ResMsg) {
	req.Lock()
	defer req.Unlock()

	if req.chatMsg.Len() > 0 {
		parentMsg, err := req.chatMsg.Dequeue()
		if err != nil {
			log.Println(err)
			return
		}
		msg := parentMsg.(*ChatMsg)
		msg.Res = response
		req.chatMsg.Enqueue(msg)
	}
}
