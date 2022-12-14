package linebotservice

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"strconv"
)

func (s *service) ShowTmpBalanceFlexMessage(tmpBalanceObj structs.TmpBalanceObj) *linebot.FlexMessage {
	return linebot.NewFlexMessage("你有一筆紀錄尚未填寫完!", &linebot.BubbleContainer{
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   "你有一筆紀錄尚未填寫完!",
					Weight: linebot.FlexTextWeightTypeBold,
					Color:  "#1DB446",
					Size:   linebot.FlexTextSizeTypeSm,
				},
				&linebot.BoxComponent{
					Type:    linebot.FlexComponentTypeBox,
					Layout:  linebot.FlexBoxLayoutTypeVertical,
					Margin:  linebot.FlexComponentMarginTypeXxl,
					Spacing: linebot.FlexComponentSpacingTypeSm,
					Contents: []linebot.FlexComponent{
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  global.DateZhTw,
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#555555",
								},
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  tmpBalanceObj.Date,
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#111111",
									Align: linebot.FlexComponentAlignTypeEnd,
								},
							},
						},
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  global.TypeZhTw,
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#555555",
								},
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: func() string {
										if tmpBalanceObj.Type == "" {
											return "-"
										}
										return helper.KeyNameConvertToZhTw(tmpBalanceObj.Type)
									}(),
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#111111",
									Align: linebot.FlexComponentAlignTypeEnd,
								},
							},
						},
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  global.ItemZhTw,
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#555555",
								},
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: func() string {
										if tmpBalanceObj.Item == "" {
											return "-"
										}
										return helper.KeyNameConvertToZhTw(tmpBalanceObj.Item)
									}(),
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#111111",
									Align: linebot.FlexComponentAlignTypeEnd,
								},
							},
						},
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  global.AmountZhTw,
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#555555",
								},
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: func() string {
										if tmpBalanceObj.Amount == 0 {
											return "-"
										}
										return strconv.Itoa(tmpBalanceObj.Amount)
									}(),
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#111111",
									Align: linebot.FlexComponentAlignTypeEnd,
								},
							},
						},
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  global.PaymentZhTw,
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#555555",
								},
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: func() string {
										if tmpBalanceObj.Payment == "" {
											return "-"
										}
										return helper.KeyNameConvertToZhTw(tmpBalanceObj.Payment)
									}(),
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#111111",
									Align: linebot.FlexComponentAlignTypeEnd,
								},
							},
						},
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  global.RemarkZhTw,
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#555555",
								},
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: func() string {
										if tmpBalanceObj.Remark == "" {
											return "-"
										}
										return tmpBalanceObj.Remark
									}(),
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#111111",
									Align: linebot.FlexComponentAlignTypeEnd,
								},
							},
						},
					},
				},
			},
		},
	})
}
