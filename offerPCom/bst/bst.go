package bst

// 树节点结构体
type Node struct{
	left *Node
	right *Node
	pre *Node
	Ele string
}

type BST struct{
	root *Node
	SIZE int
}

func InitBST()(root *BST){
	root = &BST{root:nil, SIZE:0}
	return root
}


func (bst *BST)InsertNode(value string){
	root := bst.root
	bst.root = bst._insertNode(root, value)
}


// 二叉树插入元素函数
func (bst *BST)_insertNode(node *Node, value string)(root *Node){
	if node == nil{
		root = &Node{Ele:value}
		return
	}
	if node.Ele > value {
		leftNode := bst._insertNode(node.left, value)
		node.left = leftNode
		leftNode.pre = node
		bst.SIZE += 1
	}else if node.Ele < value {
		rightNode := bst._insertNode(node.right, value)
		node.right = rightNode
		rightNode.pre = node
		bst.SIZE += 1
	}else{
		return node
	}
	return node
}


//
func (bst *BST)Min()(minNode *Node){
	root := bst.root
	if root == nil{
		return nil
	}
	min := bst._min(root)
	return min
}

func (bst *BST)_min(root *Node)(*Node){
	if root.left == nil{
		return root
	}else{
		return bst._min(root.left)
	}
}

func (bst *BST)Max()(maxNode *Node){
	root := bst.root
	if root == nil{
		return nil
	}
	return bst._max(root)
}

func (bst *BST)_max(root *Node)(maxNode *Node){
	if root.right != nil{
		return bst._max(root.right)
	}else{
		return root
	}
}


// 寻找当前节点为树根的子树的最小节点, 并将最小节点和上一个节点进行断开
func (bst *BST)findMin(currentNode *Node, preNode *Node)(min *Node){
	if currentNode.left != nil{
		minNode := bst.findMin(currentNode.left, currentNode)
		return minNode
	}else{
		preNode.left = nil
		return currentNode
	}
}


// 删除二叉树的一个节点
func (bst *BST)DeleteNode(value string)bool{
	min := bst.Min()
	max := bst.Max()
	if min == nil || value < min.Ele || max == nil && value > max.Ele{
		return false
	}
	root := bst.root
	// 被删除节点是根节点
	if root.Ele == value{
		if root.right != nil{
			// 注意在寻找最小节点的时候将其余其上一个节点断开
			minNode := bst.findMin(root.right, root)
			minNode.left = root.left
			if root.left != nil{
				root.left.pre = minNode
			}
			minNode.right = root.right
			if root.right != nil{
				root.right.pre = minNode
			}
			bst.root = minNode
			bst.root.pre = nil       // 这里遗忘了
		}else if root.left != nil{
			// 上升为根节点， 要将pre设置为nil
			bst.root = root.left
			bst.root.pre = nil
		}else{
			bst.root = nil
		}
		bst.SIZE -= 1
		return true
	}
	if root.Ele > value{
		root.left = bst._deleteNode(root.left, value)
	}else{
		root.right = bst._deleteNode(root.right, value)
	}
	return true
}

func (bst *BST)_deleteNode(node *Node, value string)(*Node){

	// 如果找到要删除的节点
	if node.Ele == value{
		bst.SIZE -= 1
		if node.right != nil{
			minNode := bst.findMin(node.right, node)
			return minNode
		}else if node.left != nil{
			return node.left
		}else{
			return nil
		}
	}else if node.Ele > value{
		node.left = bst._deleteNode(node.left, value)
		if node.left != nil{
			node.left.pre = node
		}
		return node
	}else{
		node.right = bst._deleteNode(node.right, value)
		if node.right != nil{
			node.right.pre = node
		}
		return node
	}

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
	//fmt.Printf("%s,", root.Ele)
	*list = append(*list, root.Ele)
	bst._preOrder(root.left, list)
	bst._preOrder(root.right, list)
}

func (bst *BST)InOrder()(orderList []string){
	// 中序遍历
	orList := make([]string, 0)
	root := bst.root
	bst._inOrder(root, &orList)
	return orList
}

func (bst *BST)_inOrder(root *Node, list *[]string){
	if root == nil{
		return
	}
	bst._inOrder(root.left, list)
	*list = append(*list, root.Ele)
	//fmt.Printf("%s,",root.Ele)
	bst._inOrder(root.right, list)
}

// 二叉树的后续遍历
func (bst *BST)PostOrder()(postOrder []interface{}){
	postOrder = make([]interface{}, 0)
	root := bst.root
	bst._postOrder(root, &postOrder)
	return postOrder
}

func (bst *BST)_postOrder(root *Node, list *[]interface{}){
	if root == nil{
		return
	}
	bst._postOrder(root.left, list)
	bst._postOrder(root.right, list)
	*list = append(*list, root.Ele)
}

// 判断接待是否在二叉树里面
func (bst *BST)IsNodeInTree(value string)(target *Node){
	min := bst.Min()
	max := bst.Max()
	if (min != nil && value < min.Ele) || (max != nil && value > max.Ele){
		return nil
	}
	root := bst.root
	return bst._isNodeInTree(root, value)
}

func (bst *BST)_isNodeInTree(root *Node, value string)(target *Node){
	if root == nil{
		return nil
	}
	if root.Ele > value {
		return bst._isNodeInTree(root.left, value)
	}else if root.Ele < value{
		return bst._isNodeInTree(root.right, value)
	}else{
		return root
	}
}

// 寻找二叉树中序遍历的下一个节点
func (bst *BST)InOrderNextNode(value string)(nextNode *Node){
	currentNode := bst.IsNodeInTree(value)
	if currentNode == nil{
		return nil
	}
	return bst._inOrderNextNode(currentNode)
}

func (bst *BST)_inOrderNextNode(currentNode *Node)(nextNode *Node){

	// 右子树存在，则找右子树最小的节点返回
	if currentNode.right != nil{
		currentNode = currentNode.right
		for ;currentNode.left != nil;{
			currentNode = currentNode.left
		}
		return currentNode
	}else{
		/*以下是这种情况, 当前节点是B
			\
		  	 A
				\
		  	     B   */
		fatherNode := currentNode.pre
		// 注意根节点的pre为nil
		for ; fatherNode !=  nil && fatherNode.left != currentNode;{
			currentNode = fatherNode
			fatherNode = fatherNode.pre
		}
		return fatherNode
	}

}
