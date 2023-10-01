package config

import (
	"fmt"

	"project/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func ConnectDBTest() *gorm.DB {
    config := Config{
        DB_Username: "root",
        DB_Password: "",
        DB_Port:     "3306",
        DB_Host:     "127.0.0.1",
        DB_Name:     "crud_go", 
    }

    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        config.DB_Username,
        config.DB_Password,
        config.DB_Host,
        config.DB_Port,
        config.DB_Name,
    )

    db, err := gorm.Open("mysql", connectionString)
    if err != nil {
        panic(err)
    }

    return db
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "127.0.0.1",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.User{}, &model.Book{})
}
