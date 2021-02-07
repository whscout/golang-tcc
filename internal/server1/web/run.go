package web

import (
	"go.uber.org/zap"
	"golang-tcc/config"
)

type server struct {
	l      *zap.Logger
	config *config.Config
}

func NewServer(c *config.Config) error {
	return nil
}
