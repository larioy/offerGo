package main

import (
	"fmt"
	"offerPCom/queue"
)

type Stack struct {
	Q *queue.Queue
}

func InitStack()(st *Stack){
	st = &Stack{}
	st.Q = queue.InitQueue()
	return st
}

func (st *Stack)Push(ele interface{}){
	size := st.Q.Size
	st.Q.Push(ele)
	for i:=0;i<size;i++{
		st.Q.Push(st.Q.Pop())
	}
}

func (st *Stack)Pop()(ele interface{}){
	return st.Q.Pop()
}

func stackT(pushStackData []interface{})([]interface{}){
	popStackData := make([]interface{}, 0)
	st := InitStack()
	for _, value := range pushStackData{
		st.Push(value)
	}
	for value := st.Pop(); value != nil;{
		popStackData = append(popStackData, value)
		value = st.Pop()
	}
	return popStackData
}

func main(){
	pushStackData1 := []interface{}{}
	pushStackData2 := []interface{}{"A"}
	pushStackData3 := []interface{}{"A", "B", "C", "D", "E", "F"}
	popStackData1 := stackT(pushStackData1)
	popStackData2 := stackT(pushStackData2)
	popStackData3 := stackT(pushStackData3)
	fmt.Println("pushStackData1 is: ", pushStackData1, "popStackData1 is: ", popStackData1)
	fmt.Println("pushStackData2 is: ", pushStackData2, "popStackData2 is: ", popStackData2)
	fmt.Println("pushStackData3 is: ", pushStackData3, "popStackData3 is: ", popStackData3)
}
