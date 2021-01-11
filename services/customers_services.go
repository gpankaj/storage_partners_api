package services

import (
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/domains/customers_domains"

	)

var (
	CustomerService customerServiceInterface = &customerService{}
)

type customerService struct {
}

type customerServiceInterface interface {
	Create_Customer(customers_domains.Customer) (*customers_domains.Customer, *rest_errors_package.RestErr)
	Get_Customer(customer_id int64) (*customers_domains.Customer, *rest_errors_package.RestErr)
	Update_Customer(bool,customers_domains.Customer) (*customers_domains.Customer, *rest_errors_package.RestErr)
}

func (c *customerService) Create_Customer(customer_domain customers_domains.Customer) (*customers_domains.Customer, *rest_errors_package.RestErr) {
	return nil,nil
}

func (c *customerService) Get_Customer(customer_id int64) (*customers_domains.Customer, *rest_errors_package.RestErr) {
	return nil,nil
}

func (c *customerService) Update_Customer(bool,customers_domains.Customer) (*customers_domains.Customer, *rest_errors_package.RestErr) {
	return nil,nil
}
