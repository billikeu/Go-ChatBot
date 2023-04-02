package params

type AskParams struct {
	ConversationId    string
	Prompt            string
	BotType           string
	SystemRoleMessage string
	Callback          func(params *CallParams, err error)
	Proxy             string // optional , If RefreshProxy is true, the proxy value will be used to refresh the proxy
	SecretKey         string // optional , If RefreshSecretKey is true, the proxy value will be used to refresh the SecretKey
	RefreshProxy      bool   // optional
	RefreshSecretKey  bool   // optional
}
