package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sandeshkhebudkar/Producer-App/producer-app/database"
	"github.com/sandeshkhebudkar/Producer-App/producer-app/service"
)

//Book is used to store info of book
type Book struct {
	gorm.Model
	Name  string
	Price uint
}

func add(db *gorm.DB, book Book) *gorm.DB {
	db = db.Create(&book)
	return db
}
func main() {

	db, err := database.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//var books []Book
	// Migrate the schema
	db.AutoMigrate(&Book{})

	var book Book

	book.Name = "xyz"
	book.Price = 300

	/* for i := 1; i < 100; i++ {
		db = add(db, book)
		time.Sleep(2)
	} */
	// Show
	db = service.Show(db)

}
