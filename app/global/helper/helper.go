package helper

import (
	"ever-book/app/global"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

func ValidateTimeString(timeStr string) (parseDate string, isDate bool) {
	if len(timeStr) != 8 {
		return
	}
	timeArr := []string{}
	timeArr = append(timeArr, timeStr[0:4], timeStr[4:6], timeStr[6:8])
	for k, str := range timeArr {
		parse, err := strconv.Atoi(str)
		if err != nil {
			return
		}
		switch k {
		case 0:
			if parse <= 2099 && parse >= 1970 {
				return
			}
		case 1:
			if parse <= 12 && parse >= 1 {
				return
			}
		case 2:
			if parse <= 31 && parse >= 1 {
				return
			}
		}
	}
	return fmt.Sprintf("%s-%s-%s", timeStr[0:4], timeStr[4:6], timeStr[6:8]), true
}

func ValidateBalanceTypeString(typeStr string) (balanceType string, isCorrect bool) {
	switch typeStr {
	case global.Income:
		return global.Income, true
	case global.Expense:
		return global.Expense, true
	default:
		return
	}
}

func PrintLog(param interface{}) {
	byteData, _ := jsoniter.Marshal(param)
	show := string(byteData)
	fmt.Println("PrintLog:" + show)
	panic(show)
}
