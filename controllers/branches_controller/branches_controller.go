package branches_controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/common-go-oauth/oauth"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/domains/branches_domains"
	"github.com/gpankaj/storage_partners_api/services"
	"log"
	"net/http"
	"strconv"
)

func getPartnerId(partnerIdParams string) (int64, *rest_errors_package.RestErr) {
	partner_id, partnerIdError := strconv.ParseInt(partnerIdParams,10,64)
	if partnerIdError!= nil {
		return 0, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", partnerIdError.Error()))

	}
	return partner_id, nil
}


func getBranchId(branchIdParams string) (int64, *rest_errors_package.RestErr) {
	branch_id, branchIdError := strconv.ParseInt(branchIdParams,10,64)
	if branchIdError!= nil {
		return 0, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", branchIdError.Error()))

	}
	return branch_id, nil
}

func getOwnerId(ownerIdParams string) (int64, *rest_errors_package.RestErr) {
	owner_id, ownerIdError := strconv.ParseInt(ownerIdParams,10,64)
	if ownerIdError!= nil {
		return 0, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", ownerIdError.Error()))

	}
	return owner_id, nil
}


func getPartnerStatus(partnerActiveParams string) (bool, *rest_errors_package.RestErr) {
	fmt.Println("Inside getPartnerStatus ",partnerActiveParams)
	status, statusError := strconv.ParseBool(partnerActiveParams)
	if statusError!= nil {
		//True is the default status.
		return true, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", statusError.Error()))

	}
	return status, nil
}

func CreateBranch(c *gin.Context) {


	if err:=oauth.AuthenticateRequest(c.Request);err!= nil {
		c.JSON(err.Code,err)
		return
	}

	log.Println("Caller Id", oauth.GetCallerId(c.Request))

	log.Println("Client Id", oauth.GetClientId(c.Request))


	branch_domain := branches_domains.NewBranch()
	if err:= c.ShouldBindJSON(&branch_domain); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together

		restError := rest_errors_package.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}
	branch_domain.Id = oauth.GetClientId(c.Request)

	fmt.Println("Branch in Controller is ", branch_domain)

	result, save_error := services.BranchService.Create_Branch_Service(*branch_domain)
	if save_error != nil {
		//TODO: Handle user creation Error
		c.JSON(save_error.Code, save_error)
		return
	}
	//fmt.Println("Partner Domain ",partner_domain)
	//c.String(http.StatusNotImplemented, "Implement Me!")
	c.JSON(http.StatusCreated, result)
	//c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))

}


func FindBranches(c *gin.Context) {
	//
	//status, statusError:=getPartnerStatus(c.Param("status"))
	allBranches, errFindBranches:= services.FindBranches();
	if errFindBranches!=nil{
		c.JSON(errFindBranches.Code, errFindBranches)
		return
	}
	c.JSON(http.StatusOK,allBranches)
}

//

func FindPartnerBranches(c *gin.Context) {
	//
	//status, statusError:=getPartnerStatus(c.Param("status"))
	log.Println("===============================")
	partner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}

	allPartnerBranches, errFindBranches:= services.FindPartnerBranches(partner_id);
	if errFindBranches!=nil{
		c.JSON(errFindBranches.Code, errFindBranches)
		return
	}
	c.JSON(http.StatusOK,allPartnerBranches)
}

func UpdateBranch(c *gin.Context) {

	if err:=oauth.AuthenticateRequest(c.Request);err!= nil {
		c.JSON(err.Code,err)
		return
	}

	log.Println("Caller Id", oauth.GetCallerId(c.Request))

	log.Println("Client Id", oauth.GetClientId(c.Request))

	branch_id, idError:=getBranchId(c.Param("branch_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}


	partner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}


	if oauth.GetClientId(c.Request) != partner_id {
		log.Println("Owner id ", partner_id)
		log.Println("Client id ", oauth.GetClientId(c.Request));

		log.Println("Given token does not belong to the owner")
		c.JSON(http.StatusBadRequest, errors.New("Client Id does not match with given token id"))
		return
	}
	log.Println("===You are the owner of this Branch..==")


	branch_domain := branches_domains.NewBranch()

	//Populate the partner with given user
	if err:= c.ShouldBindJSON(&branch_domain); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together

		restError := rest_errors_package.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}
	branch_domain.Id = partner_id
	branch_domain.Branch_id = branch_id


	isPartial := c.Request.Method == http.MethodPatch
	result, update_error := services.BranchService.Update_Branch_Service(isPartial,*branch_domain)
	if update_error != nil {
		//TODO: Handle user creation Error
		c.JSON(update_error.Code, update_error)
		return
	}
	//fmt.Println("Partner Domain ",partner_domain)
	//c.String(http.StatusNotImplemented, "Implement Me!")
	c.JSON(http.StatusOK, result)
	//c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//DeleteBranch


func DeleteBranch(c *gin.Context) {

	if err:=oauth.AuthenticateRequest(c.Request);err!= nil {
		c.JSON(err.Code,err)
		return
	}

	log.Println("Caller Id", oauth.GetCallerId(c.Request))

	log.Println("Client Id", oauth.GetClientId(c.Request))

	branch_id, idError:=getBranchId(c.Param("branch_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}


	partner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}


	if oauth.GetClientId(c.Request) != partner_id {
		log.Println("Owner id ", partner_id)
		log.Println("Client id ", oauth.GetClientId(c.Request));

		log.Println("Given token does not belong to the owner")
		c.JSON(http.StatusBadRequest, errors.New("Client Id does not match with given token id"))
		return
	}
	log.Println("===You are the owner of this Branch..==")

	if errDelete:= services.BranchService.Delete_Branch_Service(branch_id); errDelete!=nil {
		c.JSON(errDelete.Code, errDelete)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": fmt.Sprintf("deleted branch with id %d", branch_id)})
}