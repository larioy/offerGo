package main

import (
	"fmt"
	"offerPCom/stack"
)


type Queue struct{
	St1 *stack.Stack
	St2 *stack.Stack
}

func InitQueue()(q *Queue){
	q = &Queue{}
	q.St1 = stack.InitStack()
	q.St2 = stack.InitStack()
	return q
}

func (q *Queue)Push(ele interface{}){
	q.St1.Push(ele)
}

func (q *Queue)_shift(){
	st1Len := q.St1.Size
	for i:=0; i<st1Len;i++{
		q.St2.Push(q.St1.Pop())
	}

}

func (q *Queue)Pop()(ele interface{}){
	if q.St2.Size == 0{
		q._shift()
	}
	if q.St2.Size == 0{
		return nil
	}else{
		return q.St2.Pop()
	}
}


func queueT(pushQueueData []interface{})([]interface{}){
	popQueueData := make([]interface{}, 0)
	q := InitQueue()
	for _, value := range pushQueueData{
		q.Push(value)
	}
	for value := q.Pop(); value != nil;{
		popQueueData = append(popQueueData, value)
		value = q.Pop()
	}
	return popQueueData
}

func main(){
	pushQueueData1 := []interface{}{}
	pushQueueData2 := []interface{}{"A"}
	pushQueueData3 := []interface{}{"A", "B", "C", "D", "E", "F"}
	popQueueData1 := queueT(pushQueueData1)
	popQueueData2 := queueT(pushQueueData2)
	popQueueData3 := queueT(pushQueueData3)
	fmt.Println("pushQueueData1 is: ", pushQueueData1, "popQueueData1 is: ", popQueueData1)
	fmt.Println("pushQueueData2 is: ", pushQueueData2, "popQueueData2 is: ", popQueueData2)
	fmt.Println("pushQueueData3 is: ", pushQueueData3, "popQueueData3 is: ", popQueueData3)
}
