# chinese-calendar-golang

[![Build Status](https://travis-ci.org/Lofanmi/chinese-calendar-golang.svg)](https://travis-ci.org/Lofanmi/chinese-calendar-golang)
[![codecov](https://codecov.io/gh/Lofanmi/chinese-calendar-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/Lofanmi/chinese-calendar-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/Lofanmi/chinese-calendar-golang)](https://goreportcard.com/report/github.com/Lofanmi/chinese-calendar-golang)

** 本專案因為自己需要它, 但原作者沒有採用 GO MODULD 和正體中文

** 所以我 forked from Lofanmi/chinese-calendar-golang 來改.

公曆, 農曆, 干支歷轉換包, 提供精確的日曆轉換.

使用 `Go` 編寫, 覺得好用的給 `Star` 唄~

覺得**好用**, 而不是覺得**有用**. 如果不好用, 歡迎向我提 issue, 我會抽空不斷改進它!

真希望有人打錢打錢打錢給我啊哈哈哈哈!!!

Go 1.7+測試通過, 1.6及以下應該也可以, 不過單元測試跑不了.

# 如何安裝

```bash
go get -u -v github.com/Lofanmi/chinese-calendar-golang
```

# 用法

```bash
# 確保使用的是東八區(北京時間)
export TZ=PRC

# 查看時間
date -R
```

```go
import (
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/calendar"
)

func xxx() {
    t := time.Now()
    // 1. ByTimestamp
    // 時間戳
    c := calendar.ByTimestamp(t.Unix())
    // 2. BySolar
    // 公曆
    c := calendar.BySolar(year, month, day, hour, minute, second)
    // 3. ByLunar
    // 農曆(最後一個參數表示是否閏月)
    c := calendar.ByLunar(year, month, day, hour, minute, second, false)

    bytes, err := c.ToJSON()

    fmt.Println(string(bytes))
}
```

# 原理

1. 公曆的計算比較容易, 生肖, 星座等的轉換見源碼, 不囉嗦哈;

2. 農曆的數據和算法參考了[chinese-calendar](https://github.com/overtrue/chinese-calendar)的實現, 還有 [1900年至2100年公曆、農曆互轉Js代碼 - 晶晶的博客](http://blog.jjonline.cn/userInterFace/173.html), 節氣的捨去了, 因為在干支歷和生肖的計算上不夠精確;

3. 干支歷的轉換是最繁瑣的, 其計算依據是二十四節氣. 我結合了網上的文章, 用天文方法計算了 1904-2024 的二十四節氣時間, 精確到秒, 並與[香港天文台](http://data.weather.gov.hk/gts/astron2018/Solar_Term_2018_uc.htm)進行比較, 誤差應該在 `1分鐘` 以內, 大家如果有以往的數據, 可以看下源碼比對一下.

# 使用須知

1. 干支歷的時間範圍只有 1904-2024 年, 主要也是因為 1904 年之前的時間戳, 在32位的 `PHP` 下會識別不出(溢出), 而且再早的時間, 其實意義也不大, 2024之後...好像也不會用到吧. 研究二十四節氣這個算法的時間花了很長很長, 反而擼代碼的時間不算太多;

2. 實際上, 農曆的時間也是可以根據天文方法計算出來的. 計算的過程, 也需要先計算二十四節氣和每月的正月初一的日期(日月合朔), 才能知道閏月的信息(所以農曆是陰陽曆). 不過已經有了數據, 我也就懶了, 直接拿來用...所以農曆的算法我還沒有實現;

3. 農曆的時間範圍是 1900-2100 年, 但是這個日曆只支持到 1904-2024 年, 如果有需要我再加吧, 後續會有個 `PHP` 版本, 輸出的 `JSON` 格式會保持一致.

# 公曆(陽曆) - 字段說明

```
{
    // 生肖, 以立春分界
    //     如 2018-02-04 05:28:26 立春
    //        2018-02-04 04:59:59 屬雞年
    //        2018-02-04 05:00:00 屬狗年
    "animal": "鼠",

    // 星座
    "constellation": "天秤",

    // 今年是否為閏年
    "is_leep": true,

    // 年
    "year": 2020,
    // 月
    "month": 9,
    // 日
    "day": 20,

    // 時
    "hour": 5,
    // 分
    "minute": 15,
    // 秒
    "second": 26,
    // 納秒
    "nanosecond": 0,

    // 星期日
    "week_alias": "日",
    // 星期序數, 0表示週日
    "week_number": 0,
}
```

# 農曆(陰陽曆) - 字段說明

```
{
    // 生肖, 以每年正月初一分界
    "animal": "鼠",

    // 年
    "year": 2020,
    // 年(漢字)
    "year_alias": "二零二零",
    // 月
    "month": 8,
    // 月(漢字)
    "month_alias": "八月",
    // 日
    "day": 4,
    // 日(漢字)
    "day_alias": "初四",

    // 是否閏年
    "is_leap": true,
    
    // 這個月是否為閏月
    "is_leap_month": false,

    // 今年閏四月
    "leap_month": 4,
}
```

# 干支歷 - 字段說明

```
{
    // 生肖, 以立春分界
    //     如 2018-02-04 05:28:26 立春
    //        2018-02-03 05:28:25 屬雞年
    //        2018-02-04 05:28:26 屬狗年
    "animal": "鼠",

    // 年干支
    "year": "庚子",
    // 年干支六十甲子序數
    "year_order": 37,

    // 月干支
    "month": "乙酉",
    // 月干支六十甲子序數
    "month_order": 22,

    // 日干支
    "day": "丙寅",
    // 日干支六十甲子序數
    "day_order": 3,

    // 時干支
    "hour": "辛卯",
    // 時干支六十甲子序數
    "hour_order": 28,
}
```

# 輸出示例

```json
{
    "ganzhi": {
        "animal": "鼠",
        "day": "丙寅",
        "day_order": 3,
        "hour": "辛卯",
        "hour_order": 28,
        "month": "乙酉",
        "month_order": 22,
        "year": "庚子",
        "year_order": 37
    },
    "lunar": {
        "animal": "鼠",
        "day": 4,
        "day_alias": "初四",
        "is_leap": true,
        "is_leap_month": false,
        "leap_month": 4,
        "month": 8,
        "month_alias": "八月",
        "year": 2020,
        "year_alias": "二零二零"
    },
    "solar": {
        "animal": "鼠",
        "constellation": "天秤",
        "day": 20,
        "hour": 5,
        "is_leep": true,
        "minute": 15,
        "month": 9,
        "nanosecond": 0,
        "second": 26,
        "week_alias": "日",
        "week_number": 0,
        "year": 2020
    }
}
```

# TODO

1. 完善單元測試
2. 完善註釋
3. 完善文檔
4. 支持更大範圍的時間 / 把農曆的算法實現一下?

# 參考資料

- [算法系列之十八：用天文方法計算二十四節氣（上）](https://blog.csdn.net/orbit/article/details/7910220)
- [算法系列之十八：用天文方法計算二十四節氣（下）](https://blog.csdn.net/orbit/article/details/7944248)
- [overtrue/chinese-calendar](https://github.com/overtrue/chinese-calendar)
- [1900年至2100年公曆、農曆互轉Js代碼 - 晶晶的博客](http://blog.jjonline.cn/userInterFace/173.html)
- [香港天文台](http://data.weather.gov.hk/)
- [五虎遁元](https://baike.baidu.com/item/%E4%BA%94%E8%99%8E%E9%81%81%E5%85%83/5471492)
- [五鼠遁元](https://baike.baidu.com/item/%E4%BA%94%E9%BC%A0%E9%81%81%E5%85%83/5471935)
- [NASA](https://eclipse.gsfc.nasa.gov/)

# License

MIT