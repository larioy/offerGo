package stack

import (
	"offerPCom/common"
)

type Stack struct{
	List []interface{}
	Size int
}

// 初始化一个栈空间
func InitStack()(st *Stack){
	st = &Stack{}
	st.List = make([]interface{}, 0)
	st.Size = 0
	return st
}

// 栈插入元素
func (st *Stack)Push(ele interface{}){
	if st.Size == 0{
		st.List = append(st.List, ele)
		st.Size += 1
		return
	}
	if !common.IsSameType(ele, st.List[0]){
		panic("stack insert error ele is not the same type")
	}
	// 头部插入
	st.List = append([]interface{}{ele}, st.List...)
	st.Size += 1
}

func (st *Stack)Pop()(ele interface{}){
	if st.Size == 0{
		return nil
	}
	ele = st.List[0]
	if st.Size == 1 {
		// 这里是进行重新生成比较好吗？
		st.List = make([]interface{}, 0)
	}else {
		st.List = st.List[1:]
	}
	st.Size -= 1
	return ele
}
