package src

import "database/sql"

type RepositoryList interface {
	Insert(query string) (sql.Result, error)
	Update(query string) (sql.Result, error)
	Select(query string) (*sql.Rows, error)
}
