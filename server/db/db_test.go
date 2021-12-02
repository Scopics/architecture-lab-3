package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection := &Connection{
		Host:       "localhost",
		Port:       "5432",
		User:       "postgres",
		Password:   "password",
		DbName:     "restaurant",
		DisableSSL: false,
	}

	expectedConnectionString := "host=localhost port=5432 user=postgres password=password dbname=restaurant sslmode=disable"

	assert.Equal(t, connection.GetConnectionString(), expectedConnectionString)
}

func TestPing(t *testing.T) {
	connection := &Connection{
		Host:       "localhost",
		Port:       "5432",
		User:       "postgres",
		Password:   "password",
		DbName:     "restaurant",
		DisableSSL: false,
	}

	db, err := connection.Open()

	if err != nil {
		t.Error("Unable to connect to database")
	}

	err = db.Ping()
	if err != nil {
		t.Error("No response from server")
	}
}
