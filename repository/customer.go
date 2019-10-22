package repository

import (
	"../model"
	"../util"
)

type CustomerRepository struct {
	util.IDatabaseConnection
}

type ICustomerRepository interface {
	FindAll() []model.Customer
	FindById(id int) model.Customer
	New(customer model.Customer) model.Customer
	Update(customer model.Customer) model.Customer
	Delete(id int) model.Customer
}

func (repo *CustomerRepository) FindAll() []model.Customer {
	var customers []model.Customer
	db := repo.GetDb()
	db.Find(&customers)
	return customers
}

func (repo *CustomerRepository) FindById(id int) model.Customer {
	var customer model.Customer
	db := repo.GetDb()
	db.Where("id = ?", id).First(&customer)
	return customer
}

func (repo *CustomerRepository) New(customer model.Customer) model.Customer {
	db := repo.GetDb()
	isNotExists := db.NewRecord(customer)
	if isNotExists {
		db.Create(&customer)
	}
	return customer
}

func (repo *CustomerRepository) Update(customer model.Customer) model.Customer {
	var oldCustomer model.Customer
	db := repo.GetDb()
	db.Where("id = ?", customer.ID).First(&oldCustomer)
	oldCustomer = customer
	db.Save(&oldCustomer)
	return customer
}

func (repo *CustomerRepository) Delete(id int) model.Customer {
	var customer model.Customer
	db := repo.GetDb()
	db.Where("id = ?", id).First(&customer)
	db.Delete(&customer)
	return customer
}
