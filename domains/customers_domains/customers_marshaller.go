package customers_domains

import (
"encoding/json"
"fmt"
)

type PublicCustomer struct {
	Id 									int64
	Customer_name		 				string
	Customer_city						string
	Customer_state						string

	Customer_phone_verified			 	bool

	Customer_active 					bool

	Date_created 						string
}


type PrivateCustomer struct {
	Id 									int64
	Customer_name		 				string
	Customer_phone				 		string
	Customer_address					string
	Customer_city						string
	Customer_state						string

	Customer_phone_verified			 	bool

	Customer_email 						string
	Customer_active 					bool

	Date_created 						string
}

func (customer *Customer) Marshall(isPublic bool) (interface{}){

	customerJson, _ := json.Marshal(customer)

	if isPublic {
		var publicCustomer PublicCustomer

		if err:= json.Unmarshal(customerJson, &publicCustomer); err!= nil{
			fmt.Println("Error ", err.Error())
			return nil
		}
		return publicCustomer
	}

	var privateCustomer PrivateCustomer
	if err:= json.Unmarshal(customerJson, &privateCustomer); err!= nil{
		fmt.Println("Error ", err.Error())
		return nil
	}
	return privateCustomer
}


type Customers []Customer

func (customers Customers) Marshall(isPublic bool) ([]interface{}){
	results := make([]interface{}, len(customers))

	for index, customer := range customers {
		results[index] = customer.Marshall(isPublic)
	}
	return results
}

