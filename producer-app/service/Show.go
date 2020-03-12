package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sandeshkhebudkar/Producer-App/producer-app/model"
)

//Show is used to create book
func Show(db *gorm.DB) *gorm.DB {

	var total int
	var response []model.Book
	var Midobj []model.Book
	db.Table("books").Count(&total)
	fmt.Println(total)
	for i := 1; i <= total; i = i + 10 {

		db, Midobj = model.FindWithLimit(db, i, i+9)
		response = append(response, Midobj...)
	}

	fmt.Println(response)

	return db
}
