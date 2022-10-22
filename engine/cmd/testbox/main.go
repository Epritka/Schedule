package main

import (
	"fmt"
	"time"
)

func NumberOfTheWeekInMonth(now time.Time) int {
	beginningOfTheMonth := time.Date(now.Year(), now.Month(), 1, 1, 1, 1, 1, time.UTC)
	_, thisWeek := now.ISOWeek()
	fmt.Println(thisWeek)
	_, beginningWeek := beginningOfTheMonth.ISOWeek()
	return 1 + thisWeek - beginningWeek
}

func main() {
	fmt.Printf("This is %d, the %d week in the month", time.Now().Day(), NumberOfTheWeekInMonth(time.Now()))
}
