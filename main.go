package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sandeshkhebudkar/Producer-App/platform/mysql"
	"github.com/sandeshkhebudkar/Producer-App/service"

	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "book"
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

	db, err1 := mysql.Connect()
	if err1 != nil {
		panic("failed to connect database")
	}

	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Book{})
	db = service.Show(db, producer)

	// read command line input
	/* reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		// publish without goroutene
		publish(msg, producer)

		// publish with go routene
		// go publish(msg, producer)
	}
	*/
	/* for i := 1; i < 100; i++ {
		db = add(db, book)
		time.Sleep(2)
	} */
	// Show

}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}
