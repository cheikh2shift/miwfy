package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseWrapper interface {
	Create(any) error
	Read(int, any) error
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

	db := NewInMemoryDb()

	if err := initDB(db); err != nil {
		panic(err)
	}

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
	DROP TABLE IF EXISTS entries;
	CREATE TABLE entries(id INTEGER PRIMARY KEY, text TEXT, created_at INT);`); err != nil {
		panic(err)
	}

	sqlWrapper := NewSQLite(sqldb)

	if err := initDB(sqlWrapper); err != nil {
		panic(err)
	}

}

func initDB(db DatabaseWrapper) error {

	record := memRecord{Text: "Hello World"}
	record2 := memRecord{Text: "Record 2"}
	if err := db.Create(record); err != nil {
		return err
	}

	if err := db.Create(record2); err != nil {
		return err
	}
	record2.Text = "Hello World 2"

	if err := db.Update(1, record2); err != nil {
		return err
	}

	var queryResult memRecord

	if err := db.Read(1, &queryResult); err != nil {
		return err
	}

	fmt.Println("\n\nQuery Result\n\n", queryResult)

	err := db.Delete(1)

	return err
}
