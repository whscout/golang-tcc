package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type DB = gorm.DB

func NewDB(c Config) (*DB, error) {
	dbURI := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC",
		c.Username, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}
	db.LogMode(c.Verbose)
	db.SingularTable(c.SingTable)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(5)
	return db, err
}
