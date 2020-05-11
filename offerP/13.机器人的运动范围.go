package main

import "fmt"

// 判断当前的运动坐标是否小于最大值
func isInArea(positionX, positionY, maxNumber int)bool{
	count := 0
	for ;positionX>0;{
		count += positionX%10
		positionX = positionX/10
	}
	for ;positionY>0;{
		count += positionY%10
		positionY = positionY/10
	}
	if count <= maxNumber{
		return true
	}else{
		return false
	}
}

// 判断某一个节点是否被访问了
func isVisited(target int)bool{
	if target != 0 {
		return true
	}else{
		return false
	}
}

func sportArea(matrix [][]int, maxData int, positionX int, positionY int, maxPositionX, maxPositionY int){
	directions := [][]int{{0, 1}, {1, 0}}
	if positionX >=0 && positionX < maxPositionX && positionY >= 0 && positionY < maxPositionY{
	}else{
		return
	}

	if isInArea(positionX, positionY, maxData) && !isVisited(matrix[positionX][positionY]) {
		matrix[positionX][positionY] = 1
		for _, direction := range directions{
			currentX, currentY := positionX + direction[0], positionY + direction[1]
			sportArea(matrix, maxData, currentX, currentY, maxPositionX, maxPositionY)
		}
	}
}

func main(){
	matrix := make([][]int, 0)
	for j:=0;j<40;j++{
		item := make([]int, 0)
		for i:=0;i<40;i++{
			item = append(item, 0)
		}
		matrix = append(matrix, item)
	}
	yMaxLen := len(matrix)
	xMaxLen := len(matrix[0])
	sportArea(matrix, 18, 0, 0, yMaxLen, xMaxLen)
	count := 0
	for _, item := range matrix{
		for i:=0;i<len(item);i++{
			if item[i] == 1{
				count += 1
			}
		}
	}
	fmt.Println(count)
}
