package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("hello world")

	//timeLocation, _ := time.LoadLocation("America/Mexico_City")
	//time, _ := time.ParseInLocation("2006-01-02 15:04:05", "2024-08-01 00:00:00", timeLocation)
	//fmt.Println(time)

	var cstZone = time.FixedZone("CST", -6*3600)       // 东八区
	t := time.Now().In(cstZone)
	fmt.Println("SH : ", t.Format("2006-01-02 15:04:05"))
	var mxZone = time.FixedZone("CST",  8*3600)       // 东八区
	mt := t.In(mxZone)
	fmt.Println("MX : ", mt.String())

	mxTimeLocation, _ := time.LoadLocation("America/Mexico_City")
	mxTime, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:04:05"), mxTimeLocation)
	fmt.Println("MX : ", mxTime.String())
}
