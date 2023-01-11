package linebotservice

import (
	"ever-book/app/global"
	"ever-book/app/global/helper"
	"ever-book/app/global/structs"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"strconv"
	"strings"
)

func (s *service) ShowTmpBalanceFlexMessage(msgTitle string, balanceObj structs.ShowBalanceObj) *linebot.FlexMessage {
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
									Text:  balanceObj.Date,
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
									Type:  linebot.FlexComponentTypeText,
									Text:  balanceObj.Type,
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
									Type:  linebot.FlexComponentTypeText,
									Text:  balanceObj.Item,
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
									Type:  linebot.FlexComponentTypeText,
									Text:  balanceObj.Amount,
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
									Text:  "占比",
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#555555",
								},
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  balanceObj.Proportion,
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
									Type:  linebot.FlexComponentTypeText,
									Text:  balanceObj.Payment,
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
									Type:  linebot.FlexComponentTypeText,
									Text:  balanceObj.Remark,
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
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeHorizontal,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "日期",
								Size:   linebot.FlexTextSizeTypeXs,
								Color:  "#555555",
								Weight: linebot.FlexTextWeightTypeBold,
								Align:  linebot.FlexComponentAlignTypeCenter,
								Flex:   linebot.IntPtr(1),
							},
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "花費",
								Size:   linebot.FlexTextSizeTypeXs,
								Color:  "#555555",
								Weight: linebot.FlexTextWeightTypeBold,
								Align:  linebot.FlexComponentAlignTypeCenter,
								Flex:   linebot.IntPtr(1),
							},
						},
					},
					&linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeHorizontal,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "項目",
								Size:   linebot.FlexTextSizeTypeXs,
								Color:  "#555555",
								Weight: linebot.FlexTextWeightTypeBold,
								Align:  linebot.FlexComponentAlignTypeCenter,
								Flex:   linebot.IntPtr(1),
							},
						},
					},
				},
			},
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   "備註",
						Size:   linebot.FlexTextSizeTypeXs,
						Color:  "#555555",
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeCenter,
						Flex:   linebot.IntPtr(1),
					},
				},
			},
		},
	})

	for _, BalanceObj := range summaryFlexMsg.BalanceObjs {
		flexComponents = append(flexComponents, &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeHorizontal,
					Contents: []linebot.FlexComponent{
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  BalanceObj.Date,
									Size:  linebot.FlexTextSizeTypeXs,
									Color: "#555555",
									Align: linebot.FlexComponentAlignTypeCenter,
									Flex:  linebot.IntPtr(1),
								},
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: func() (amount string) {
										amount = strconv.Itoa(BalanceObj.Amount)
										if BalanceObj.Type == global.Income {
											var balanceAmount strings.Builder
											balanceAmount.WriteString("+")
											balanceAmount.WriteString(amount)
											return balanceAmount.String()
										}
										return
									}(),
									Size:  linebot.FlexTextSizeTypeXs,
									Color: "#555555",
									Align: linebot.FlexComponentAlignTypeCenter,
									Flex:  linebot.IntPtr(1),
								},
							},
						},
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  helper.KeyNameConvertToZhTw(BalanceObj.Item),
									Size:  linebot.FlexTextSizeTypeXs,
									Color: "#555555",
									Align: linebot.FlexComponentAlignTypeCenter,
									Flex:  linebot.IntPtr(1),
								},
							},
						},
					},
				},
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeHorizontal,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type: linebot.FlexComponentTypeText,
							Text: func() string {
								if BalanceObj.Remark == "" {
									return "-"
								}
								return BalanceObj.Remark
							}(),
							Size:  linebot.FlexTextSizeTypeXs,
							Color: "#555555",
							Align: linebot.FlexComponentAlignTypeCenter,
							Flex:  linebot.IntPtr(1),
						},
					},
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
					Text: func() string {
						var totalExpenses strings.Builder
						totalExpenses.WriteString("$")
						totalExpenses.WriteString(strconv.Itoa(summaryFlexMsg.TotalExpenses))
						return totalExpenses.String()
					}(),
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
					Text: func() string {
						var totalBalance strings.Builder
						totalBalance.WriteString("$")
						totalBalance.WriteString(strconv.Itoa(summaryFlexMsg.TotalBalance))
						return totalBalance.String()
					}(),
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: func() string {
						var duringDate strings.Builder
						duringDate.WriteString(summaryFlexMsg.StartDate)
						duringDate.WriteString(" - ")
						duringDate.WriteString(summaryFlexMsg.EndDate)
						return duringDate.String()
					}(),
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
