package customers_domains


import (
	"fmt"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/datasources/mysql/partners_db"
	"github.com/gpankaj/storage_partners_api/utils/mysql_utils"
	"log"
)

const  (
	//indexUniqueCompanyActiveName = "company_name_listing_active_unique"
	noRowInResultSet = "no rows in result set"
	queryInsertCustomer = "INSERT INTO customers_table(Customer_name,Customer_phone_number," +
		"Customer_address, Customer_city,Customer_state," +
		"Customer_phone_verified,Customer_comments,Customer_email_id,Customer_password,Customer_active,Customer_date_created) VALUES(?,?,?,?,?,?,?,?,?,?,?);"

	queryGetCustomer = "SELECT Customer_id, Customer_name,Customer_phone_number,Customer_address, " +
		"Customer_city,Customer_state," +
		"Customer_phone_verified,Customer_comments,Customer_email_id,Customer_active,Customer_date_created,Customer_verified FROM customers_table WHERE Customer_id=?;"

	queryUpdatePartner = "UPDATE partners_table SET Storage_partner_name=?, Storage_partner_company_name=?, Storage_partner_company_gst=?," +
		"Provides_goods_transport_service=?, Provides_goods_packaging_service=?," +
		"Provides_goods_insurance_service=?, Listing_active=?, Phone_numbers=?, Email_id=? WHERE Id=?;"

	queryDeletePartner = "DELETE FROM partners_table WHERE id=?;"

	queryFindPartnerIfActive = "SELECT Id, Storage_partner_name,Storage_partner_company_name,Storage_partner_company_gst, " +
		"Provides_goods_transport_service,Provides_goods_packaging_service," +
		"Provides_goods_insurance_service,Listing_active,Phone_numbers,Email_id,Date_created,Verified FROM partners_table WHERE Listing_active=?;"

	queryFindPartnerByOwner = "SELECT Id, Storage_partner_name,Storage_partner_company_name,Storage_partner_company_gst, " +
		"Provides_goods_transport_service,Provides_goods_packaging_service," +
		"Provides_goods_insurance_service,Listing_active,Phone_numbers,Email_id,Date_created,Verified FROM partners_table WHERE Id=?;"




	queryFindEmailAndPassword="SELECT Customer_id, Customer_name,Customer_phone_number,Customer_address," +
		"Customer_city,Customer_state," +
		"Customer_phone_verified,Customer_comments,Customer_email_id,Customer_active,Customer_date_created,Customer_verified " +
		"FROM customers_table WHERE Customer_email_id=? AND Customer_password=?;"

)

func (customer *Customer)  Get() (*rest_errors_package.RestErr){

	if err:=partners_db.Client.Ping(); err!= nil {
		panic(err)
	}

	fmt.Println("Got customer with id", customer.Customer_id)

	stmt,err:=partners_db.Client.Prepare(queryGetCustomer)
	if err!=nil {
		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in GET ",err.Error()), err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(customer.Customer_id)
	if getErr:= result.Scan(&customer.Customer_id,&customer.Customer_name, &customer.Customer_phone_number,
		&customer.Customer_address,&customer.Customer_city,&customer.Customer_state,
		&customer.Customer_phone_verified,&customer.Customer_comments,&customer.Customer_email_id,&customer.Customer_active,
		&customer.Customer_date_created,
		&customer.Customer_verified); getErr!=nil{

		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (customer *Customer) Save() *rest_errors_package.RestErr{
	fmt.Println(customer)
	stmt ,err := partners_db.Client.Prepare(queryInsertCustomer)

	if err!=nil {
		log.Println("Failed in DAO while preparing stmt " , queryInsertCustomer)

		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt ",err.Error()),err)
	}
	//Storage_partner_company_name
	defer stmt.Close()
	//partner.Date_created = date_utils.GetNowString()

	log.Println("Customer in save of DAO is ", customer);

	result, saveError := stmt.Exec(
		customer.Customer_name, customer.Customer_phone_number, customer.Customer_address,
		customer.Customer_city, customer.Customer_state, customer.Customer_phone_verified,
		customer.Customer_comments, customer.Customer_email_id, customer.Customer_password, customer.Customer_active, customer.Customer_date_created)

	if saveError!= nil {
		return mysql_utils.ParseError(saveError)
	}

	customer_id, err := result.LastInsertId()

	if err!= nil {
		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while trying to get LastInsertId %s ", err.Error()),
			err)
	}
	customer.Customer_id = customer_id

	return nil
}

//FindByEmailAndPassword

func (customer *Customer) FindByEmailAndPassword() (*rest_errors_package.RestErr){

	if err:=partners_db.Client.Ping(); err!= nil {
		panic(err)
	}

	fmt.Println("Got customer id", customer.Customer_id)

	stmt,err:=partners_db.Client.Prepare(queryFindEmailAndPassword)
	if err!=nil {
		return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while preparing stmt in FindByEmailAndPassword ",err.Error()),err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(customer.Customer_email_id, customer.Customer_password)
	log.Println("result ", result)


	if getErr:= result.Scan(&customer.Customer_id,&customer.Customer_name, &customer.Customer_phone_number,
		&customer.Customer_address,&customer.Customer_city,&customer.Customer_state,
		&customer.Customer_phone_verified,&customer.Customer_comments,&customer.Customer_email_id,&customer.Customer_active,
		&customer.Customer_date_created,
		&customer.Customer_verified); getErr!=nil{

		return mysql_utils.ParseError(getErr)
	}
	return nil
}