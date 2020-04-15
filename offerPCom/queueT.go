package main

import (
	"fmt"
	"offerPCom/queue"
)


func queueT(pushQueueData []interface{})([]interface{}){
	popQueueData := make([]interface{}, 0)
	q := queue.InitQueue()
	for _, value := range pushQueueData{
		st.Push(value)
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

