package main

import "fmt"

type DataNode struct{
	cmd string
	desc string
	handler func()
	next *DataNode
}


func FindCmd(head *DataNode, cmd string) *DataNode {
	if head == nil || len(cmd) == 0 {
		return nil
	}
	for p := head; p != nil; p = p.next {
		if p.cmd == cmd {
			return p
		}
	}
	return nil
}

func ShowAllCmd(head *DataNode) {
	fmt.Println("Menu List:")
	for p := head; p != nil; p = p.next {
		fmt.Println(p.cmd, " - ", p.desc)
	}
}
