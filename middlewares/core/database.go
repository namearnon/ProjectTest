package core

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Name     string
	Host     string
	User     string
	Password string
	Port     string
	Location string
}

func NewDatabase() *Database {
	return &Database{
		Name: "BeerData",
		//Host:     "mariadb",
		Host:     "localhost",
		User:     "user",
		Password: "1234",
		Port:     "3306",
		Location: "Asia%2fBangkok",
	}
}

// ConnectDB to connect database
func (db *Database) Connect() (*gorm.DB, error) {
	newDB, err := gorm.Open("mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&multiStatements=True&loc=%v",
			db.User, db.Password, db.Host, db.Port, db.Name, db.Location,
		))
	newDB.SingularTable(true)
	if err != nil {
		return nil, err
	}
	newDB.LogMode(true)
	return newDB, nil
}

func (db *Database) PingDB() error {
	_, err := db.Connect()
	return err
}
