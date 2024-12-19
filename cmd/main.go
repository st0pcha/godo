package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/st0pcha/godo/internal/executor"
	"github.com/st0pcha/godo/internal/loader"
	"github.com/st0pcha/godo/internal/utils"
)

func main() {
	helpFlag := flag.Bool("help", false, "show help message")
	initFlag := flag.Bool("init", false, "create a .godo file")
	flag.Parse()

	if *helpFlag {
		utils.ShowHelp()
		os.Exit(1)
	}
	if *initFlag {
		if err := loader.CreateGodoFile(); err != nil {
			exitWithMessage(err.Error())
		}
		os.Exit(1)
	}

	verifyArgs()

	commands, err := loader.LoadGodoFile()
	if err != nil {
		exitWithMessage(err.Error())
	}

	script, err := loader.FindCommand(commands, os.Args[1])
	if err != nil {
		exitWithMessage(err.Error())
	}

	if err := executor.Execute(script); err != nil {
		exitWithMessage("failed to execute command")
	}
}

func verifyArgs() {
	if len(os.Args) < 2 {
		exitWithMessage("no command provided. use 'godo --help' for more information")
	}
}

func exitWithMessage(message string) {
	fmt.Println("[GODO]", message)
	os.Exit(1)
}
