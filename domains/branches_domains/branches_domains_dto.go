package branches_domains

import (
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/utils/date_utils"
	"strings"
)

type Branch struct {
	Branch_id 							int64
	City				 				string
	Point_of_contact1			 		string
	Point_of_contact2 					string
	Point_of_contact3				 	string
	Branch_email_id 					string
	Remarks		 						string

	Branch_verified						bool
	Branch_listing_active 				bool

	Branch_date_created 				string
	Id									int64
}
// Constructor return intreface
func NewBranch() *Branch {
	return &Branch{City: "", Point_of_contact1: "", Point_of_contact2:"",
		Point_of_contact3: "", Branch_email_id: "", Remarks: "",
		Branch_listing_active: true, Branch_date_created : date_utils.GetNowDB(), Branch_verified : false}
}

func (branch *Branch) Validate(isPatch bool) (*rest_errors_package.RestErr){

	branch.City = strings.TrimSpace(branch.City)
	branch.Point_of_contact1 = strings.TrimSpace(branch.Point_of_contact1)

	branch.Point_of_contact2 = strings.TrimSpace(branch.Point_of_contact2)

	branch.Point_of_contact3 = strings.TrimSpace(branch.Point_of_contact3)

	branch.Branch_email_id = strings.TrimSpace(branch.Branch_email_id)
	branch.Remarks = strings.TrimSpace(branch.Remarks)

	if (!isPatch) {

		if branch.Point_of_contact1 == "" {
			return rest_errors_package.NewBadRequestError("Point_of_contact1 cannot be empty")
		}

		if branch.City == "" {
			//TEMP
			//partner.Password = "xx"
			return rest_errors_package.NewBadRequestError("Branch City can not be empty")
		}
	}
	return nil
}

type Branches []Branch


