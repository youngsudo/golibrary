package getTime

import (
	"fmt"
	"time"
)

func GetTime() string {
	now := time.Now()      //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func GetTimeStamp() int64 {
	return time.Now().Unix() //时间戳 s
}

// 时间戳转时间
// time.Unix(record.TradingTime, 0).Format("2006-01-02 15:04:05")
