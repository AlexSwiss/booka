package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func FetchConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:rosecransB430@/bookworm")
	if err != nil {
		panic(err)
	}
	return db
}

type Author struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Bookid    int    `json:"bookid"`
}

type Book struct {
	ID       int       `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Authors  []*Author `json:"authors"`
}
