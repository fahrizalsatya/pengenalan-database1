package database

import (
	"database/sql"
	"log"
)

//Customer represent (ID, FirstName, LastName, NpwpID, Age, CustomerType, Street, City, State, ZipCode, PhoneNumber)
type Customer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	NpwpID       string `json:"npwp_id"`
	Age          int    `json:"age"`
	CustomerType string `json:"customer_type"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	PhoneNumber  string `json:"phone_number"`
}

//InsertCustomer for insert new data to database
func InsertCustomer(customer Customer, db *sql.DB) {
	_, err := db.Exec("insert into customers(first_name, last_name,npwp_id,age,customer_type,street,city,state,zip_code,phone_number) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		customer.FirstName,
		customer.LastName,
		customer.NpwpID,
		customer.Age,
		customer.CustomerType,
		customer.Street,
		customer.City,
		customer.State,
		customer.ZipCode,
		customer.PhoneNumber)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("insert success!")
}
