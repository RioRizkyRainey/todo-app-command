package db

import (
	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivikmock"
)

type Database struct {
	Client *kivik.Client
	Mock   *kivikmock.MockClient
}

func Initialize() (*Database, error) {
	client, err := kivik.New("couch", "http://admin:iniadmin@13.250.43.79:5984")
	if err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}
