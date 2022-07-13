package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	interval := [][]string{
		{"2022-04-08 10:04:02", "3"},
		{"2022-04-08 11:12:35", "2"},
		{"2022-04-08 17:08:04", "3"},
		{"2022-04-11 16:30:42", "3"},
		{"2022-04-08 09:04:15", "7"},
		{"2022-04-09 00:59:18", "2"},
		{"2022-04-09 00:40:18", "6"},
	}
	sort.Slice(interval, func(i, j int) bool {
		t1, err := time.Parse("2006-01-02 15:04:05", interval[i][0])
		t2, err := time.Parse("2006-01-02 15:04:05", interval[j][0])
		if err == nil {
			return t1.Before(t2)
		}
		return false
	})
	for _, in := range interval{
		fmt.Println(in)
	}


	time2 := "2022-04-14 20:16:28"
	time1 := "2022-04-18 16:19:23"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		fmt.Println("true")
	}


}