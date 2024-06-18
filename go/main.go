package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	y, m, d := now.Date()
	today := time.Date(y, m, d, now.Hour(), now.Minute(), 0, 0, time.Local)

	lw := today.AddDate(0, 0, -7)

	from, to := strconv.FormatInt(lw.Unix(), 10), strconv.FormatInt(today.Unix(), 10)

	fmt.Printf("%s 〜 %s\n", lw, today)
	fmt.Printf("%s 〜 %s\n", from, to)
}
