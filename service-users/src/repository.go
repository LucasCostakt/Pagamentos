package src

import "database/sql"

type RepositoryList interface {
	Insert(query string) (sql.Result, error)
}
