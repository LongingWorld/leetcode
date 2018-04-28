package main

import "fmt"

var ints = []int{3,3,3}
var target = 6
func main() {
	fmt.Println(twoSum(ints,target))
	fmt.Println(twoSum1(ints,target))
}
//using map to save the traversed elements of the slice
func twoSum(ints []int,target int) [][]int  {
	result := [][]int{}
	intMaps := make(map[int]int)
	var diff = 0
	for i:=0;i < len(ints);i++  {
		diff = target -ints[i]
		if  _,ok := intMaps[diff]; ok {
			if intMaps[i] != i {
				result = append(result,[]int{i,intMaps[diff]})
			}
		}
		intMaps[ints[i]] = i

	}
	return result
}

func twoSum1(ints []int, target int) []int {
	intMaps := make(map[int]int)
	var res []int
	var diff = 0
	for i:=0;i < len(ints);i++  {
		diff = target -ints[i]
		if  _,ok := intMaps[diff]; ok {
			if intMaps[i] != i  {
				res = append(res,intMaps[diff],i)
			}
		}
		intMaps[ints[i]] = i

	}
	return  res
}