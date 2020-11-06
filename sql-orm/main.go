package main

import (
	"fmt"
	"log"

	"github.com/FahrizalSatya/pengenalan-database1/sql-orm/config"
	"github.com/FahrizalSatya/pengenalan-database1/sql-orm/database"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	// database.InsertCustomer(database.CustomerORM{
	// 	FirstName:    "Fahrizal",
	// 	LastName:     "Satya",
	// 	NpwpID:       "id-1",
	// 	Age:          10,
	// 	CustomerType: "Premium",
	// 	Street:       "Str",
	// 	City:         "Jakarta",
	// 	State:        "Indo",
	// 	ZipCode:      "55555",
	// 	PhoneNumber:  "0812384",
	// }, db)
	database.GetCustomers(db)
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("db successfully connected")

	return db, nil
}
