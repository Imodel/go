package main 

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)

	t = time.Now().UTC()

	fmt.Println(t)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(week)

	fmt.Println(week_from_now)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2016 11:18"))
	s := t.Format("20160102")
	fmt.Println(t,"=>",s)
}
