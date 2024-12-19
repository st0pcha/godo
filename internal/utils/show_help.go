package utils

import "fmt"

func ShowHelp() {
	fmt.Println("[GODO] Correct usage")
	fmt.Println("- godo --init | create a .godo file")
	fmt.Println("- godo --help | show this message")
	fmt.Println("- godo <command> | execute a script")
}
