package DS

import (
	"testing"
)

/**
输入n个节点，返回根节点
测试用例:
ex1:
in:
5
88 70 61 96 120
out:
70

ex2:
in:
7
88 70 61 96 120 90 65
out:
88
*/

// 测试AVL
func TestAvl(t *testing.T) {
	avl := Avl{}
	arr := []int{88, 70, 61, 96, 120}
	for _, v := range arr {
		avl.Insert(v)
	}
	if avl.root.val != 70 {
		t.Error("错误")
	}
	arr = []int{88, 70, 61, 96, 120, 90, 65}
	for _, v := range arr {
		avl.Insert(v)
	}
	if avl.root.val != 88 {
		t.Error("错误")
	}
}
