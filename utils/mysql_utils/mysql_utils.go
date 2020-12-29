package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"strings"
)

const (
	noRowInResultSet = "no rows in result set"
)
func ParseError(err error) (*rest_errors_package.RestErr) {
	sqlErr, not_mysql_type_error := err.(*mysql.MySQLError)
	if !not_mysql_type_error {
		if strings.ContainsAny(err.Error(),noRowInResultSet) {
			return rest_errors_package.NewInternalServerError(fmt.Sprintf("No Row Selected for given id %s", err.Error()),
				err)
		}
		return rest_errors_package.NewInternalServerError(fmt.Sprintf("While trying to Parse Query error %s ", err.Error()),
			err)
	}


	fmt.Println("Some SQL error Message ", sqlErr.Message)
	fmt.Println("Some SQL error Number", sqlErr.Number)
	fmt.Println("Some SQL error ", sqlErr.Error())
	switch sqlErr.Number {
	case 1062:
		return rest_errors_package.NewUniqueContraintViolationcompany_name_listing_active_uniqueError(fmt.Sprintf("Error while trying to Exec stmt %s ", sqlErr.Error()))
	}

	return rest_errors_package.NewInternalServerError(fmt.Sprintf("Error while trying to Exec stmt %s ", sqlErr.Error()),
		err)
}
