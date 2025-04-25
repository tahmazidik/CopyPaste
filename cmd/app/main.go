package main

import (
	"fmt"

	"github.com/tahmazidik/Copy_Paste/internal/editor"
)

func main() {
	result := editor.Check("Hello, World!")
	fmt.Println("Result:", result)
}
