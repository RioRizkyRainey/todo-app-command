package internal

import (
	"context"
	"fmt"
	"todo-app/internal/db"
	"todo-app/internal/models"
)

func GetTodos(database *db.Database) {
	todoDB := database.Client.DB(context.TODO(), "todos_rrzky")
	rows, err := todoDB.Query(context.TODO(), "_design/design", "_view/view")

	fmt.Println("My Todo")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("No Data")
		return
	}

	for rows.Next() {
		var id string
		if err := rows.ScanKey(&id); err != nil {
			fmt.Println(err.Error())
		}
		row := todoDB.Get(context.TODO(), id)

		var todo models.Todo
		row.ScanDoc(&todo)
		doneText := "Not Done"
		if todo.IsDone {
			doneText = "Done"
		}
		fmt.Println(" - " + todo.Text)
		fmt.Println("     " + doneText)
		fmt.Println("     created at: " + todo.CreatedAt)
		fmt.Println("     id: " + todo.ID)
	}
}
