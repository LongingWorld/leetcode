package main

import "fmt"

//{0,1,2,3,0,4,5,6,0}
//{1,23,0,0,123}
//{0,0,1}
//{1,0,0}
func main() {
	number := []int{1, 23, 0, 0, 0, 0, 123}
	moveZeroes(number)
	fmt.Println(number)
}

//将非零数字向左移动，pos记录左边非零数字个数
func moveZeroes(ints []int) {
	pos := 0
	for i := 0; i < len(ints); i++ {
		if ints[i] != 0 {
			ints[pos] = ints[i]
			pos++
		}
	}
	for i := pos; i < len(ints); i++ {
		ints[i] = 0
	}
	/*for i:=0;i <= tarZero;i++  {
		if ints[tarZero] != 0 {
			if ints[i] == 0 {
				for j:=i; j <= tarZero; j++ {
					if j + 1 > tarZero {
						break
					}
					ints[j] = ints[j+1]
				}
				ints[tarZero] = 0
				tarZero--
			}
		}else {
			i--
			tarZero--
		}
	}*/
}
