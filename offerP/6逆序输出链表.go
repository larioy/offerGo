package main

import (
	"fmt"
	"offerPCom/common"
)

type lNode struct{
	pre *lNode
	next *lNode
	ele interface{}
}

type Chain struct {
	Head *lNode
	Size int
}

func (ch *Chain)Create(arr []interface{}){
	if len(arr) == 0{
		return
	}
	ch.Head = &lNode{}
	head := ch.Head
	for _, value := range arr{
		node := lNode{pre:head, next:nil, ele:value}
		head.next = &node
		head = &node
		ch.Size += 1
	}
}

func (ch *Chain)Traverse()(list []interface{}){
	head := ch.Head.next
	_list := make([]interface{},0)
	for ; head != nil;{
		_list = append(_list, head.ele)
		head = head.next
	}
	return _list
}


// 对链表进行逆序操作
func (ch *Chain)Reverse()(rList []interface{}){
	stack := make([]interface{}, 0)
	head := ch.Head.next
	for ; head != nil;{
		stack = common.SliceInsert(0, head.ele, stack)
		head = head.next
	}
	return stack
}

func main(){
	initList := []interface{}{"a", "b", "c", "d", "e", "f"}
	fmt.Println("initList: ", initList)
	chain := Chain{Size:0}
	chain.Create(initList)
	list := chain.Traverse()
	fmt.Println("list: ", list)
	rlist := chain.Reverse()
	fmt.Println("rlist: ", rlist)
}
