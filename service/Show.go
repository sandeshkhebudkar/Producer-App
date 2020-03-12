package service

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sandeshkhebudkar/Producer-App/model"
	"github.com/sandeshkhebudkar/Producer-App/platform/kafka"
)

//Show is used to create book
func Show(db *gorm.DB, producer sarama.SyncProducer) *gorm.DB {

	var total int
	var response []model.Book
	var Midobj []model.Book
	db.Table("books").Count(&total)
	fmt.Println(total)
	for i := 1; i <= total; i = i + 10 {
		db, Midobj = model.FindWithLimit(db, i, i+9)
		response = append(response, Midobj...)
	}
	for _, data := range response {
		d := []byte(fmt.Sprintf("%v", data))

		//	d := fmt.Sprintf("%v", data)
		//fmt.Printf("%+v", string(d))

		fmt.Println(string(d))
		kafka.Publish(d, producer)
	}
	return db
}
