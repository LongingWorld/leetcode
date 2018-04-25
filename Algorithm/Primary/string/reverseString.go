package main

import "fmt"

func main() {
	var s = "Hello"
	fmt.Println(reverseString(s))
}

func reverseString(s string) string  {
	str := make([]rune, len(s))
	count := len(s)-1
	for _,v := range s  {
		str[count] = v
		count--
	}
	return string(str)
}
