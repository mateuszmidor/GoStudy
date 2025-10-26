package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()

	p(now)
	then := time.Date(2009, 11, 17, 20, 34, 58, 654312, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())   // tuesday
	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	diff := now.Sub(then) // 92850h43m30.525018143s
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))  // 2020-06-21 15:19:48.293943545 +0000 UTC
	p(then.Add(-diff)) // 1999-04-16 01:50:07.707365079 +0000 UTC
}
