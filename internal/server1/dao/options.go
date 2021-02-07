package dao

import "github.com/jinzhu/gorm"

type Option interface {
	Apply(*Dao) error
}

type withDB struct {
	db *gorm.DB
}

func WithDB(db *gorm.DB) *withDB {
	return &withDB{db}
}

func (w *withDB) Apply(d *Dao) error {
	d.db = w.db
	return nil
}
