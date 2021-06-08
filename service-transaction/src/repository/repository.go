package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) Insert(query string) (sql.Result, error) {
	res, err := r.db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("Insert().Exec(): %w", err)
	}
	return res, nil
}

func (r *Repository) Update(query string) (sql.Result, error) {
	res, err := r.db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("Insert().Exec(): %w", err)
	}
	return res, nil
}

func (r *Repository) Select(query string) (*sql.Rows, error) {
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Select().Query(): %w", err)
	}
	return rows, nil
}
