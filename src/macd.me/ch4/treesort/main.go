package main

import (
	"fmt"
)

// tree 二叉树结构体
type tree struct {
	value       int
	left, right *tree
}

// Sort 使用二叉树进行排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	// 此处的values[:0]应该是将values数组指向0的位置
	appendValues(values[:0], root)
}

// 使用中序遍历读取数据
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 构建二叉树，保证L<D<=R
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	values := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	Sort(values)
	fmt.Println(values)
}
