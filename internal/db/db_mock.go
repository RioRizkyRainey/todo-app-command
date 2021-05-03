package db

import (
	"github.com/go-kivik/kivikmock"
)

func MockInit() (*Database, error) {
	client, mock, err := kivikmock.New()
	if err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
		Mock:   mock,
	}, nil

}
