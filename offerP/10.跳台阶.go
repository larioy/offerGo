package main

import "fmt"

// 每次可以跳一级也可以跳2级， 采用自顶向下推测， fn = f(n-1)+f(n-2)
func JumpStepTwo(number int)(int){
	i := 0
	j := 1
	for count:=0; count<number;count++{
		j = i + j
		i = j
	}
	return j
}

// 一次跳多级台阶, 根据自顶向下的推理， fn = f(1)+f(2)+...f(n-1), 其实就是一个阶乘
func JumpStepMany(number int)(int){
	current := 0
	for count := 0; count < number; count++{
		i := 0
		j := 1
		// 每次计算出count下的值， 然后进行相加
		for m := 0; m < count; m ++{
			i = j
			j = i + j
		}
		current += j
	}
	return current
}

func main(){
	num := JumpStepTwo(10)
	fmt.Println("total count is: ", num)
	num1 := JumpStepMany(10)
	fmt.Println("total count is: ", num1)
}
