package partners_domains

import (
	"encoding/json"
	"fmt"
)

type PublicPartner struct {
	Id 									int64
	Storage_partner_name 				string
	Storage_partner_company_name 		string
	//Storage_partner_company_gst 		string
	Provides_goods_transport_service 	bool
	Provides_goods_packaging_service 	bool
	Provides_goods_insurance_service	bool
	Listing_active 						bool
	//Email_id 							string
	//Phone_numbers 						string

	Verified							bool
	//Password							string

	Date_created 						string
}


type PrivatePartner struct {
	Id 									int64
	Storage_partner_name 				string
	Storage_partner_company_name 		string
	Storage_partner_company_gst 		string
	Provides_goods_transport_service 	bool
	Provides_goods_packaging_service 	bool
	Provides_goods_insurance_service	bool
	Listing_active 						bool
	Email_id 							string
	Phone_numbers 						string

	Verified							bool
	//Password							string

	Date_created 						string
}

func (partner *Partner) Marshall(isPublic bool) (interface{}){

	partnerJson, _ := json.Marshal(partner)

	if isPublic {
		var publicPartner PublicPartner

		if err:= json.Unmarshal(partnerJson, &publicPartner); err!= nil{
			fmt.Println("Error ", err.Error())
			return nil
		}
		return publicPartner
	}

	var privatePartner PrivatePartner
	if err:= json.Unmarshal(partnerJson, &privatePartner); err!= nil{
		fmt.Println("Error ", err.Error())
		return nil
	}
	return privatePartner
}


type Partners []Partner

func (partners Partners) Marshall(isPublic bool) ([]interface{}){
	results := make([]interface{}, len(partners))

	for index, partner := range partners {
		results[index] = partner.Marshall(isPublic)
	}
	return results
}
