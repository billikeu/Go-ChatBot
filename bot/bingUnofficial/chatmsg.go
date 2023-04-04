package bingunofficial

type ChatMsg struct {
	MsgId string
	Req   *ReqMsg
	Res   *ResMsg
}

type ReqMsg struct {
	Content string
}

type ResMsg struct {
	Content string
}

func (msg *ChatMsg) ID() string {
	return msg.MsgId
}
