package services

import (
	"fmt"
	"github.com/gpankaj/storage_partners_api/domains/partners_domains"
	"github.com/gpankaj/storage_partners_api/utils/crypto_utils"
	"github.com/gpankaj/storage_partners_api/utils/errors"
	"log"
)
var (
	PartnerService partnerServiceInterface = &partnerService{}
)

type partnerService struct {
}

type partnerServiceInterface interface {
	Create_Partner_Service( partners_domains.Partner) (*partners_domains.Partner, *errors.RestErr)
	Get_Partner_Service(partner_id int64) (*partners_domains.Partner, *errors.RestErr)
	Update_Partner_Service(bool,partners_domains.Partner) (*partners_domains.Partner, *errors.RestErr)
	Delete_Partner_Service(int64) *errors.RestErr
	LoginUser(partners_domains.PartnerLoginRequest) (*partners_domains.Partner,*errors.RestErr)
}

func (s *partnerService) Create_Partner_Service(partner partners_domains.Partner) (*partners_domains.Partner, *errors.RestErr) {
	/*
	partner_model := partners_domains.Partner{}

	partner_model.Email_id = strings.TrimSpace(strings.ToLower(partner.Email_id))
	*/
	if err:= partner.Validate(false); err!=nil{
		return nil, err
	}

	partner.Password = crypto_utils.GetMd5(partner.Password)

	if err:= partner.Save(); err!= nil {
		return nil, err
	}
	return &partner, nil
}




func (s *partnerService) Get_Partner_Service(partner_id int64) (*partners_domains.Partner, *errors.RestErr) {
	if partner_id <0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("Invalid partner id %d", partner_id))
	}

	partner_model := &partners_domains.Partner{Id: partner_id}

	if err:= partner_model.Get(); err!= nil {
		return nil, err
	}
	return partner_model, nil
}

func (s *partnerService) Update_Partner_Service(isPartial bool,partner partners_domains.Partner) (*partners_domains.Partner, *errors.RestErr) {

	//User from DB
	partner_from_db, err := PartnerService.Get_Partner_Service(partner.Id);


	if err!=nil{
		return nil, err
	}



	if isPartial {
		if err:= partner.Validate(isPartial); err!=nil {
			return nil, err
		}


		if partner.Storage_partner_name != "" {
			partner_from_db.Storage_partner_name = partner.Storage_partner_name
		}

		if partner.Storage_partner_company_name != "" {
			partner_from_db.Storage_partner_company_name = partner.Storage_partner_company_name
		}

		if partner.Provides_goods_transport_service {
			partner_from_db.Provides_goods_transport_service = true
		} else {
			partner_from_db.Provides_goods_transport_service = false
		}

		if partner.Provides_goods_packaging_service {
			partner_from_db.Provides_goods_packaging_service = true
		} else{
			partner_from_db.Provides_goods_packaging_service = false
		}

		if partner.Provides_goods_insurance_service {
			partner_from_db.Provides_goods_insurance_service = true
		} else{
			partner_from_db.Provides_goods_insurance_service = false
		}

		log.Println("Parter listing ", partner.Listing_active);

		if partner.Listing_active {
			partner_from_db.Listing_active = true
		} else{
			partner_from_db.Listing_active = false
		}

		if partner.Email_id !="" {
			partner_from_db.Email_id = partner.Email_id
		}
		if partner.Phone_numbers !="" {
			partner_from_db.Phone_numbers = partner.Phone_numbers
		}

	} else {
		if err:= partner.Validate(isPartial); err!=nil {
			return nil, err
		}

		partner_from_db.Storage_partner_name = partner.Storage_partner_name
		partner_from_db.Storage_partner_company_name = partner.Storage_partner_company_name

		partner_from_db.Provides_goods_transport_service = partner.Provides_goods_transport_service
		partner_from_db.Provides_goods_packaging_service = partner.Provides_goods_packaging_service

		partner_from_db.Provides_goods_insurance_service = partner.Provides_goods_insurance_service
		partner_from_db.Listing_active = partner.Listing_active

		partner_from_db.Email_id = partner.Email_id
		partner_from_db.Phone_numbers = partner.Phone_numbers
	}


	if err:=partner_from_db.Update(); err!=nil {
		return nil, err
	}

	return partner_from_db, nil
	//partner_from_db

	//Now we know the user is in DB so let us update it.


	return nil, nil
}

func (s *partnerService) Delete_Partner_Service(partner_id int64) *errors.RestErr {
	//Check if user is in DB
	partner_from_db, err := PartnerService.Get_Partner_Service(partner_id);
	if err != nil {
		return err
	}

	if partner_from_db == nil {
		return errors.NewNotFoundError(fmt.Sprintf("No such used with id %d", partner_id ))
	}

	deleteError := partner_from_db.Delete()
	if deleteError!= nil {

	}
	return nil
}

func (s *partnerService) LoginUser(request partners_domains.PartnerLoginRequest) (*partners_domains.Partner,*errors.RestErr) {
	dao := &partners_domains.Partner{
		Email_id: request.Email_id,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err:= dao.FindByEmailAndPassword(); err!= nil {
		return nil, err
	}

	return dao, nil
}

func FindByPartnerActive(status bool)  (partners_domains.Partners , *errors.RestErr) {
	return partners_domains.FindByPartnerActive(status)
}