package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Connection struct {
	Host       string
	Port       string
	User       string
	Password   string
	DbName     string
	DisableSSL bool
}

func (c *Connection) GetConnectionString() string {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DbName)

	return connectionString
}

func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("postgres", c.GetConnectionString())
}
