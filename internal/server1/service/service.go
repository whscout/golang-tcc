package service

import (
	"golang-tcc/config"
	"golang-tcc/internal/server1/dao"
)

type Service struct {
	conf *config.Config
	d    *dao.Dao
}

func New(opts ...Option) (*Service, error) {
	s := &Service{}
	for _, o := range opts {
		if err := o.Apply(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}
