package branches_domains


import (
	"fmt"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/datasources/mysql/partners_db"
	"github.com/gpankaj/storage_partners_api/utils/mysql_utils"
	"github.com/gpankaj/storage_partners_api/logger"

	"log"
)

const  (
	//indexUniqueCompanyActiveName = "company_name_listing_active_unique"
	noRowInResultSet = "no rows in result set"
	queryInsertBranch = "INSERT INTO branches_table(City,Point_of_contact1," +
		"Point_of_contact2, Point_of_contact3,Branch_email_id," +
		"Remarks,Branch_verified,Branch_listing_active,Branch_date_created, Id) VALUES(?,?,?,?,?,?,?,?,?,?);"

	queryGetBranch = "SELECT Branch_id, City,Point_of_contact1,Point_of_contact2, " +
		"Point_of_contact3,Branch_email_id," +
		"Remarks,Branch_verified,Branch_listing_active,Branch_date_created, Id FROM branches_table WHERE Branch_id=?;"

	queryUpdateBranch = "UPDATE branches_table SET City=?, Point_of_contact1=?, Point_of_contact2=?," +
		"Point_of_contact3=?, Branch_email_id=?," +
		"Remarks=?, Branch_listing_active=? WHERE Branch_id=?;"

	queryDeleteBranch = "DELETE FROM branches_table WHERE Branch_id=?;"


	queryFindBranches = "SELECT Branch_id, City,Point_of_contact1,Point_of_contact2, " +
		"Point_of_contact3,Branch_email_id," +
		"Remarks,Branch_verified,Branch_date_created, Id, Branch_listing_active FROM branches_table"

	queryFindPartnerBranches = "SELECT Branch_id, City,Point_of_contact1,Point_of_contact2, " +
		"Point_of_contact3,Branch_email_id," +
		"Remarks,Branch_verified,Branch_date_created, Id, Branch_listing_active FROM branches_table WHERE Id=?;"

	queryFindPartnerByOwner = "SELECT Id, Storage_partner_name,Storage_partner_company_name,Storage_partner_company_gst, " +
		"Provides_goods_transport_service,Provides_goods_packaging_service," +
		"Provides_goods_insurance_service,Listing_active,Phone_numbers,Email_id,Date_created,Verified FROM partners_table WHERE Id=?;"




	queryFindEmailAndPassword="SELECT Id, Storage_partner_name,Storage_partner_company_name,Storage_partner_company_gst," +
		"Provides_goods_transport_service,Provides_goods_packaging_service," +
		"Provides_goods_insurance_service,Listing_active,Phone_numbers,Email_id,Date_created,Verified " +
		"FROM partners_table WHERE Email_id=? AND Password=?;"

)

func (branch *Branch)  Get() (*rest_errors_package.RestErr){

	if err:=partners_db.Client.Ping(); err!= nil {
		panic(err)
	}

	fmt.Println("Got branch id", branch.Branch_id)

	stmt,err:=partners_db.Client.Prepare(queryGetBranch)
	if err!=nil {
		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in GET ",err.Error()), err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(branch.Branch_id)
	if getErr:= result.Scan(&branch.Branch_id,&branch.City, &branch.Point_of_contact1,
		&branch.Point_of_contact2,&branch.Point_of_contact3,&branch.Branch_email_id,
		&branch.Remarks,&branch.Branch_verified, &branch.Branch_listing_active,&branch.Branch_date_created,
		&branch.Id); getErr!=nil{

		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (branch *Branch) Save() *rest_errors_package.RestErr{
	fmt.Println(branch)
	stmt ,err := partners_db.Client.Prepare(queryInsertBranch)

	if err!=nil {
		log.Println("Failed in DAO while preparing stmt " , queryInsertBranch)

		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt ",err.Error()),err)
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()

	log.Println("Branch in save of DAO is ", branch);

	result, saveError := stmt.Exec(
		branch.City, branch.Point_of_contact1, branch.Point_of_contact2,
		branch.Point_of_contact3, branch.Branch_email_id, branch.Remarks, branch.Branch_verified, branch.Branch_listing_active,
		branch.Branch_date_created, branch.Id)

	if saveError!= nil {
		return mysql_utils.ParseError(saveError)
	}

	branch_id, err := result.LastInsertId()

	if err!= nil {
		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while trying to get LastInsertId %s ", err.Error()),
			err)
	}
	branch.Branch_id = branch_id

	return nil

}
func FindBranches() ([]Branch,*rest_errors_package.RestErr) {
	stmt ,err := partners_db.Client.Prepare(queryFindBranches)
	if err!=nil {
		log.Println("Failed to prepare query ", queryFindBranches)
		return nil,rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in FindByPartnerActive ",err.Error()),
			err)
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()
	rows, findBranchesError := stmt.Query()
	if findBranchesError!= nil {
		log.Println("Failed to find branches while running stmt ", findBranchesError)
		return nil,mysql_utils.ParseError(findBranchesError)
	}

	defer rows.Close()
	//We do not know how many results will be there, so we make a slice of size 0 of data.
	results := make([]Branch,0)

	for rows.Next() {
		var branch Branch
		err:=rows.Scan(&branch.Branch_id, &branch.City, &branch.Point_of_contact1,
			&branch.Point_of_contact2,&branch.Point_of_contact3, &branch.Branch_email_id,
			&branch.Remarks, &branch.Branch_verified,&branch.Branch_date_created,&branch.Id,
			&branch.Branch_listing_active,
			)
		if err!= nil{
			log.Println("Failed to get branches..")
			return nil,mysql_utils.ParseError(findBranchesError)
		}
		results = append(results, branch)
	}

	if len(results) == 0 {
		return nil, rest_errors_package.NewNotFoundError(fmt.Sprintf("No Branches Found"))
	}
	return results, nil
}

//


func FindPartnerBranches(partner_id int64) ([]Branch,*rest_errors_package.RestErr) {
	stmt ,err := partners_db.Client.Prepare(queryFindPartnerBranches)

	if err!=nil {
		log.Println("Failed to prepare query ", queryFindPartnerBranches)
		return nil,rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in FindPartnerBranches ",err.Error()),
			err)
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()
	rows, findPartnerBranchesError := stmt.Query(partner_id)
	if findPartnerBranchesError!= nil {
		log.Println("Failed to find branches while running stmt ", findPartnerBranchesError)
		return nil,mysql_utils.ParseError(findPartnerBranchesError)
	}

	defer rows.Close()
	//We do not know how many results will be there, so we make a slice of size 0 of data.
	results := make([]Branch,0)

	for rows.Next() {
		var branch Branch
		err:=rows.Scan(&branch.Branch_id, &branch.City, &branch.Point_of_contact1,
			&branch.Point_of_contact2,&branch.Point_of_contact3, &branch.Branch_email_id,
			&branch.Remarks, &branch.Branch_verified,&branch.Branch_date_created, &branch.Id,
			&branch.Branch_listing_active,
		)
		if err!= nil{
			log.Println("Failed to get branches for a partner.. %d" , partner_id)
			return nil,mysql_utils.ParseError(findPartnerBranchesError)
		}
		results = append(results, branch)
	}

	if len(results) == 0 {
		return nil, rest_errors_package.NewNotFoundError(fmt.Sprintf("No Branches Found"))
	}
	return results, nil
}



func (branch *Branch)Update() *rest_errors_package.RestErr{
	stmt ,err := partners_db.Client.Prepare(queryUpdateBranch)

	if err!=nil {
		logger.Error("Error while preparing statement inside Update function called by used", err)
		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt ",err.Error()), err)
	}
	//Storage_partner_company_name
	defer stmt.Close()
	log.Println("Printing inside Update ", branch)

	//partner.Date_created = date_utils.GetNowString()
	result, updateError := stmt.Exec(branch.City, branch.Point_of_contact1, branch.Point_of_contact2,
		branch.Point_of_contact3, branch.Branch_email_id, branch.Remarks,
		branch.Branch_listing_active, branch.Branch_id)

	rows_affected, error := result.RowsAffected()
	if error!= nil {
		log.Println(error)
	}

	log.Println("Update result ", rows_affected)

	if updateError!= nil {
		return mysql_utils.ParseError(updateError)
	}
	return nil
}
//Delete


func (branch *Branch)Delete() *rest_errors_package.RestErr{
	stmt ,err := partners_db.Client.Prepare(queryDeleteBranch)

	if err!=nil {
		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in DELETE ",err.Error()),
			err)
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()
	_, deleteError := stmt.Exec(branch.Branch_id)

	if deleteError!= nil {
		return mysql_utils.ParseError(deleteError)
	}
	return nil
}
