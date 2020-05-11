package main

import (
	"fmt"
	bst2 "offerPCom/bst"
)

func main(){
	oriTreeData := []string{"D", "B", "H", "E", "I", "A", "F", "K", "C", "G"}
	bst := bst2.InitBST()
	for _, value := range oriTreeData{
		bst.InsertNode(value)
	}
	preO := bst.PreOrder()
	//inO := bst.InOrder()
	preO = preO
	fmt.Println("before delete preO: ", preO)
	postO := bst.PostOrder()
	fmt.Println("before delete postO: ", postO)
	if maxNode := bst.Max(); maxNode != nil{
		fmt.Println("max node value is: ", maxNode.Ele)
	}
	if minNode := bst.Min();minNode != nil{
		fmt.Println("min node value is: ", minNode.Ele)
	}

	levelOrder := bst.LevelOrder()
	fmt.Println("bst level order is: ", levelOrder)

	//bst.DeleteNode("D")
	//preO = bst.PreOrder()
	//inO = bst.InOrder()
	//postO = bst.PostOrder()
	//fmt.Println("after delete ", "D", " preOrder is ", preO)
	//fmt.Println("after delete ", "D", " inOrder is ", inO)
	//for _, value := range inO{
	//	nextNode := bst.InOrderNextNode(value)
	//	if nextNode != nil{
	//		fmt.Println(value, "next node value is ", nextNode.Ele)
	//	}else{
	//		fmt.Println(value, "next node value is nil")
	//	}
	//
	//}
	//fmt.Println("after delete ", "D", " postOrder is ", postO)
	//
	//
	//bst.DeleteNode("K")
	//preO = bst.PreOrder()
	//inO = bst.InOrder()
	//postO = bst.PostOrder()
	//fmt.Println("after delete ", "K", " preOrder is ", preO)
	//fmt.Println("after delete ", "K", " inOrder is ", inO)
	//fmt.Println("after delete ", "K", " postOrder is ", postO)
	//
	//bst.DeleteNode("A")
	//preO = bst.PreOrder()
	//inO = bst.InOrder()
	//postO = bst.PostOrder()
	//fmt.Println("after delete ", "A", " preOrder is ", preO)
	//fmt.Println("after delete ", "A", " inOrder is ", inO)
	//fmt.Println("after delete ", "A", " postOrder is ", postO)

	LevelOrderTest()
}

func LevelOrderTest(){
	bst := bst2.InitBST()
	initData := []string{"F","B","A","D","H","G","I","C","E"}
	for _, value := range initData{
		bst.InsertNode(value)
	}
	levOrder := bst.LevelOrder()
	fmt.Println("level order is: ", levOrder)
	levelSliceOrder := bst.LevelSliceOrder()
	fmt.Println("level slice order is: ", levelSliceOrder)
	turnLevOrder := bst.TurnLevelOrder()
	fmt.Println("turn level order is: ", turnLevOrder)

}

/*

				F
			   / \
			  B   H
             / \  / \
            A  D  G  I
              /  \
             C    E

*/