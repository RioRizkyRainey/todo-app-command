package internal

import (
	"context"
	"fmt"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/models"

	"github.com/google/uuid"
)

func CreateTodos(database *db.Database, text string) {
	todoDB := database.Client.DB(context.TODO(), "todos_rrzky")

	todo := models.Todo{
		ID:        uuid.New().String(),
		Text:      text,
		IsDone:    false,
		CreatedAt: time.Now().String(),
	}

	rev, err := todoDB.Put(context.TODO(), uuid.New().String(), todo)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error save Data")
		return
	}
	todo.Rev = rev

	fmt.Println("Todo created")
}
