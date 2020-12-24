package partners_domains

import (
	"fmt"
	"github.com/gpankaj/storage_partners_api/datasources/mysql/partners_db"
	"github.com/gpankaj/storage_partners_api/utils/errors"
	"github.com/gpankaj/storage_partners_api/utils/mysql_utils"
)

const  (
	//indexUniqueCompanyActiveName = "company_name_listing_active_unique"
	noRowInResultSet = "no rows in result set"
	queryInsertPartner = "INSERT INTO partners_table(Storage_partner_name,Storage_partner_company_name," +
		"Storage_partner_company_gst, Provides_goods_transport_service,Provides_goods_packaging_service," +
		"Provides_goods_insurance_service,Listing_active,Phone_numbers,Email_id,Date_created, Password) VALUES(?,?,?,?,?,?,?,?,?,?,?);"

	queryGetPartner = "SELECT Id, Storage_partner_name,Storage_partner_company_name,Storage_partner_company_gst, " +
		"Provides_goods_transport_service,Provides_goods_packaging_service," +
		"Provides_goods_insurance_service,Listing_active,Phone_numbers,Email_id,Date_created,Verified FROM partners_table WHERE Id=?;"

	queryUpdatePartner = "UPDATE partners_table SET Storage_partner_name=?, Storage_partner_company_name=?, Storage_partner_company_gst=?," +
		"Provides_goods_transport_service=?, Provides_goods_packaging_service=?," +
		"Provides_goods_insurance_service=?, Listing_active=?, Phone_numbers=?, Email_id=? WHERE Id=?;"

	queryDeletePartner = "DELETE FROM partners_table WHERE id=?;"

	queryFindPartnerIfActive = "SELECT Id, Storage_partner_name,Storage_partner_company_name,Storage_partner_company_gst, " +
		"Provides_goods_transport_service,Provides_goods_packaging_service," +
		"Provides_goods_insurance_service,Listing_active,Phone_numbers,Email_id,Date_created,Verified FROM partners_table WHERE Listing_active=? ;"

)

func (partner *Partner)  Get() (*errors.RestErr){

	if err:=partners_db.Client.Ping(); err!= nil {
		panic(err)
	}

	fmt.Println("Got partner id", partner.Id)

	stmt,err:=partners_db.Client.Prepare(queryGetPartner)
	if err!=nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in GET ",err.Error()))
	}
	defer stmt.Close()

	result := stmt.QueryRow(partner.Id)
	if getErr:= result.Scan(&partner.Id,&partner.Storage_partner_name, &partner.Storage_partner_company_name,
		&partner.Storage_partner_company_gst,&partner.Provides_goods_transport_service,&partner.Provides_goods_packaging_service,
		&partner.Provides_goods_insurance_service,&partner.Listing_active,&partner.Phone_numbers,&partner.Email_id,&partner.Date_created,
		&partner.Verified); getErr!=nil{

		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (partner *Partner) Save() *errors.RestErr{
	fmt.Println(partner)
	stmt ,err := partners_db.Client.Prepare(queryInsertPartner)

	if err!=nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while preparing stmt ",err.Error()))
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()
	result, saveError := stmt.Exec(
		partner.Storage_partner_name, partner.Storage_partner_company_name, partner.Storage_partner_company_gst,
		partner.Provides_goods_transport_service, partner.Provides_goods_packaging_service, partner.Provides_goods_insurance_service,
		partner.Listing_active, partner.Phone_numbers, partner.Email_id, partner.Date_created, partner.Password)

	if saveError!= nil {
		return mysql_utils.ParseError(saveError)
	}

	partner_id, err := result.LastInsertId()

	if err!= nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to get LastInsertId %s ", err.Error()))
	}
	partner.Id = partner_id

	return nil

}

func (partner *Partner)Update() *errors.RestErr{
	stmt ,err := partners_db.Client.Prepare(queryUpdatePartner)

	if err!=nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while preparing stmt ",err.Error()))
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()
	_, updateError := stmt.Exec(partner.Storage_partner_name, partner.Storage_partner_company_name, partner.Storage_partner_company_gst,
		partner.Provides_goods_transport_service, partner.Provides_goods_packaging_service, partner.Provides_goods_insurance_service,
		partner.Listing_active, partner.Phone_numbers, partner.Email_id, partner.Id)

	if updateError!= nil {
		return mysql_utils.ParseError(updateError)
	}
	return nil
}

func (partner *Partner)Delete() *errors.RestErr{
	stmt ,err := partners_db.Client.Prepare(queryDeletePartner)

	if err!=nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in DELETE ",err.Error()))
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()
	_, deleteError := stmt.Exec(partner.Id)

	if deleteError!= nil {
		return mysql_utils.ParseError(deleteError)
	}
	return nil
}
//

func FindByPartnerActive(status bool) ([]Partner, *errors.RestErr){
	stmt ,err := partners_db.Client.Prepare(queryFindPartnerIfActive)

	if err!=nil {
		return nil,errors.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in FindByPartnerActive ",err.Error()))
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()
	rows, findByPartnerActiveError := stmt.Query(status)
	if findByPartnerActiveError!= nil {
		return nil,mysql_utils.ParseError(findByPartnerActiveError)
	}

	defer rows.Close()
	//We do not know how many results will be there, so we make a slice of size 0 of data.
	results := make([]Partner,0)

	for rows.Next() {
		var partner Partner
		err:=rows.Scan(&partner.Id, &partner.Storage_partner_name, &partner.Storage_partner_company_name,
			&partner.Storage_partner_company_gst,&partner.Provides_goods_transport_service, &partner.Provides_goods_packaging_service,
			&partner.Provides_goods_insurance_service, &partner.Listing_active,&partner.Phone_numbers,
			&partner.Email_id,&partner.Date_created,&partner.Verified)
		if err!= nil{
			return nil,mysql_utils.ParseError(findByPartnerActiveError)
		}
		results = append(results, partner)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No matching user found for given status %t", status))
	}
	return results, nil
}