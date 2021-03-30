package main

import (
	"fmt"
)

func main() {
	count := 12
	suffix := "minutes of bonus footage"
	format := "DVD"
	language := append([]string{}, "English", "Korean")
	fmt.Println(count, suffix, format, language)
}
