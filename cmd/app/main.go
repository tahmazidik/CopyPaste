package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tahmazidik/Copy_Paste/internal/editor"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: copypaste <input> <commands> <output>")
		os.Exit(1)
	}

	e := editor.NewEditor()

	if err := e.LoadFile(os.Args[1]); err != nil {
		fmt.Printf("Error loading file: %v\n", err)
		os.Exit(1)
	}

	processCommands(os.Args[2], e)

	if err := e.SaveFile(os.Args[3]); err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		os.Exit(1)
	}
}

func processCommands(filename string, e *editor.Editor) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening commands file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e.ProcessCommand(scanner.Text())
	}
}
