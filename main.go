package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	weekday := now.Weekday()

	fmt.Println(weekday)
}
