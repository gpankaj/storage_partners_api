package services

import (
	"fmt"
	"github.com/gpankaj/storage_partners_api/domains/partners_domains"
	"github.com/gpankaj/storage_partners_api/utils/errors"
	"strings"
)

func Create_Partner_Service(partner partners_domains.Partner) (*partners_domains.Partner, *errors.RestErr) {
	partner_model := partners_domains.Partner{}
	partner_model.Email_id = strings.TrimSpace(strings.ToLower(partner.Email_id))

	if err:= partner.Validate(); err!=nil{
		return nil, err
	}

	if err:= partner.Save(); err!= nil {
		return nil, err
	}
	return &partner, nil
}




func Get_Partner_Service(partner_id int64) (*partners_domains.Partner, *errors.RestErr) {
	if partner_id <0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("Invalid partner id %d", partner_id))
	}

	partner_model := &partners_domains.Partner{Id: partner_id}

	if err:= partner_model.Get(); err!= nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("Partner with %d not found",partner_id))
	}
	return partner_model, nil
}