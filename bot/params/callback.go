package params

type CallParams struct {
	ConversationId string
	MsgId          string // message id: chatgptuno id,
	ParentId       string
	Chunk          string
	Text           string
	Done           bool
	Model          string
	ChunkIndex     int
}

// create params for ask callback
/*
callback(NewCallParams(msgId, parentId, chunk, text, false), err)
*/
func NewCallParams(msgId, parentId, chunk, text, model string, done bool, chunkIndex int) *CallParams {
	p := &CallParams{
		MsgId:      msgId,
		ParentId:   parentId,
		Chunk:      chunk,
		Text:       text,
		Done:       done,
		Model:      model,
		ChunkIndex: chunkIndex,
	}
	return p
}

func (p *CallParams) SetConversationId(convId string) {
	p.ConversationId = convId
}
