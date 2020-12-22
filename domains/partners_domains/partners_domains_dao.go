package partners_domains

import (
	"fmt"
	"github.com/gpankaj/storage_partners_api/utils/errors"
)

var (
	partnerDB = make(map[int64] *Partner)
)
func (partner *Partner)  Get() (*errors.RestErr){

	fmt.Println("Got partner id", partner.Id)
	fmt.Println("Content of partnerDB ", partnerDB)

	partner_from_db := partnerDB[partner.Id]
	if partner_from_db == nil {
		return errors.NewNotFoundError(fmt.Sprintf("Partner %d not found ", partner.Id))
	}

	partner.Email_id = partner_from_db.Email_id
	partner.Storage_partner_name = partner_from_db.Storage_partner_name

	partner.Storage_partner_company_name = partner_from_db.Storage_partner_company_name
	partner.Storage_partner_company_gst = partner_from_db.Storage_partner_company_gst

	partner.Provides_goods_transport_service = partner_from_db.Provides_goods_transport_service
	partner.Provides_goods_packaging_service = partner_from_db.Provides_goods_packaging_service

	partner.Provides_goods_insurance_service = partner_from_db.Provides_goods_insurance_service
	partner.Listing_active = partner_from_db.Listing_active

	partner.Phone_numbers = partner_from_db.Phone_numbers

	return nil
}

func (partner *Partner) Save() *errors.RestErr{
	partner_from_db := partnerDB[partner.Id]
	if partner_from_db!= nil {
		if partner_from_db.Email_id == partner.Email_id {
			return errors.NewBadRequestError(fmt.Sprintf("Partner %s already exists in db with same email", partner.Email_id))
		} else if (partner_from_db.Storage_partner_company_name == partner.Storage_partner_company_name) {
			return errors.NewBadRequestError(
				fmt.Sprintf(
					"Partner %s company already exists, you may contact %s for more info",
					partner.Storage_partner_company_name, partner.Email_id))
		}
		return errors.NewBadRequestError(fmt.Sprintf("Partner %d already exists in db", partner.Id))
	}
	partnerDB[partner.Id] = partner
	return nil
}