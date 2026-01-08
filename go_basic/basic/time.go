package main

import (
	"fmt"
	"time"
)

const (
	TIME_FMT = "2006-01-02 15:04:05.000"
	DATE_FMT = "20060102"
)

func parse_format() {
	t := time.Now()

	fmt.Printf("Current Time: %s\n", t.Format(TIME_FMT))
	fmt.Printf("Current Date: %s\n", t.Format(DATE_FMT))
	ts := t.Format(TIME_FMT)

	t2, err := time.Parse(TIME_FMT, ts)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return
	}
	fmt.Printf("Parsed Time: %s\n", t2.Format(TIME_FMT))

	loc, _ := time.LoadLocation("Asia/Shanghai")
	t3, err := time.ParseInLocation(TIME_FMT, ts, loc)
	if err != nil {
		fmt.Printf("Error parsing time in location: %v\n", err)
		return
	}
	fmt.Printf("Parsed Time in Location: %s\n", t3.Format(TIME_FMT))
}

// 周期执行任务
func ticker() {
	fmt.Printf("当前时间是: %s\n", time.Now().Format(TIME_FMT))
	tk := time.NewTicker(time.Second)
	defer tk.Stop()

	for i := 0; i < 10; i++ {
		<-tk.C
		fmt.Printf("当前时间是: %s\n", time.Now().Format(TIME_FMT))
	}
}

func timer() {
	fmt.Printf("当前时间是: %s\n", time.Now().Format(TIME_FMT))
	tm := time.NewTimer(time.Second)
	defer tm.Stop()
	<-tm.C
	fmt.Printf("1秒后时间是: %s\n", time.Now().Format(TIME_FMT))

	fmt.Printf("当前时间是: %s\n", time.Now().Format(TIME_FMT))
	<-time.After(time.Second)
	fmt.Printf("1秒后时间是: %s\n", time.Now().Format(TIME_FMT))
}

func main() {
	// parse_format()
	// ticker()
	timer()
}
