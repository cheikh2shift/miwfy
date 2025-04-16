package main

import (
	"context"
	"log"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

type User struct {
	Id    int    `ksql:"id"` // Primary key for User
	Name  string `ksql:"name"`
	Tasks []Task // field tag omitted. Field will be populated with query
}

type Task struct {
	Id          int    `ksql:"id"`          // Primary key for Task
	Title       string `ksql:"title"`       // Bind to SQL column "title"
	Description string `ksql:"description"` // Bind to SQL column "description"
	UserId      int    `ksql:"user_id"`     // Bind to SQL column "user_id"
}

func main() {
	connStr := "postgres://postgres:postgres@localhost:5433/individuals?sslmode=disable"
	ctx := context.Background()

	db, err := kpgx.New(ctx, connStr, ksql.Config{})
	if err != nil {
		log.Fatalf("unable connect to database: %s", err)
	}
	defer db.Close()

	_, err = db.Exec(ctx, `
		CREATE TABLE if not exists users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);

		CREATE TABLE if not exists tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
		);
	`)

	if err != nil {
		log.Fatalf("error creating tables: %s", err)
	}

	users := ksql.NewTable("users", "id")

	u := User{
		Name: "Cheikh",
	}

	// passing as pointer will result in id field being populated
	err = db.Insert(ctx, users, &u)

	if err != nil {
		log.Fatalf("error inserting user record: %s", err)
	}

	log.Printf("Inserted user: %+v", u)

	// construct instance of ksql.Table
	tasks := ksql.NewTable("tasks", "id")
	// insert task
	err = db.Insert(ctx, tasks, &Task{
		Title:       "1st Task",
		Description: "Test of relational capabilities",
		UserId:      u.Id,
	})

	if err != nil {
		log.Fatalf("error inserting task record: %s", err)
	}

	// get tasks for newly created user
	err = db.Query(ctx, &u.Tasks, "FROM tasks WHERE user_id = $1", u.Id)

	if err != nil {
		log.Fatalf("unable to query user tasks: %s", err)
	}

	log.Printf("user with posts: %+v\n", u)

}
