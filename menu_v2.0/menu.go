package main

import (
	"fmt"
	"os"
)

var head *DataNode
var cmds []*DataNode = []*DataNode{
	{"help", "this is help cmd!", Help, nil},
	{"version", "menu program v2.0", ShowVersion, nil},
	{"q", "exit menu", Exit, nil},
}

func main() {
	head = &DataNode{}
	tail := head
	for _, cmd := range cmds {
		tail.next = cmd
		tail = tail.next
	}
	tail.next = nil
	head = head.next
	var cmd string
	for {
		fmt.Print("Input a cmd > ")
		fmt.Scanln(&cmd)
		p := FindCmd(head, cmd)
		if p == nil {
			fmt.Println("This is a wrong cmd!")
			continue
		}
		fmt.Println(p.cmd, " - ", p.desc)
		if p.handler != nil {
			p.handler()
		}
		fmt.Println()
	}
}

func Help() {
	ShowAllCmd(head)
}
func ShowVersion() {
	fmt.Println("menu v2.0")
}
func Exit() {
	os.Exit(0)
}
