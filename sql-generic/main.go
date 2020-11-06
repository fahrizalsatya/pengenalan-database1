package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FahrizalSatya/pengenalan-database1/sql-generic/config"
	"github.com/FahrizalSatya/pengenalan-database1/sql-generic/database"

	"github.com/spf13/viper"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := connect(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}
	database.InsertCustomer(database.Customer{
		FirstName:    "Fahrizal",
		LastName:     "Satya",
		NpwpID:       "id-1",
		Age:          10,
		CustomerType: "Premium",
		Street:       "Str",
		City:         "Jakarta",
		State:        "Indo",
		ZipCode:      "55555",
		PhoneNumber:  "0812384",
	}, db)
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

func connect(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Config))
	if err != nil {
		return nil, err
	}

	log.Println("db successfully connected")
	return db, nil
}
