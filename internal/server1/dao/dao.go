package dao

import (
	"golang-tcc/library/database/mysql"
)

type Dao struct {
	db *mysql.DB
}

func New(opts ...Option) (*Dao, error) {
	d := &Dao{}
	for _, opt := range opts {
		if err := opt.Apply(d); err != nil {
			return nil, err
		}
	}
	return d, nil
}
