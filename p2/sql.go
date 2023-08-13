package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type PostTable struct {
	db           *sql.DB
	commentTable DatabaseWrapper
}

type Object struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	Object
	Text         string    `json:"text"`
	CommentCount int       `json:"comment_count"`
	Comments     []Comment `json:"comments,omitempty"`
}

type Comment struct {
	Object
	Comment string `json:"comment"`
	PostId  int    `json:"post_id,omitempty"`
}

func NewPostTable(
	db *sql.DB,
	comments DatabaseWrapper,
) *PostTable {

	return &PostTable{
		db,
		comments,
	}
}

func (db *PostTable) Create(r any) error {

	now := time.Now()
	entry, ok := r.(Post)

	if !ok {
		return errors.New("wrong concrete type passed")
	}

	_, err := db.db.Exec(
		fmt.Sprintf(
			`INSERT INTO posts(text, comment_count, created_at) VALUES('%s', 0, %d);`,
			entry.Text,
			now.Unix(),
		),
	)

	return err
}

func (db *PostTable) ReadAll(q int, r any) error {

	results := []Post{}
	rows, err := db.db.Query("SELECT id,text,comment_count,created_at FROM posts")

	if err != nil {
		return err
	}

	for rows.Next() {

		var text string
		var id, timestamp, comments int

		err = rows.Scan(
			&id,
			&text,
			&comments,
			&timestamp,
		)

		if err != nil {
			return err
		}

		result := Post{
			Object: Object{
				CreatedAt: time.Unix(
					int64(timestamp),
					0,
				),
				ID: id,
			},
			Text:         text,
			CommentCount: comments,
		}

		results = append(results, result)
	}

	applyDataToPointer(results, r)

	return nil
}

func (db *PostTable) Read(q int, r any) error {

	var text string
	var timestamp,
		comment_count, id int
	var comments []Comment

	row := db.db.QueryRow("SELECT id,text,comment_count, created_at FROM posts WHERE id = ?", q)
	err := row.Scan(&id, &text, &comment_count, &timestamp)

	if err != nil {
		return err
	}

	result := Post{
		Object: Object{
			CreatedAt: time.Unix(
				int64(timestamp),
				0,
			),
			ID: id,
		},
		Text:         text,
		CommentCount: comment_count,
	}

	err = db.commentTable.ReadAll(q, &comments)

	if err != nil {
		return err
	}

	result.Comments = comments
	applyDataToPointer(result, r)

	return nil
}

func (db *PostTable) Update(q int, r any) error {
	record := r.(Post)

	_, err := db.db.Exec(
		"UPDATE posts SET text = '?' WHERE id = ?;",
		record.Text,
		q,
	)

	return err
}

func (db *PostTable) Delete(q int) error {

	_, err := db.db.Exec("DELETE FROM posts WHERE id = ?;", q)

	return err
}
