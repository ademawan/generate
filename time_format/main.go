package main

import (
	"fmt"
	"time"
)

func main() {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2023-07-07 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	res := date.Format("Jan 02, 2006 15:04:05 GMT+7")
	fmt.Println(value, "\t->", date)
	fmt.Println(value, "\t->", res)

}
