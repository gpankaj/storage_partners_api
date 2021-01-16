package customers_domains

import (
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/utils/date_utils"
	"strings"
)

const (
	STORAGE = "true"
	TRANSPORT = "true"
	PACKAGING = "true"
	INSURANCE = "true"
)
type Customer struct {
	Customer_id 						int64
	Customer_name		 				string
	Customer_phone_number		 		string
	Customer_address					string
	Customer_city						string
	Customer_state						string
	Customer_phone_verified			 	bool
	Customer_phone_verification_code	string
	Customer_comments					string
	Customer_email_id 					string
	Customer_active 					bool
	Customer_date_created 				string
	Customer_password					string
	Customer_verified					bool
}

// Constructor return intreface
func NewCustomer() *Customer {
	return &Customer{Customer_name: "", Customer_phone_number: "", Customer_address:"",
		Customer_city: "", Customer_state: "", Customer_phone_verified: true,
		Customer_phone_verification_code: "",Customer_comments: "",Customer_email_id: "",
		Customer_active: true, Customer_password: "",Customer_date_created : date_utils.GetNowDB(), Customer_verified : false}
}

func (customer *Customer) Validate(isPatch bool) (*rest_errors_package.RestErr){
	customer.Customer_name = strings.TrimSpace(customer.Customer_name)
	customer.Customer_phone_number = strings.TrimSpace(customer.Customer_phone_number)
	customer.Customer_address = strings.TrimSpace(customer.Customer_address)

	customer.Customer_city = strings.TrimSpace(customer.Customer_city)

	customer.Customer_state = strings.TrimSpace(customer.Customer_state)

	customer.Customer_comments = strings.TrimSpace(customer.Customer_comments)

	customer.Customer_phone_verification_code = strings.TrimSpace(customer.Customer_phone_verification_code)
	customer.Customer_email_id = strings.TrimSpace(customer.Customer_email_id)



	if (!isPatch) {
		if customer.Customer_email_id == "" {
			return rest_errors_package.NewBadRequestError("Email id can not be empty")
		}
		if customer.Customer_password == "" {
			return rest_errors_package.NewBadRequestError("Company Password field can not be empty.")
		}
	}
	return nil
}
