package main

import "fmt"


func Fib()func()(int){
	i, j := 0, 1
	target := i  // 闭包需要一个中间变量来进行暂存
	innerFunc := func ()(int){
		target = i
		i = j
		j = target + j
		return target
	}
	return  innerFunc
}


// 通过goroutine来创建斐波那锲数列迭代器
func generateFib(count int)(<-chan int){
	i, j := 0, 1
	number := 0
	fFib := make(chan int)
	go func (){
		for {
			if number == count{
				close(fFib)
				break
			}
			fFib <- i
			i, j = j, i + j
			number ++
		}
	}()
	return  fFib
}



func main(){
	//fib := Fib()
	//for i := 0; i < 10; i++{
	//	fmt.Println(fib())
	//}

	for value := range generateFib(6){
		fmt.Println(value)
	}
}