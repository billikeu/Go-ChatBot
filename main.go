package main

import (
	"context"
	"log"

	bingunofficial "github.com/billikeu/Go-ChatBot/bot/bingUnofficial"
	chatgptunofficial "github.com/billikeu/Go-ChatBot/bot/chatgptUnofficial"
	"github.com/billikeu/Go-ChatBot/bot/params"
	"github.com/billikeu/go-chatgpt/chatgptuno"

	"github.com/billikeu/Go-ChatBot/bot"
)

var coversationId string = "rq1p21s32as138zj7f9qrjv4b"

var callback = func(params *params.CallParams, err error) {
	if params == nil {
		return
	}
	if err != nil {
		log.Println(params.MsgId, err)
	}
	coversationId = params.ConversationId
	if params.Done {
		log.Println("answer: ", coversationId, params.Text)
	}

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mybot := bot.NewBot(&bot.Config{
		// chatgp config
		Proxy: "", // http://10.0.0.13:3127 , socks5://10.0.0.13:3126
		ChatGPT: bot.ChatGPTConf{
			SecretKey: "your secret key", // your secret key
		},
		// bing config
		BingUnofficial: &bingunofficial.BingConfig{
			Proxy:      "", // http://127.0.0.1:10809
			CookiePath: "./data/bingCookie.json",
		},
		ChatGPTUnofficial: &chatgptunofficial.Config{
			Proxy:       "",
			AccessToken: "",
			ChatGPTUno:  &chatgptuno.ChatGPTUnoConfig{},
		},
	})

	var sysMessage = `You are Go-ChatBot, a large language model trained by billikeu. Follow the user's instructions carefully. Respond using markdown.`
	questions := []string{
		`Give me a joke of no more than 20 characters, it must start with "he"`,
		"Is not funny",
		"接下来你将用中文回复我,并在括号显示对应的英文",
		"你叫什么名字?",
		"你确定你叫这个名字？",
		"你是ChatGPT吗",
	}
	for _, prompt := range questions {
		log.Println("questions:", prompt)
		err := mybot.Ask(
			context.Background(),
			&params.AskParams{
				ConversationId:    coversationId, // if chatgpt uno, the first conversion should be nil
				Prompt:            prompt,
				Callback:          callback,
				ChatEngine:        params.ChatGPT, // params.ChatGPT params.NewBingUnofficial
				SystemRoleMessage: sysMessage,
			},
		)
		if err != nil {
			panic(err)
		}
	}
	log.Println("done")
}
