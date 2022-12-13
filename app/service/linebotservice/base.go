package linebotservice

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
	"os"
	"sync"
)

type Interface interface {
	GetClient() *linebot.Client
	ShowBalanceDateOptionTemplate() *linebot.TemplateMessage
	ShowBalanceTypeOptionTemplate() *linebot.TemplateMessage
	ShowBalanceItemOptionTemplate() *linebot.TemplateMessage
	ShowBalancePaymentOptionTemplate() *linebot.TemplateMessage
}

type service struct{}

var bot *linebot.Client
var singleton *service
var once sync.Once

func New() Interface {
	once.Do(func() {
		singleton = &service{}
	})
	return singleton
}

func (s *service) GetClient() *linebot.Client {
	var err error
	bot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		linebot.WithHTTPClient(&http.Client{}))
	if err != nil {
		log.Fatalf("[❌ Fatal ❌] New LineBot Error" + err.Error())
	}
	return bot
}
