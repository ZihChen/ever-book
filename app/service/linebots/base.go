package linebots

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
	"sync"
)

type Interface interface {
	GetClient() *linebot.Client
}

type service struct{}

var bot *linebot.Client
var singleton *service
var once sync.Once

func NewService() Interface {
	once.Do(func() {
		singleton = &service{}
	})
	return singleton
}

func (s *service) GetClient() *linebot.Client {
	var err error
	bot, err = linebot.New("", "", linebot.WithHTTPClient(&http.Client{}))
	if err != nil {
		log.Fatalf("New LineBot Error" + err.Error())
	}
	return bot
}
