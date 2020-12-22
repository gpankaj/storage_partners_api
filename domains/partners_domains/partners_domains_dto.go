package partners_domains

import "github.com/gpankaj/storage_partners_api/utils/errors"

type Partner struct {
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

	Date_created 						string
}

func (partner *Partner) Validate() (*errors.RestErr){
	if partner.Email_id == "" {
		return errors.NewBadRequestError("Email id can not be empty")
	}
	if partner.Storage_partner_company_name == "" {
		return errors.NewBadRequestError("Partner Company Name can not be empty.")
	}
	return nil
}

