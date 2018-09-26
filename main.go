package restday

import (
	"strconv"
	"time"
)

var (
	restdays   = make(map[string]string) // 休息日
	restdayMax = make(map[string]int)    // 休息日天数
)

const REST = "休息日"

func init() {

	// 合并休息日
	for k, v := range restdays_2018 {
		restdays[k] = v
	}

	// 合并休息日天数
	for k, v := range restdayMax_2018 {
		restdayMax[k] = v
	}
}

// 获取是否是工作日
func Is_WorkDay(date string) string {
	res, has := restdays[date]
	if has {
		return res
	}
	return "工作日"
}

// 获取月份天数
func GetMonthDay(month string) (days int) {
	date, _ := time.Parse("2006-01", month)
	days, _ = strconv.Atoi(date.AddDate(0, 1, -1).Format("02"))
	return
}

// 获取月份休息日天数
func GetMonthRestDay(date string) int {
	res, has := restdayMax[date]
	if has {
		return res
	}
	return 0
}

// 获取月份工作日天数
func GetMonthWorkDay(date string) int {
	return GetMonthDay(date) - GetMonthRestDay(date)
}
