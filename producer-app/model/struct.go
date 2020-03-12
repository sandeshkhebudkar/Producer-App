package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Book is DB layer structure
type Book struct {
	gorm.Model
	Name  string
	Price uint
}

//FindWithLimit is use to find records
func FindWithLimit(db *gorm.DB, i, j int) (*gorm.DB, []Book) {

	var books []Book

	db = db.Find(&books, "id>=? AND id<=?", i, j)
	//fmt.Println("DB", books)
	return db, books
}
