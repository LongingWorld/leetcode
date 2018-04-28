package main

import "fmt"

func main() {
	nums := []int{0,0,1,1,2,2,3,3,4,4,5,5,6,6,7,7,8,8,9,9}
	count := removeDuplicates(nums)
	fmt.Println("length of slice nums is ",count)
	for i:=0;i < count ;i++  {
		fmt.Printf("%d ",nums[i])
	}
	fmt.Println()
}


/*Remove Duplicates from Sorted Array
Given a sorted array nums, remove the duplicates in-place such
that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this
by modifying the input array in-place with O(1) extra memory.
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0{    //length of slice is zero,return 0
		return 0
	}
	size := 0  
	for i,_ := range nums  {
		if nums[size] != nums[i] {
			size+=1
			nums[size] =nums[i]
		}
	}
	nums = nums[0:size+1]
	return size + 1
}