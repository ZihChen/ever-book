package linebotservice

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"strconv"
)

func (s *service) ShowTmpBalanceFlexMessage(msgTitle string, tmpBalanceObj structs.TmpBalanceObj) *linebot.FlexMessage {
	return linebot.NewFlexMessage(msgTitle, &linebot.BubbleContainer{
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   msgTitle,
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

func (s *service) ShowCancelBalanceFlexMessage(msgTitle string, tmpBalanceObj structs.TmpBalanceObj) *linebot.FlexMessage {
	return linebot.NewFlexMessage(msgTitle, &linebot.BubbleContainer{
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   msgTitle,
					Weight: linebot.FlexTextWeightTypeBold,
					Color:  "#E63946",
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

func (s *service) ShowSummaryFlexMessage(msgTitle string, summaryFlexMsg structs.SummaryFlexMsg) *linebot.FlexMessage {
	var flexComponents []linebot.FlexComponent
	flexComponents = append(flexComponents, &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   "編號",
				Size:   linebot.FlexTextSizeTypeSm,
				Color:  "#555555",
				Weight: linebot.FlexTextWeightTypeBold,
				Align:  linebot.FlexComponentAlignTypeCenter,
			},
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   "日期",
				Size:   linebot.FlexTextSizeTypeSm,
				Color:  "#555555",
				Weight: linebot.FlexTextWeightTypeBold,
				Align:  linebot.FlexComponentAlignTypeCenter,
			},
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   "項目",
				Size:   linebot.FlexTextSizeTypeSm,
				Color:  "#555555",
				Weight: linebot.FlexTextWeightTypeBold,
				Align:  linebot.FlexComponentAlignTypeCenter,
			},
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   "花費",
				Size:   linebot.FlexTextSizeTypeSm,
				Color:  "#555555",
				Weight: linebot.FlexTextWeightTypeBold,
				Align:  linebot.FlexComponentAlignTypeCenter,
			},
		},
	})

	for _, BalanceObj := range summaryFlexMsg.BalanceObjs {
		flexComponents = append(flexComponents, &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   strconv.Itoa(BalanceObj.ID),
					Size:   linebot.FlexTextSizeTypeXxs,
					Color:  "#555555",
					Weight: linebot.FlexTextWeightTypeBold,
					Align:  linebot.FlexComponentAlignTypeCenter,
				},
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   BalanceObj.Date,
					Size:   linebot.FlexTextSizeTypeXxs,
					Color:  "#555555",
					Weight: linebot.FlexTextWeightTypeBold,
					Align:  linebot.FlexComponentAlignTypeCenter,
				},
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   helper.KeyNameConvertToZhTw(BalanceObj.Item),
					Size:   linebot.FlexTextSizeTypeXxs,
					Color:  "#555555",
					Weight: linebot.FlexTextWeightTypeBold,
					Align:  linebot.FlexComponentAlignTypeCenter,
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: func() string {
						if BalanceObj.Type == global.Income {
							return "+" + strconv.Itoa(BalanceObj.Amount)
						}
						return strconv.Itoa(BalanceObj.Amount)
					}(),
					Size:   linebot.FlexTextSizeTypeXxs,
					Color:  "#555555",
					Weight: linebot.FlexTextWeightTypeBold,
					Align:  linebot.FlexComponentAlignTypeCenter,
				},
			},
		})
	}

	return linebot.NewFlexMessage(msgTitle, &linebot.BubbleContainer{
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   msgTitle,
					Weight: linebot.FlexTextWeightTypeBold,
					Color:  "#1DB446",
					Size:   linebot.FlexTextSizeTypeSm,
				},
				&linebot.SeparatorComponent{
					Margin: linebot.FlexComponentMarginTypeSm,
				},
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   "已支出",
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeLg,
					Margin: linebot.FlexComponentMarginTypeMd,
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "$" + strconv.Itoa(summaryFlexMsg.TotalExpenses),
				},
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   "結餘",
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeLg,
					Margin: linebot.FlexComponentMarginTypeMd,
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "$" + strconv.Itoa(summaryFlexMsg.TotalBalance),
				},
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   summaryFlexMsg.StartDate + " - " + summaryFlexMsg.EndDate,
					Size:   linebot.FlexTextSizeTypeXs,
					Color:  "#aaaaaa",
					Wrap:   true,
					Margin: linebot.FlexComponentMarginTypeMd,
				},
				&linebot.SeparatorComponent{
					Margin: linebot.FlexComponentMarginTypeSm,
				},
				&linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Margin:   linebot.FlexComponentMarginTypeLg,
					Spacing:  linebot.FlexComponentSpacingTypeSm,
					Contents: flexComponents,
				},
			},
		},
	})
}
