package customers_domains

const (
	STORAGE = "true"
	TRANSPORT = "true"
	PACKAGING = "true"
	INSURANCE = "true"
)
type Customer struct {

	Id 									int64
	Customer_name		 				string
	Customer_phone				 		string
	Customer_address					string
	Customer_city						string
	Customer_state						string

	Customer_phone_verified			 	bool
	Customer_phone_verification_code	string

	Customer_comments					string

	Customer_email 						string
	Customer_active 					bool

	Date_created 						string

}