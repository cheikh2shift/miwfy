package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type CommentTable struct {
	db *sql.DB
}

func NewCommentTable(db *sql.DB) *CommentTable {

	return &CommentTable{
		db: db,
	}
}

func (db *CommentTable) Create(r any) error {

	now := time.Now()
	entry, ok := r.(Comment)

	if !ok {
		return errors.New("wrong concrete type passed")
	}

	_, err := db.db.Exec(
		"UPDATE posts SET comment_count = comment_count + 1 WHERE id = ?;",
		entry.PostId,
	)

	if err != nil {
		return err
	}

	_, err = db.db.Exec(
		fmt.Sprintf(
			`INSERT INTO comments(text, post_id, created_at) VALUES('%s', %d, %d);`,
			entry.Comment,
			entry.PostId,
			now.Unix(),
		),
	)

	return err
}

func (db *CommentTable) ReadAll(q int, r any) error {

	results := []Comment{}
	rows, err := db.db.Query("SELECT id,text,created_at FROM comments WHERE post_id = ?", q)

	if err != nil {
		return err
	}

	for rows.Next() {

		var text string
		var id, timestamp int

		err = rows.Scan(
			&id,
			&text,
			&timestamp,
		)

		if err != nil {
			return err
		}

		result := Comment{
			Object: Object{
				CreatedAt: time.Unix(
					int64(timestamp),
					0,
				),
				ID: id,
			},
			Comment: text,
		}

		results = append(results, result)
	}

	applyDataToPointer(results, r)

	return nil
}

func (db *CommentTable) Read(q int, r any) error {

	var text string
	var timestamp int

	row := db.db.QueryRow("SELECT text, created_at FROM comments WHERE id = ?", q)
	err := row.Scan(&text, &timestamp)

	if err != nil {
		return err
	}

	result := Post{
		Object: Object{CreatedAt: time.Unix(
			int64(timestamp),
			0,
		)},
		Text: text,
	}

	applyDataToPointer(result, r)

	return nil
}

func (db *CommentTable) Update(q int, r any) error {
	record := r.(Post)

	_, err := db.db.Exec(
		"UPDATE comments SET text = '?' WHERE id = ?;",
		record.Text,
		q,
	)

	return err
}

func (db *CommentTable) Delete(q int) error {

	_, err := db.db.Exec("DELETE FROM comments WHERE id = ?;", q)

	return err
}
