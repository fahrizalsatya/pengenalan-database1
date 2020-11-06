package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/FahrizalSatya/pengenalan-database1/sql-generic/config"

	"github.com/spf13/viper"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}
	_, err = connect(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}
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
	if err != nil{
		reutrn nil, err
	}

	log.Println("db successfully connected")
	return db, nil
}
