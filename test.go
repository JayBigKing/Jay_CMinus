package main

import "fmt"

const baseIndex123 = 10 //起始索引
const (
	SUNDAY = baseIndex123 + iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

const (
	SUNDAY1 = SATURDAY + 1 + iota
	MONDAY1
	TUESDAY1
	WEDNESDAY1
	THURSDAY1
	FRIDAY1
	SATURDAY1
)

func main() {

	fmt.Print(FRIDAY1)
}
