package main

import "fmt"

var nums = []int{1,4,2,3,5,6,7,8,9,1,23}
func main() {
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool  {
	con := make(map[int]bool)
	for _,v := range nums{
		if !con[v] {
			con[v] = true
		}else {
			return true
		}
	}
	return false
}