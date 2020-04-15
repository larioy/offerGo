package common

import (
	"reflect"
)


// 用于判断元素的类型是否一致
func IsSameType(ori interface{}, target interface{})(bool){
	//fmt.Println(reflect.ValueOf(ori).Type(), reflect.ValueOf(target).Type())
	if reflect.ValueOf(ori).Type() == reflect.ValueOf(target).Type(){
		return true
	}else{
		return false
	}
}

// 切片插入元素操作
func SliceInsert(index int, value interface{}, arr []interface{})([]interface{}){
	if len(arr) > 0{
		if !IsSameType(value, arr[0]){
			panic("data type not same")
		}
	}
	res := append([]interface{}{value}, arr[:index]...)
	return append(res, arr[index:]...)
}

// 获取元素在数组中的下标
func GetIndex(target interface{}, arr []interface{})(index int){
	// 边界条件检查
	if len(arr) == 0{
		return -1
	}
	if !IsSameType(target, arr[0]){
		return -1
	}
	for index, value := range arr{
		if value == target{
			return index
		}
	}
	return -1
}
