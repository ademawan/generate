package main

import "time"

func main() {

}

func CheckIsNew(createdAt string) (float64, bool) {
	var status bool
	layoutFormat := "2006-01-02T15:04:05Z07:00"
	timeCreate, _ := time.Parse(layoutFormat, createdAt)
	timeNow := time.Now().Format(time.RFC3339)
	timeNowNew, _ := time.Parse(layoutFormat, timeNow)
	diff := timeNowNew.Sub(timeCreate.Local())
	if diff.Hours() > float64(24) {

		return diff.Hours(), status
	}
	status = true
	return diff.Hours(), status
}
