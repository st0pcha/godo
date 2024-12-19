package main

import (
	"fmt"
	"os"

	"github.com/st0pcha/godo/pkg/loader"
)

func main() {
	verifyArgs()

	commands, err := loader.LoadGodoFile()
	if err != nil {
		exitWithMessage("failed to load .godo file")
	}

	fmt.Println(commands)
}

func verifyArgs() {
	if len(os.Args) < 2 {
		exitWithMessage("no command provided")
	}
}

func exitWithMessage(message string) {
	fmt.Println("[GODO]", message)
	os.Exit(1)
}
