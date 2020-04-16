package main

import (
	"fmt"
	"offerPCom/common"
)

// 树节点结构体
type Node struct{
	left *Node
	right *Node
	pre *Node
	ele interface{}
}

type BST struct{
	root *Node
}

func (bst *BST)PreOrder()([]interface{}){
	// 树的中序遍历
	preList := make([]interface{}, 0)
	root := bst.root
	bst._preOrder(root, &preList)
	return preList
}

func (bst *BST)_preOrder(root *Node, list *[]interface{}){
	if root == nil{
		return
	}
	//fmt.Printf("%s,", root.ele)
	*list = append(*list, root.ele)
	bst._preOrder(root.left, list)
	bst._preOrder(root.right, list)
}

func (bst *BST)InOrder()(orderList []interface{}){
	// 中序遍历
	orList := make([]interface{}, 0)
	root := bst.root
	bst._inOrder(root, &orList)
	return orList
}

func (bst *BST)_inOrder(root *Node, list *[]interface{}){
	if root == nil{
		return
	}
	bst._inOrder(root.left, list)
	*list = append(*list, root.ele)
	//fmt.Printf("%s,",root.ele)
	bst._inOrder(root.right, list)
}





func generateBST(preOrder []interface{}, inOrder []interface{})*Node{
	var root *Node
	if len(preOrder) == 0{
		return nil
	}
	if len(preOrder) == 1{
		root = &Node{ele:preOrder[0]}
		return root
	}
	root = &Node{ele:preOrder[0]}
	pos := common.GetIndex(preOrder[0], inOrder)
	root.left = generateBST(preOrder[1:pos+1], inOrder[:pos])
	root.right = generateBST(preOrder[pos+1:], inOrder[pos+1:])
	return root
}

// 用于验证生成的树的中序遍历和前序遍历是否和给的输入是否一致
func VerifyTree(preOrder []interface{}, inOrder []interface{}){
	fmt.Println("oriPreOrder: ", preOrder)
	fmt.Println("oriInOrder: ", inOrder)
	root := generateBST(preOrder, inOrder)
	bst := BST{root:root}
	bstPreOrder := bst.PreOrder()
	fmt.Println("bstPreOrder: ", bstPreOrder)
	bstInOrder := bst.InOrder()
	fmt.Println("bstInOrder: ", bstInOrder)
}

func main(){
	/*
		测试元素类型是否一致
		testA := int(2)
		preOrder := []int{1,2,3,4,5,6,7,9}
		fmt.Println(IsSameType(testA, preOrder[0]))
	*/
	preOrder := []interface{}{"A", "B", "D", "E", "H", "I", "C", "F", "K", "G"}
	inOrder := []interface{}{"D", "B", "H", "E", "I", "A", "F", "K", "C", "G"}
	VerifyTree(preOrder, inOrder)

	// 下面为验证bst的create后， 进行前序遍历和中序遍历得到的序列是否对
	preOrder = []interface{}{"D", "B", "A", "C", "H", "E", "F", "G", "I", "K"}
	inOrder = []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "K"}
	VerifyTree(preOrder, inOrder)
}
