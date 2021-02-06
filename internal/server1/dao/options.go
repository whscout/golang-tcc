package dao

import (
	"golang-tcc/library/database/mysql"
)

type Option interface {
	Apply(*Dao) error
}

func WithDB(db *mysql.DB) *withDB {
	return &withDB{db}
}

type withDB struct {
	db *mysql.DB
}

func (w *withDB) Apply(d *Dao) error {
	d.db = w.db
	return nil
}
