package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "list":
	case "update":
	case "create":
	case "delete":
	default:
		fmt.Println("command not found")
	}

}
