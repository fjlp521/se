package main

import "fmt"

func main() {
	var cmd string
	for {
		fmt.Print("请输入命令(Q退出)：")
		n, _ := fmt.Scanln(&cmd)
		if n == 0 {
			continue
		}
		if cmd == "Q" || cmd == "q" {
			break
		}
		switch cmd {
		case "help":
			fmt.Println("do something about help")
		case "list":
			fmt.Println("do somethis about list")
		default:
			fmt.Println("sorry, The command \"" + cmd + "\" is not defined")
		}
	}
}
