package main

import "fmt"

// 判断某一个节点是否被访问了
func hasVisited(target int)bool{
	if target != 0 {
		return true
	}else{
		return false
	}
}

func findPath(matrix [][]string, matrixFlag [][]int,  targetStr []string, positionX int, positionY int,
	maxPositionX, maxPositionY int)bool{
	directions := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	if positionX >= 0 && positionX <= maxPositionX &&  positionY >= 0 && positionY <= maxPositionY  {
		if len(targetStr) == 1  {
			if matrix[positionX][positionY] == targetStr[0]{
				matrixFlag[positionX][positionY] = 1
				return true
			}else{
				return false
			}
		}
	}else{
		return false
	}
	if !hasVisited(matrixFlag[positionX][positionY]) && matrix[positionX][positionY] == targetStr[0]{
		matrixFlag[positionX][positionY] = 1
		for _, direction := range directions{
			currentX, currentY := positionX + direction[0], positionY + direction[1]
			// 如果节点已经访问过了, 那么不要重新进行递归函数
			if currentX >= 0 && currentX <= maxPositionX &&  currentY >= 0 && currentY <= maxPositionY &&
				hasVisited(matrixFlag[currentX][currentY]){
				continue
			}
			if findPath(matrix, matrixFlag, targetStr[1:], currentX, currentY, maxPositionX, maxPositionY){
				return true
			}
		}
		matrixFlag[positionX][positionY] = 0
	}
	return false
}

func main(){
	matrix := [][]string{{"A","B","C","E"}, {"S","F","C","S"}, {"A","D","E","E"} }
	matrixFlag := [][]int{{0,0,0,0}, {0,0,0,0}, {0,0,0,0} }
	xMaxLen := len(matrix[0])-1
	yMaxLen := len(matrix)-1
	fmt.Println("generation matrix is ")
	for _, item := range matrix{
		fmt.Println(item)
	}
	targetStr := []string{"B","C", "C", "E"}

	flag := 0
	for i:=0;i<xMaxLen;i++{
		for j:=0;j<yMaxLen;j++{
			if findPath(matrix, matrixFlag, targetStr, i, j, len(matrix), len(matrix[0])){
				fmt.Println("find path")
				for _, item := range matrixFlag{
					fmt.Println(item)
				}
				flag = 1
				break
			}
		}
		if flag==1{
			break
		}
	}
}

/*
存在疑问的问题

*/