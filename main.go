package main

import (
	"fmt"
	"time"

	"github.com/kuochaoyi/chinese-calendar-golang/calendar"
)

func main() {
	// t := time.Now()
	// 1. ByTimestamp
	// 時間戳
	// c := calendar.ByTimestamp(t.Unix())
	// 2. BySolar
	// 公曆
	// c := calendar.BySolar(year, month, day, hour, minute, second
	// c := calendar.BySolar(2020, 12, 25, 13, 24, 99)
	// 3. ByLunar
	// 農曆(最後一個參數表示是否閏月)
	// c := calendar.ByLunar(year, month, day, hour, minute, second, false)
	d := calendar.ByTimestamp(time.Now().Unix())

	// bytes, _ := c.ToJSON()
	bytes1, _ := d.ToJSON()

	// fmt.Println(string(bytes))
	fmt.Println(string(bytes1))
}
