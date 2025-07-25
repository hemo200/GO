package models

import "time"

type  snippet  struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type  SnippetModel  struct {
	DB *sql.DB
}
//this function enters a new value to the Database
func (m *SnippetModel) Insert (title string, content  string, expires int) (int, error){
	return 0, nil
}