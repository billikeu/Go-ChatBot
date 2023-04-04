package main

import (
	"context"
	"log"

	"github.com/billikeu/Go-ChatBot/bot/params"

	"github.com/billikeu/Go-ChatBot/bot"
)

var callback = func(params *params.CallParams, err error) {
	if params == nil {
		return
	}
	if err != nil {
		log.Println(params.MsgId, err)
	}
	if params.Done {
		log.Println("answer: ", params.Text)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mybot := bot.NewBot(&bot.Config{
		Proxy: "", // socks5://10.0.0.13:3126 , http://10.0.0.13:3127
		ChatGPT: bot.ChatGPTConf{
			SecretKey: "your secret key", // your secret key
		},
	})
	var coversationId string = "rq1p21s32as138zj7f9qrjv4b"
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
				ConversationId:    coversationId,
				Prompt:            prompt,
				Callback:          callback,
				ChatEngine:        params.ChatGPT,
				SystemRoleMessage: sysMessage,
			},
		)
		if err != nil {
			panic(err)
		}
	}
	log.Println("done")
}
