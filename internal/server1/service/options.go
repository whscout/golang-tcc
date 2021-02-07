package service

import "golang-tcc/internal/server1/dao"

type Option interface {
	Apply(*Service) error
}

type withDao struct {
	d *dao.Dao
}

func WithDao(d *dao.Dao) *withDao {
	return &withDao{d}
}

func (w *withDao) Apply(s *Service) error {
	s.d = w.d
	return nil
}