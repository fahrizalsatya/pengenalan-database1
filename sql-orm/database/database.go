package database

import (
	"log"

	"gorm.io/gorm"
)

//CustomerORM represents (ID, FirstName, LastName, NpwpID, Age, CustomerType, Street, City, State, ZipCode, PhoneNumber, AccountORM)
type CustomerORM struct {
	ID           int          `gorm:"primary_key" json:"-"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	NpwpID       string       `json:"npwp_id"`
	Age          int          `json:"age"`
	CustomerType string       `json:"customer_type"`
	Street       string       `json:"street"`
	City         string       `json:"city"`
	State        string       `json:"state"`
	ZipCode      string       `json:"zip_code"`
	PhoneNumber  string       `json:"phone_number"`
	AccountORM   []AccountORM `gorm:"ForeignKey:IDCustomerRefer";json:"account_orm"`
}

//AccountORM represents (ID, IdCustomerRefer, Balance, AccounType)
type AccountORM struct {
	ID              int    `gorm:"primary_key" json:"-"`
	IDCustomerRefer int    `json:"-"`
	Balance         int    `json:"balance"`
	AccountType     string `json:"account_type"`
}

//InsertCustomer add data to database, using gorm
func InsertCustomer(customer CustomerORM, db *gorm.DB) {
	if err := db.Create(&customer).Error; err != nil {
		log.Println("failed to insert :", err.Error())
		return
	}
	log.Println("Success insert data")
}

//GetCustomers show all data from database customer, using gorm
func GetCustomers(db *gorm.DB) {
	var customer []CustomerORM
	if err := db.Preload("AccountORM").Find(&customer).Error; err != nil {
		log.Println("failed to get data :", err.Error())
		return
	}
	log.Println(customer)
}

//DeleteCustomer delete data based on their customer_id, using gorm
func DeleteCustomer(id int, db *gorm.DB) {
	var customer CustomerORM
	if err := db.Where(&CustomerORM{ID: id}).Delete(&customer).Error; err != nil {
		log.Println("failed to delete data :", err.Error())
		return
	}

	log.Println("Success delete data")
}

//UpdateCustomer update data based on their customer_id, using gorm
func UpdateCustomer(customer CustomerORM, id int, db *gorm.DB) {
	if err := db.Model(&CustomerORM{}).Where(&CustomerORM{ID: id}).Updates(customer).Error; err != nil {
		log.Println("failed to update data :", err.Error())
		return
	}

	log.Println("Success update data")
}
