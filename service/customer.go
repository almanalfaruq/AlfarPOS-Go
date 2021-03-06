package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type CustomerService struct {
	customer customerRepositoryIface
}

func NewCustomerService(customerRepo customerRepositoryIface) *CustomerService {
	return &CustomerService{
		customer: customerRepo,
	}
}

func (service *CustomerService) GetAllCustomer() ([]model.Customer, error) {
	return service.customer.FindAll()
}

func (service *CustomerService) GetOneCustomer(id int64) (model.Customer, error) {
	return service.customer.FindById(id)
}

func (service *CustomerService) NewCustomer(customerData string) (model.Customer, error) {
	var customer model.Customer
	customerDataByte := []byte(customerData)
	err := json.Unmarshal(customerDataByte, &customer)
	if err != nil {
		return customer, err
	}
	customer.Code = fmt.Sprintf("CUST-%s", time.Now().Format("20060102150405"))
	return service.customer.New(customer)
}

func (service *CustomerService) UpdateCustomer(customerData string) (model.Customer, error) {
	var customer model.Customer
	customerDataByte := []byte(customerData)
	err := json.Unmarshal(customerDataByte, &customer)
	if err != nil {
		return customer, err
	}
	return service.customer.Update(customer)
}

func (service *CustomerService) DeleteCustomer(id int64) (model.Customer, error) {
	return service.customer.Delete(id)
}
