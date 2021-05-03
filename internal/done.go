package internal

import (
	"context"
	"fmt"
	"todo-app/internal/db"
	"todo-app/internal/models"
)

func DoneTodos(database *db.Database, id string) {
	todoDB := database.Client.DB(context.TODO(), "todos_rrzky")

	var todo models.Todo
	row := todoDB.Get(context.TODO(), id)

	row.ScanDoc(&todo)

	todo = models.Todo{
		Rev:       todo.Rev,
		ID:        id,
		IsDone:    true,
		CreatedAt: todo.CreatedAt,
		Text:      todo.Text,
	}

	rev, err := todoDB.Put(context.TODO(), id, todo)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error save Data")
		return
	}
	todo.Rev = rev

	fmt.Println("Todo updated")
}
