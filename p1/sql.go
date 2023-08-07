package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type SQLite struct {
	db *sql.DB
}

func NewSQLite(db *sql.DB) *SQLite {

	return &SQLite{
		db,
	}
}

func (db *SQLite) Create(r any) error {

	now := time.Now()
	entry, ok := r.(memRecord)

	if !ok {
		return errors.New("wrong concrete type passed")
	}

	_, err := db.db.Exec(
		fmt.Sprintf(
			`INSERT INTO entries(text, created_at) VALUES('%s',%d);`,
			entry.Text,
			now.Unix(),
		),
	)

	return err
}

func (db *SQLite) Read(q int, r any) error {

	var text string
	var timestamp int

	row := db.db.QueryRow("SELECT text,created_at FROM entries WHERE id = ?", q)
	err := row.Scan(&text, &timestamp)

	if err != nil {
		return err
	}
	result := memRecord{
		CreatedAt: time.Unix(
			int64(timestamp),
			0,
		),
		Text: text,
	}

	applyDataToPointer(result, r)

	return nil
}

func (db *SQLite) Update(q int, r any) error {
	record := r.(memRecord)

	_, err := db.db.Exec(
		"UPDATE entries SET text = '?' WHERE id = ?;",
		record.Text,
		q,
	)

	return err
}

func (db *SQLite) Delete(q int) error {

	_, err := db.db.Exec("DELETE FROM entries WHERE id = ?;", q)

	return err
}
