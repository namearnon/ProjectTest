package core

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type IContext interface {
	echo.Context
	DB() (*gorm.DB, error)
	PingDB() error
}

type Context struct {
	echo.Context
}

func (c *Context) DB() (*gorm.DB, error) {
	db, err := NewDatabase().Connect()
	if err != nil {
		return nil, err
	}
	db.LogMode(false)
	return db, nil
}

func (c *Context) PingDB() error {
	err := NewDatabase().PingDB()
	if err != nil {
		return err
	}

	return nil
}
