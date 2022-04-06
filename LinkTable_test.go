package linktable

import (
	"fmt"
	"testing"
)

func TestLinkTable(t *testing.T) {
	lt := CreateLinkTable()
	one := &LinkTableNode{1, nil}
	two := &LinkTableNode{2, nil}
	three := &LinkTableNode{3, nil}
	four := &LinkTableNode{4, nil}
	lt.AddLinkTableNode(one)
	lt.AddLinkTableNode(two)
	lt.AddLinkTableNode(three)
	lt.AddLinkTableNode(four)
	fmt.Println("当前链表：", lt)		// 1 2 3 4
	fmt.Println("链表头部：", lt.GetLinkTableHead())	//1
	lt.DelLinkTableNode(three)	// 1 2 4
	lt.DelLinkTableNode(four)	// 1 2
	fmt.Println("one节点的后继节点：", lt.GetNextLinkTableNode(one))	// 2 
	fmt.Println("查询val为1的节点：", lt.SearchLinkTableNode(func(ltn *LinkTableNode) bool {
		return ltn.Val == 1
	}))		// 1
	lt.DeleteLinkTable()
	fmt.Println(lt)		//LinkTable is null
}

