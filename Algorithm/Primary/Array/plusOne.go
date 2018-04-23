package main

import (
	"fmt"
	"math"
)

func plusOne(digits []int) []int  {
	var count int64
	var num float64
	var maxNum = math.Pow10(len(digits))
	var maxPow = len(digits)
	var result []int
	for i,v :=range digits {
		//fmt.Println(i,v)
		num = math.Pow10(len(digits)-i-1)
		count += int64(num*float64(v))
	}
	fmt.Println("count is ",count)
	count +=1
	fmt.Println("count is ",count)
	fmt.Println("maxPow is ",maxPow)
	if count/int64(maxNum) > 0 {
		maxPow += 1
	}
	fmt.Println("maxPow is ",maxPow)
	for i := 0;i < int(maxPow);i++{
		con := count / int64(math.Pow10(maxPow-i-1))
		count = count % int64(math.Pow10(maxPow-i-1))
		fmt.Printf("con :%d i :%d count :%d\n",con,i,count)
		result = append(result,int(con))
	}
	return result
}


/*Given a non-empty array of digits representing a non-negative integer,
plus one to the integer.The digits are stored such that the most significant
digit is at the head of the list,and each element in the array contain a
single digit.You may assume the integer does not contain any leading zero,
except the number 0 itself.*/
/*
将一个数字的每个位上的数字分别存到一个一维向量中，最高位在最开头，我们需要给这个数字加一，即在末尾数字加一，
如果末尾数字是9，那么则会有进位问题，而如果前面位上的数字仍为9，则需要继续向前进位。
具体算法如下：
首先判断最后一位是否为9，若不是，直接加一返回，若是，则该位赋0，再继续查前一位，同样的方法，知道查完第一位。
如果第一位原本为9，加一后会产生新的一位，那么最后要做的是，查运算完的第一位是否为0，如果是，则在最前头加一个1
*/
func plusOnePlus(digits []int) []int {
	 nums := len(digits)
	for i := nums -1;i >= 0;i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	digits[0] =1
	digits = append(digits,0)
	return  digits
}

var digit = []int{9,9,9,9}
func main() {
	fmt.Println(plusOne(digit))
	fmt.Println(plusOnePlus(digit))
}
