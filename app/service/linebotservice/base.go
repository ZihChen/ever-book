package linebotservice

import (
	"ever-book/app/global/structs"
	"ever-book/app/model"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
	"os"
	"sync"
)

type Interface interface {
	// GetClient 取得 Line-bot 實例
	GetClient() *linebot.Client
	// ShowIsNeedDateAndRemarkOptionTemplate 是否需要日期、備註選項視窗
	ShowIsNeedDateAndRemarkOptionTemplate() *linebot.TemplateMessage
	// ShowBalanceDateOptionTemplate 填寫指定日期視窗
	ShowBalanceDateOptionTemplate() *linebot.TemplateMessage
	// ShowBalanceItemOptionTemplate 選擇消費項目選項視窗
	ShowBalanceItemOptionTemplate() *linebot.TemplateMessage
	// ShowBalancePaymentOptionTemplate 選擇消費方式選項視窗
	ShowBalancePaymentOptionTemplate() *linebot.TemplateMessage
	// ShowBalanceRemarkOptionTemplate 是否填寫備註選項視窗
	ShowBalanceRemarkOptionTemplate() *linebot.TemplateMessage
	// ShowContinueOrDiscardOptionTemplate 繼續填寫或是捨棄選項視窗
	ShowContinueOrDiscardOptionTemplate() *linebot.TemplateMessage
	// ShowCancelOrNotOptionTemplate 確認是否刪除選項視窗
	ShowCancelOrNotOptionTemplate() *linebot.TemplateMessage
	// ShowTmpBalanceFlexMessage 顯示編輯當下的收支紀錄訊息
	ShowTmpBalanceFlexMessage(msgTitle string, tmpBalanceObj structs.TmpBalanceObj) *linebot.FlexMessage
	// ShowCancelBalanceFlexMessage 顯示即將刪除的收支紀錄訊息
	ShowCancelBalanceFlexMessage(msgTitle string, tmpBalanceObj structs.TmpBalanceObj) *linebot.FlexMessage
	// ShowMonthOptionTemplate 顯示含當月及過去兩個月的月份選項視窗
	ShowMonthOptionTemplate() *linebot.TemplateMessage
	// ShowSummaryFlexMessage 顯示收支統計訊息
	ShowSummaryFlexMessage(msgTitle string, summaryFlexMsg structs.SummaryFlexMsg) *linebot.FlexMessage
	// ShowUserGroupOption 選擇要查看其他人帳本還是綁定成員選項
	ShowUserGroupOption() *linebot.TemplateMessage
	// ShowUserListOption 顯示可選擇綁定的所有成員選項視窗
	ShowUserListOption(userList []structs.UserObj) *linebot.TemplateMessage
	// ShowMemberListOption 顯示已綁定的使用者選項視窗
	ShowMemberListOption(users []model.User) *linebot.TemplateMessage
	// ShowMembersBalanceMonthOption 選擇觀看成員帳本指定月份選項視窗
	ShowMembersBalanceMonthOption(memberID string) (template *linebot.TemplateMessage)
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
