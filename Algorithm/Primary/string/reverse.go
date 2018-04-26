package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-629))
}

func reverse(x int) int  {
	var result int64
	for x != 0 {
		result = result * 10 + int64(x) % 10
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return int(result)
}