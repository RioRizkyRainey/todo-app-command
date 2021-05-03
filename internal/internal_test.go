package internal_test

import (
	"testing"
	"time"
	"todo-app/internal"
	"todo-app/internal/db"
	"todo-app/internal/models"

	"github.com/go-kivik/kivik/driver"
	"github.com/go-kivik/kivikmock"
	"github.com/google/uuid"
)

func TestGetTodos(t *testing.T) {
	database, _ := db.MockInit()

	rows := kivikmock.NewRows().
		AddRow(&driver.Row{
			ID:    "1792e805-a2b3-d775-5aa1-787ccbab9e37",
			Key:   []byte(`1792e805-a2b3-d775-5aa1-787ccbab9e37`),
			Value: []byte(`1`),
			Doc: []byte(`{
				"_id":       "1792e805-a2b3-d775-5aa1-787ccbab9e37",
				"_rev":      "3-de639474ea0727b49858cea7b48dae19",
				"text":      "",
				"is_done":   true,
				"createdAt": "",
			}`),
		})

	database.Mock.NewDB().
		ExpectQuery().
		WithView("_view/view").
		WithDDocID("_design/design").
		WillReturn(rows)

	internal.GetTodos(database)
}

func TestCreateTodos(t *testing.T) {
	database, _ := db.MockInit()

	uuid := uuid.New().String()

	todo := models.Todo{
		ID:        uuid,
		Text:      "test",
		IsDone:    false,
		CreatedAt: time.Now().String(),
	}
	database.Mock.NewDB().
		ExpectPut().
		WithDoc(todo).
		WithDocID(uuid).
		WillReturn("rev_id")

	internal.CreateTodos(database, "test")
}

func TestUpdateTodos(t *testing.T) {
	database, _ := db.MockInit()

	uuid := uuid.New().String()

	todo := models.Todo{
		ID:        uuid,
		Text:      "test",
		IsDone:    false,
		CreatedAt: time.Now().String(),
	}
	database.Mock.NewDB().
		ExpectPut().
		WithDoc(todo).
		WithDocID(uuid).
		WillReturn("rev_id")

	internal.UpdateTodos(database, uuid, "test")
}

func TestDeleteTodos(t *testing.T) {
	database, _ := db.MockInit()

	uuid := uuid.New().String()

	todo := models.Todo{
		ID:        uuid,
		Text:      "test",
		IsDone:    false,
		CreatedAt: time.Now().String(),
	}
	database.Mock.NewDB().
		ExpectPut().
		WithDoc(todo).
		WithDocID(uuid).
		WillReturn("rev_id")

	internal.DeleteTodos(database, uuid)
}
func TestDoneTodos(t *testing.T) {
	database, _ := db.MockInit()

	uuid := uuid.New().String()

	todo := models.Todo{
		ID:        uuid,
		Text:      "test",
		IsDone:    true,
		CreatedAt: time.Now().String(),
	}
	database.Mock.NewDB().
		ExpectPut().
		WithDoc(todo).
		WithDocID(uuid).
		WillReturn("rev_id")

	internal.DoneTodos(database, uuid)
}
