package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Weekday())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2021-07-23")
	fmt.Println(parse)

	fmt.Println(now.Local())

}
