package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "help" {
		printHelp()
		return
	}



	fmt.Println("Invalid option. Use \"./make.exe help\" for usage instructions.")
}

func printHelp() {
	fmt.Println("Usage about \"make\":")
	fmt.Println("this program runs command : \"make [option1] [option2]\".")
	fmt.Println("---------------------------------------------------")
	fmt.Println("make help")
	fmt.Println("   - show this command help")
	fmt.Println()
}

/*
追加したいコマンド
pyinit
pyadd
pyremove
pyoutput

gitinit
gitupdate
gitupdate-ac
*/