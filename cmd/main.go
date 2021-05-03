package main

import (
	"fmt"
	"os"

	"todo-app/internal"
	"todo-app/internal/db"
)

func main() {
	arg := os.Args[1]

	database, err := db.Initialize()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch arg {
	case "list":
		internal.GetTodos(database)
		break
	case "update":
		id := os.Args[2]
		text := os.Args[3]
		if id == "" {
			fmt.Println("id must included in argument")
			return
		}
		if text == "" {
			fmt.Println("text must included in argument")
			return
		}
		internal.UpdateTodos(database, id, text)
		break
	case "create":
		text := os.Args[2]
		if text == "" {
			fmt.Println("text must included in argument")
			return
		}
		internal.CreateTodos(database, text)
		break
	case "delete":
		id := os.Args[2]
		if id == "" {
			fmt.Println("id must included in argument")
			return
		}
		internal.DeleteTodos(database, id)
	case "mark-done":
		id := os.Args[2]
		internal.DoneTodos(database, id)
		break
	default:
		fmt.Println("command not found")
	}

}
