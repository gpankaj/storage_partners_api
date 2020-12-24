package partners_domains

import (
	"github.com/gpankaj/storage_partners_api/utils/date_utils"
	"github.com/gpankaj/storage_partners_api/utils/errors"
	"strings"
)

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
// Constructor return intreface
func NewPartner() *Partner {
	return &Partner{Storage_partner_name: "", Storage_partner_company_name: "", Storage_partner_company_gst:"",
		Provides_goods_transport_service: true, Provides_goods_packaging_service: true, Provides_goods_insurance_service: true,
		Listing_active: true, Email_id: "",Date_created : date_utils.GetNowString()}
}

func (partner *Partner) Validate(isPatch bool) (*errors.RestErr){

	partner.Storage_partner_name = strings.TrimSpace(partner.Storage_partner_name)
	partner.Storage_partner_company_name = strings.TrimSpace(partner.Storage_partner_company_name)

	partner.Storage_partner_company_gst = strings.TrimSpace(partner.Storage_partner_company_gst)

	partner.Phone_numbers = strings.TrimSpace(partner.Phone_numbers)

	if (!isPatch) {
		if partner.Email_id == "" {
			return errors.NewBadRequestError("Email id can not be empty")
		}
		if partner.Storage_partner_company_name == "" {
			return errors.NewBadRequestError("Partner Company Name can not be empty.")
		}
	}
	return nil
}

