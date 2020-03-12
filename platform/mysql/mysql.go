package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

//Connect is use to establish connection
func Connect() (*gorm.DB, error) {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("config") // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	uri := viper.GetString("mysql.URI")
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}
