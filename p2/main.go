package main

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseWrapper interface {
	Create(any) error
	Read(int, any) error
	ReadAll(int, any) error
	Update(int, any) error
	Delete(int) error
}

// applyDataToPointer is ripped from the
// std library's JSON Unmarshal function,
// I cracked and had to find out how they
// did it with code. This is the result.
func applyDataToPointer(src, dst any) {

	rv := reflect.ValueOf(src)
	ri := reflect.ValueOf(dst).Elem()
	ri.Set(rv)
}

func main() {

	// Database setup complete
	// sqlite test. init database in memory
	sqldb, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		log.Fatal(err)
	}

	// close connection
	// after function returns.
	defer sqldb.Close()

	if _, err := sqldb.Exec(`
	DROP TABLE IF EXISTS posts;
	DROP TABLE IF EXISTS comments;
	CREATE TABLE posts(id INTEGER PRIMARY KEY, text TEXT,comment_count INT, created_at INT);
	CREATE TABLE comments(id INTEGER PRIMARY KEY,post_id INTEGER,text TEXT,created_at INT);`); err != nil {
		panic(err)
	}

	commentsDB := NewCommentTable(sqldb)
	postTable := NewPostTable(sqldb, commentsDB)

	r := gin.Default()

	posts := r.Group("/posts")
	{
		posts.POST(
			"",
			Add[Post](postTable),
		)

		posts.DELETE(
			"/:id",
			Read[Post](postTable),
		)

		posts.PUT(
			"/:id",
			Update[Post](postTable),
		)

		posts.GET(
			"/:id",
			ReadOne[Post](postTable),
		)

		posts.GET(
			"",
			Read[Post](postTable),
		)

		comments := posts.Group("/comment")
		{
			comments.POST(
				"",
				Add[Comment](commentsDB),
			)
		}
	}

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}

}

func initDB(db DatabaseWrapper) error {

	return nil
}
