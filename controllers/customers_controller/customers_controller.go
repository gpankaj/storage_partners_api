package customers_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/domains/customers_domains"
	"github.com/gpankaj/storage_partners_api/services"
	"log"
	"net/http"
	"strconv"
)

func getPartnerId(partnerIdParams string) (int64, *rest_errors_package.RestErr) {
	partner_id, partnerIdError := strconv.ParseInt(partnerIdParams,10,64)
	if partnerIdError!= nil {
		return 0, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", partnerIdError.Error()))

	}
	return partner_id, nil
}


func getCustomerId(customerIdParams string) (int64, *rest_errors_package.RestErr) {
	customer_id, customerIdError := strconv.ParseInt(customerIdParams,10,64)
	if customerIdError!= nil {
		return 0, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", customerIdError.Error()))

	}
	return customer_id, nil
}


func getBranchId(branchIdParams string) (int64, *rest_errors_package.RestErr) {
	branch_id, branchIdError := strconv.ParseInt(branchIdParams,10,64)
	if branchIdError!= nil {
		return 0, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", branchIdError.Error()))

	}
	return branch_id, nil
}

func getOwnerId(ownerIdParams string) (int64, *rest_errors_package.RestErr) {
	owner_id, ownerIdError := strconv.ParseInt(ownerIdParams,10,64)
	if ownerIdError!= nil {
		return 0, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", ownerIdError.Error()))

	}
	return owner_id, nil
}


func getCustomerStatus(partnerActiveParams string) (bool, *rest_errors_package.RestErr) {
	fmt.Println("Inside getPartnerStatus ",partnerActiveParams)
	status, statusError := strconv.ParseBool(partnerActiveParams)
	if statusError!= nil {
		//True is the default status.
		return true, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", statusError.Error()))

	}
	return status, nil
}

func CreateCustomer(c *gin.Context) {


	customer_domain := customers_domains.NewCustomer()
	if err:= c.ShouldBindJSON(&customer_domain); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together

		restError := rest_errors_package.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}

	fmt.Println("Customer in Controller is ", customer_domain)

	result, save_error := services.CustomerService.Create_Customer_Service(*customer_domain)
	if save_error != nil {
		//TODO: Handle user creation Error
		c.JSON(save_error.Code, save_error)
		return
	}
	//fmt.Println("Partner Domain ",partner_domain)
	//c.String(http.StatusNotImplemented, "Implement Me!")
	c.JSON(http.StatusCreated, result)
	//c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))

}
//Login

func Login(c *gin.Context) {
	var request customers_domains.CustomerLoginRequest


	if err:= c.ShouldBindJSON(&request); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together
		restError := rest_errors_package.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}
	customer, err := services.CustomerService.LoginUser(request)

	log.Println("Request to login with Email", request.Customer_email_id);
	log.Println("Request to login with Password ", request.Customer_password);

	if request.Customer_email_id == ""||request.Customer_password == "" {
		restError := rest_errors_package.NewBadRequestError("Wrong Args")
		c.JSON(restError.Code, restError)
		return
	}

	if err!=nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, customer)
	//c.JSON(http.StatusOK, customer.Marshall(c.GetHeader("X-Public") == "true"))

	return
}
