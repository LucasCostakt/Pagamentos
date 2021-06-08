package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) Insert(query string) (sql.Result, error) {
	res, err := r.db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("Insert().ExecContext(): %w", err)
	}
	return res, nil
}
