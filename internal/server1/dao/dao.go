package dao

import (
	"golang-tcc/library/database/mysql"
)

type Dao struct {
	db *mysql.DB
}

func New(opts ...Option) (*Dao, error) {
	dao := &Dao{}
	for _, o := range opts {
		if err := o.Apply(dao); err != nil {
			return nil, err
		}
	}
	return dao, nil
}
