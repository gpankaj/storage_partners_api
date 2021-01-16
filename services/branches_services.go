package services

import (
	"fmt"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/domains/branches_domains"
	"log"
)

var (
	BranchService branchServiceInterface = &branchService{}
)

type branchService struct {
}

type branchServiceInterface interface {
	Create_Branch_Service( branches_domains.Branch) (*branches_domains.Branch, *rest_errors_package.RestErr)
	Get_Branch_Service(branch_id int64) (*branches_domains.Branch, *rest_errors_package.RestErr)

	Update_Branch_Service(bool,branches_domains.Branch) (*branches_domains.Branch, *rest_errors_package.RestErr)
	Delete_Branch_Service(int64) *rest_errors_package.RestErr
}

func (s *branchService) Create_Branch_Service(branch branches_domains.Branch) (*branches_domains.Branch, *rest_errors_package.RestErr) {
	/*
		partner_model := partners_domains.Partner{}

		partner_model.Email_id = strings.TrimSpace(strings.ToLower(partner.Email_id))
	*/
	if err:= branch.Validate(false); err!=nil{
		return nil, err
	}

	if err:= branch.Save(); err!= nil {
		return nil, err
	}
	return &branch, nil
}




func (s *branchService) Get_Branch_Service(branch_id int64) (*branches_domains.Branch, *rest_errors_package.RestErr) {
	if branch_id <0 {
		return nil, rest_errors_package.NewBadRequestError(fmt.Sprintf("Invalid branch_id in service %d", branch_id))
	}

	branch_model := &branches_domains.Branch{Branch_id: branch_id}

	if err:= branch_model.Get(); err!= nil {
		return nil, err
	}
	return branch_model, nil
}

//FindBranches

func FindBranches()(branches_domains.Branches , *rest_errors_package.RestErr) {
	return branches_domains.FindBranches()
}

func FindPartnerBranches(partner_id int64)(branches_domains.Branches , *rest_errors_package.RestErr) {
	return branches_domains.FindPartnerBranches(partner_id)
}


func (s *branchService) Update_Branch_Service(isPartial bool,branch branches_domains.Branch) (*branches_domains.Branch, *rest_errors_package.RestErr) {

	//User from DB
	branch_from_db, err := BranchService.Get_Branch_Service(branch.Branch_id);
	if err!=nil{
		return nil, err
	}


	if isPartial {
		if err:= branch.Validate(isPartial); err!=nil {
			return nil, err
		}

		log.Println("Branch from user is ", branch)


		if branch.City != "" {
			branch_from_db.City = branch.City
		}


		if branch.Point_of_contact1 != "" {
			branch_from_db.Point_of_contact1 = branch.Point_of_contact1
		}

		if branch.Point_of_contact2 != "" {
			branch_from_db.Point_of_contact2 = branch.Point_of_contact2
		}
		if branch.Point_of_contact3 != "" {
			branch_from_db.Point_of_contact3 = branch.Point_of_contact3
		}

		if branch.Branch_email_id != "" {
			branch_from_db.Branch_email_id = branch.Branch_email_id
		}
		if branch.Remarks != "" {
			branch_from_db.Remarks = branch.Remarks
		}


		if branch.Branch_verified {
			branch_from_db.Branch_verified = true
		} else {
			branch_from_db.Branch_verified = false
		}

		if branch.Branch_listing_active {
			branch_from_db.Branch_listing_active = true
		} else{
			branch_from_db.Branch_listing_active = false
		}

	} else {
		if err:= branch.Validate(isPartial); err!=nil {
			return nil, err
		}

		branch_from_db.City = branch.City
		branch_from_db.Point_of_contact1 = branch.Point_of_contact1

		branch_from_db.Point_of_contact2 = branch.Point_of_contact2
		branch_from_db.Point_of_contact3 = branch.Point_of_contact3

		branch_from_db.Branch_email_id = branch.Branch_email_id
		branch_from_db.Remarks = branch.Remarks

	}


	if err:=branch_from_db.Update(); err!=nil {
		return nil, err
	}

	return branch_from_db, nil
	//partner_from_db

	//Now we know the user is in DB so let us update it.


	return nil, nil
}

//


func (s *branchService) Delete_Branch_Service(branch_id int64) *rest_errors_package.RestErr {
	//Check if user is in DB
	branch_from_db, err := BranchService.Get_Branch_Service(branch_id)

	if err != nil {
		return err
	}

	if branch_from_db == nil {
		return rest_errors_package.NewNotFoundError(fmt.Sprintf("No such used with id %d", branch_id ))
	}

	deleteError := branch_from_db.Delete()
	if deleteError!= nil {

	}
	return nil
}