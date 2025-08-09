package models

import (
	"database/sql"
	"errors"
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
	statement := `INSERT INTO snippets (title, content, created, expires)
	VALUES (?,?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, nil
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// this function returns the 10 most recent entries
func (m *SnippetModel) Latest() ([]snippet, error) {
	return nil, nil
}

// This function  returns a specific snippet based on its id

func (m *SnippetModel) Get(id int) (snippet, error) {
	// to write the stetement to value of a record
	statement := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(statement, id)

	var s snippet
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return snippet{}, sql.ErrNoRows
		} else {
			return snippet{}, err
		}
	}
	return s, nil
}
