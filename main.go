package main

import (
	"fmt"
)

func main(){
	var n  int = 3
	var arr = []int{1,2,3,4,5,6}
	// fmt.Scan(&n)
	twoSum(arr,n)

}

func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums)-1;i++{
		for j := i+1; j<len(nums);j++{
			
		}
	}
	return nums
}

//bubble sort
// func main() {

// 	var arr = []int{43,5,-2,32,2,0}
// 	fmt.Println(arr)

// 	for i:=0; i< len(arr)-1; i++ {
// 		for j:=0; j < len(arr)-i-1; j++ {
// 		   if (arr[j] > arr[j+1]) {
// 			  arr[j], arr[j+1] = arr[j+1], arr[j]
// 		   }
// 		}
// 	}
// 	fmt.Println(arr)
// }