package main

import "fmt"

// 查找翻转数组的最小值
func FindRotateMinNumber(arr []int)(interface{}){
	arrLen := len(arr) - 1
	if arrLen == -1{
		return nil
	}
	leftPosition := 0
	rightPosition := arrLen
	// 注意每个if/else里面进行判断的是右节点
	for ; leftPosition < rightPosition;{
		mid := leftPosition + (rightPosition-leftPosition)/2    // 这里的取半折中计算方式， 要注意
		// 中间点小于最有点， 中间点可能是最小的值
		if arr[mid] < arr[rightPosition]{
			rightPosition = mid
		// 中间值大于最右值， 中间值肯定不是最小值， 所以加一
		}else if arr[mid] > arr[rightPosition]{
			leftPosition = mid + 1
		}else {
			// 中间点等于右节点， 将右节点减一
			rightPosition -= 1
		}
	}
	return arr[leftPosition]
}


func main(){
	testArr1 := []int{3,4,5,1,2}
	testArr2 := []int{1,1,1,1,0,1}
	testArr3 := []int{2,2,2,2,2,2}
	testArr4 := []int{1}
	fmt.Println("testArr1 min is: ", FindRotateMinNumber(testArr1))
	fmt.Println("testArr2 min is: ", FindRotateMinNumber(testArr2))
	fmt.Println("testArr3 min is: ", FindRotateMinNumber(testArr3))
	fmt.Println("testArr4 min is: ", FindRotateMinNumber(testArr4))
}
