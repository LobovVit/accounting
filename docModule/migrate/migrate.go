package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	m, err := migrate.New(
		"file:///pg",
		"postgres://localhost:5433/accounting")
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
	m.Steps(2)
}
