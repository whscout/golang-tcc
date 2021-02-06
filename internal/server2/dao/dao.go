package dao

import (
	"golang-tcc/library/database/mysql"
)

type Dao struct {
	db *mysql.DB
}
