package services

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/saharak/etrackings-linebot/app/utils/linebot"
)

type LineBot struct{}

type Messages []linebot.SendingMessage

func (l *LineBot) Client() *linebot.Client {
	client := &http.Client{Timeout: 6 * time.Second}
	bot, err := linebot.New(os.Getenv("LINE_CHANNEL_SECRET"), os.Getenv("LINE_CHANNEL_TOKEN"), linebot.WithHTTPClient(client))
	if err != nil {
		log.Println("[LineBot] err -> ", err)
	}

	return bot
}

func (l *LineBot) Events(request *fasthttp.Request) (events []*linebot.Event, err error) {
	return l.Client().ParseRequest(request)
}

func (l *LineBot) ReplyStrMessage(replyToken, replyMessage string) {
	if _, err := l.Client().ReplyMessage(replyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Println("[LineBot] error -> ", err)
	}
}

func (l *LineBot) ReplyTextMessage(replyToken string, textMessage *linebot.TextMessage) {
	if _, err := l.Client().ReplyMessage(replyToken, textMessage).Do(); err != nil {
		log.Println("[LineBot] error -> ", err)
	}
}
