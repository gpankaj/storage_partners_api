package services

import (
	"github.com/gpankaj/storage_partners_api/domains/customers_domains"
	"github.com/gpankaj/storage_partners_api/utils/errors"
)

var (
	CustomerService customerServiceInterface = &customerService{}
)

type customerService struct {
}

type customerServiceInterface interface {
	Create_Customer(customers_domains.Customer) (*customers_domains.Customer, *errors.RestErr)
	Get_Customer(customer_id int64) (*customers_domains.Customer, *errors.RestErr)
	Update_Customer(bool,customers_domains.Customer) (*customers_domains.Customer, *errors.RestErr)
}

func (c *customerService) Create_Customer(customer_domain customers_domains.Customer) (*customers_domains.Customer, *errors.RestErr) {
	return nil,nil
}

func (c *customerService) Get_Customer(customer_id int64) (*customers_domains.Customer, *errors.RestErr) {
	return nil,nil
}

func (c *customerService) Update_Customer(bool,customers_domains.Customer) (*customers_domains.Customer, *errors.RestErr) {
	return nil,nil
}
