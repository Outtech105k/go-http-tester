package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlEngine = "sqlite3"
	dbFile    = "./chat.db"
)

func initDb() error {
	DbConnection, err := sql.Open(sqlEngine, dbFile)
	if err != nil {
		return err
	}
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS posts(id INTEGER PRIMARY KEY, body TEXT NOT NULL, posted TEXT DEFAULT CURRENT_TIMESTAMP)`
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		return err
	}

	return nil
}

func addPost(body string) error {
	DbConnection, err := sql.Open(sqlEngine, dbFile)
	if err != nil {
		return err
	}
	defer DbConnection.Close()

	cmd := `INSERT INTO posts(body) VALUES(?)`
	_, err = DbConnection.Exec(cmd, body)
	if err != nil {
		return err
	}

	return nil
}

func getPosts() ([]Post, error) {
	DbConnection, err := sql.Open(sqlEngine, dbFile)
	if err != nil {
		return nil, err
	}
	defer DbConnection.Close()

	cmd := `SELECT * FROM posts`
	rows, err := DbConnection.Query(cmd)
	if err != nil {
		return nil, err
	}

	var gotPost []Post
	for rows.Next() {
		r := Post{}
		rows.Scan(&r.Id, &r.Body, &r.Posted)
		gotPost = append(gotPost, r)
	}

	return gotPost, nil
}

type Post struct {
	Id     int    `json:"id"`
	Body   string `json:"body"`
	Posted string `json:"posted"`
}
