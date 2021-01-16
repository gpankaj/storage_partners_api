package services

import (
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/domains/customers_domains"
	"github.com/gpankaj/storage_partners_api/utils/crypto_utils"
)

var (
	CustomerService customerServiceInterface = &customerService{}
)

type customerService struct {
}

type customerServiceInterface interface {
	Create_Customer_Service(customers_domains.Customer) (*customers_domains.Customer, *rest_errors_package.RestErr)
	//Get_Customer(customer_id int64) (*customers_domains.Customer, *rest_errors_package.RestErr)
	//Update_Customer(bool,customers_domains.Customer) (*customers_domains.Customer, *rest_errors_package.RestErr)

	LoginUser(request customers_domains.CustomerLoginRequest) (*customers_domains.Customer,*rest_errors_package.RestErr)
}
//Create_Customer_Service

func (s *customerService) Create_Customer_Service(customer customers_domains.Customer) (*customers_domains.Customer, *rest_errors_package.RestErr) {
	/*
		partner_model := partners_domains.Partner{}
		partner_model.Email_id = strings.TrimSpace(strings.ToLower(partner.Email_id))
	*/

	if err:= customer.Validate(false); err!=nil{
		return nil, err
	}

	customer.Customer_password = crypto_utils.GetMd5(customer.Customer_password)

	if err:= customer.Save(); err!= nil {
		return nil, err
	}
	return &customer, nil
}
//LoginUser


func (s *customerService) LoginUser(request customers_domains.CustomerLoginRequest) (*customers_domains.Customer,*rest_errors_package.RestErr) {
	dao := &customers_domains.Customer{
		Customer_email_id: request.Customer_email_id,
		Customer_password: crypto_utils.GetMd5(request.Customer_password),
	}
	if err:= dao.FindByEmailAndPassword(); err!= nil {
		return nil, err
	}

	return dao, nil
}