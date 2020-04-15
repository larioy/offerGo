package queue

import "offerPCom/common"

type Queue struct{
	List []interface{}
	Size int
}

// 初始化一个栈空间
func InitQueue()(q *Queue){
	q = &Queue{}
	q.List = make([]interface{}, 0)
	q.Size = 0
	return q
}

// 栈插入元素
func (q *Queue)Push(ele interface{}){
	if q.Size == 0{
		q.List = append(q.List, ele)
		q.Size += 1
		return
	}
	if !common.IsSameType(ele, q.List[0]){
		panic("Queue insert error ele is not the same type")
	}
	q.List = append(q.List, ele)
	q.Size += 1
}

func (q *Queue)Pop()(ele interface{}){
	if len(q.List) == 0{
		return nil
	}
	ele = q.List[0]
	if len(q.List) == 1 {
		// 这里是进行重新生成比较好吗？
		q.List = make([]interface{}, 0)
	}else {
		q.List = q.List[1:]
	}
	q.Size -= 1
	return ele
}