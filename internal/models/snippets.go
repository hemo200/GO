package models

import (
	"database/sql"
	"time"
)

type snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// this function enters a new value to the Database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// This function  returns a specific snippet based on its id
func (m *SnippetModel) Get(id int) (snippet, error) {
	return snippet{}, nil
}

// this function returns the 10 most recent entries
func (m *SnippetModel) Latest() ([]snippet, error) {
	return nil, nil
}
