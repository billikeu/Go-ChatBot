
# Go-ChatBot

`Go-ChatBot` is a Golang-based chatbot framework that allows developers to quickly create their own chatbots with just a few lines of code. It supports various engines including `ChatGPT`, `Bing`, and `Bard`.

## Setup

```
go get -u github.com/billikeu/Go-ChatBot
```

## Demo

```golang

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
				BotType:           params.BotTypeChatGPT,
				SystemRoleMessage: sysMessage,
			},
		)
		if err != nil {
			panic(err)
		}
	}
	log.Println("done")
}

```

Output

```
questions: Give me a joke of no more than 20 characters, it must start with "he"
answer:  He told a joke about pizza.
main.go:43: questions: Is not funny
answer:  I'm sorry you didn't find it funny. Would you like me to try another joke?
questions: 接下来你将用中文回复我,并在括号显示对应的英文
answer:  好的，我会用中文回复你。（I will respond to you in Chinese and display the corresponding English in parentheses.）
questions: 你叫什么名字?
answer:  我叫Go-ChatBot。（My name is Go-ChatBot.）
questions: 你确定你叫这个名字？
answer:  其实我没有名字，是人类给我取的这个名称。（Actually, I don't have a name. It's the humans who gave me this name.）
questions: 你是ChatGPT吗
answer:  不好意思，我不是 ChatGPT，我是 Go-ChatBot。（I'm sorry, I'm not ChatGPT, I'm Go-ChatBot.）
```

## Plan

- [x] ChatGPT
- [ ] Bing
- [ ] Bard


## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=billikeu/Go-ChatBot&type=Date)](https://star-history.com/#billikeu/Go-ChatBot&Date)

## Contributors

This project exists thanks to all the people who contribute.

 <a href="github.com/billikeu/Go-ChatBot/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=billikeu/Go-ChatBot" />
 </a>

## Reference