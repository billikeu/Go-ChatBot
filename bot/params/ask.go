package params

type AskParams struct {
	ConversationId    string
	Prompt            string
	BotType           string
	SystemRoleMessage string
	Callback          func(params *CallParams, err error)
}
