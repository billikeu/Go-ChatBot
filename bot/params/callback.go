package params

type CallParams struct {
	MsgId    string
	ParentId string
	Chunk    string
	Text     string
	Done     bool
	Model    string
}

// create params for ask callback
/*
callback(NewCallParams(msgId, parentId, chunk, text, false), err)
*/
func NewCallParams(msgId, parentId, chunk, text, model string, done bool) *CallParams {
	p := &CallParams{
		MsgId:    msgId,
		ParentId: parentId,
		Chunk:    chunk,
		Text:     text,
		Done:     done,
		Model:    model,
	}
	return p
}
