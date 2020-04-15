package main

import "fmt"

func searchData(target int, data [][]int)(inArray bool){
	rows := len(data)
	cols := len(data[0]) - 1
	row := 0
	for ; row < rows && cols >0;{
		if target > data[row][cols]{
			row++
		}else if target == data[row][cols]{
			fmt.Println(row, cols)
			return true
		}else{
			cols--
		}
	}
	return false
}

func main(){
	var initData = make([][]int, 0)
	initData = [][]int{{1,2,8,9}, {2,4,9,12}, {4,7,10,13}, {6,8,11,15}}
	for index, value := range initData{
		fmt.Println(index, value)
	}
	target := 10
	inArray := searchData(target, initData)
	fmt.Println(target, " in ", initData, ": ", inArray)
}


