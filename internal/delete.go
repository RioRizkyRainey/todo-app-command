package internal

import (
	"context"
	"fmt"
	"todo-app/internal/db"
	"todo-app/internal/models"
)

func DeleteTodos(database *db.Database, id string) {
	todoDB := database.Client.DB(context.TODO(), "todos_rrzky")

	row := todoDB.Get(context.TODO(), id)

	var todo models.Todo
	row.ScanDoc(&todo)

	rev, err := todoDB.Delete(context.TODO(), todo.ID, todo.Rev)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error save Data")
		return
	}
	todo.Rev = rev

	fmt.Println("Todo deleted")
}
