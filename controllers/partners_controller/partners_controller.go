package partners_controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/common-go-oauth/oauth"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/gpankaj/storage_partners_api/domains/partners_domains"
	"github.com/gpankaj/storage_partners_api/services"
	"log"
	"net/http"
	"strconv"
)

func TestServiceInterface() {

}

func getPartnerId(partnerIdParams string) (int64, *rest_errors_package.RestErr) {
	partner_id, partnerIdError := strconv.ParseInt(partnerIdParams,10,64)
	if partnerIdError!= nil {
		return 0, rest_errors_package.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", partnerIdError.Error()))

	}
	return partner_id, nil
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

func CreatePartner(c *gin.Context) {
	//var partner_domain partners_domains.Partner

	/*
	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err!= nil {
		print(err.Error())
		//TODO: Handle reading request data error
		return
	}
	//bytes is an array which we use to populate from bytes to partner_domain.
	if err:= json.Unmarshal(bytes, &partner_domain); err!=nil {
		print(err.Error())
		//TODO: Handle unmarshal error
		return
	}
	*/
	partner_domain := partners_domains.NewPartner()
	if err:= c.ShouldBindJSON(&partner_domain); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together

		restError := rest_errors_package.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}
	fmt.Println("Password is ", partner_domain)
	result, save_error := services.PartnerService.Create_Partner_Service(*partner_domain)
	if save_error != nil {
		//TODO: Handle user creation Error
		log.Println("Failed to save new partner in controller ", save_error)
		c.JSON(save_error.Code, save_error)
		return
	}
	//fmt.Println("Partner Domain ",partner_domain)
	//c.String(http.StatusNotImplemented, "Implement Me!")
	//c.JSON(http.StatusCreated, result)
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdatePartner(c *gin.Context) {
	//create a partner domain.

	if err:=oauth.AuthenticateRequest(c.Request);err!= nil {
		c.JSON(err.Code,err)
		return
	}

	log.Println("Caller Id", oauth.GetCallerId(c.Request))

	log.Println("Client Id", oauth.GetClientId(c.Request))

	owner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}

	if oauth.GetClientId(c.Request) !=owner_id {
		log.Println("Owner id ", owner_id)
		log.Println("Client id ", oauth.GetClientId(c.Request));

		log.Println("Given token does not belong to the owner")
		c.JSON(http.StatusBadRequest, errors.New("Client Id does not match with given token id"))
		return
	}
	log.Println("===You are the owner of this partner..==")


	partner_domain := partners_domains.NewPartner()

	//Populate the partner with given user
	if err:= c.ShouldBindJSON(&partner_domain); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together

		restError := rest_errors_package.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}
	partner_domain.Id = owner_id


	isPartial := c.Request.Method == http.MethodPatch

	result, update_error := services.PartnerService.Update_Partner_Service(isPartial,*partner_domain)
	if update_error != nil {
		//TODO: Handle user creation Error
		c.JSON(update_error.Code, update_error)
		return
	}
	//fmt.Println("Partner Domain ",partner_domain)
	//c.String(http.StatusNotImplemented, "Implement Me!")
	//c.JSON(http.StatusOK, result)
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//


func GetSinglePartner(c *gin.Context) {

	if err:=oauth.AuthenticateRequest(c.Request);err!= nil {
		c.JSON(err.Code,err)
		return
	}

	log.Println("Caller Id", oauth.GetCallerId(c.Request))

	log.Println("Client Id", oauth.GetClientId(c.Request))

	owner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}

	if oauth.GetClientId(c.Request) !=owner_id {
		log.Println("Owner id ", owner_id)
		log.Println("Client id ", oauth.GetClientId(c.Request));

		log.Println("Given token does not belong to the owner")
		c.JSON(http.StatusBadRequest, errors.New("Client Id does not match with given token id"))
		return
	}
	log.Println("===You are the owner of this partner..==")
	result, get_error := services.PartnerService.Get_Partner_Service(owner_id)
	if get_error!= nil {
		c.JSON(get_error.Code,get_error)
		return
	}

	if oauth.GetCallerId(c.Request) == result.Id {
		c.JSON(http.StatusOK, result.Marshall(false))
		return
	}
	//c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
	fmt.Println(oauth.IsPublic(c.Request))
	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
	//c.JSON(http.StatusOK,result)
	//c.String(http.StatusNotImplemented, "Implement Me!")
}


func GetPartner(c *gin.Context) {

	if err:=oauth.AuthenticateRequest(c.Request);err!= nil {
		c.JSON(err.Code,err)
		return
	}

	log.Println("Caller Id", oauth.GetCallerId(c.Request))

	log.Println("Client Id", oauth.GetClientId(c.Request))

	partner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}
	result, get_error := services.PartnerService.Get_Partner_Service(partner_id)
	if get_error!= nil {
		c.JSON(get_error.Code,get_error)
		return
	}

	if oauth.GetCallerId(c.Request) == result.Id {
		c.JSON(http.StatusOK, result.Marshall(false))
		return
	}
	//c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
	fmt.Println(oauth.IsPublic(c.Request))
	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
	//c.JSON(http.StatusOK,result)
	//c.String(http.StatusNotImplemented, "Implement Me!")
}


func DeletePartner(c *gin.Context) {

	partner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}
	if errDelete:= services.PartnerService.Delete_Partner_Service(partner_id); errDelete!=nil {
		c.JSON(errDelete.Code, errDelete)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": fmt.Sprintf("deleted partner with id %d", partner_id)})
}

func FindPartnerByOwner(c *gin.Context) {
	if err:=oauth.AuthenticateRequest(c.Request);err!= nil {
		c.JSON(err.Code,err)
		return
	}
	owner_id, idError:=getOwnerId(c.Param("owner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}

	if oauth.GetClientId(c.Request) !=owner_id {
		log.Println("Owner id ", owner_id)
		log.Println("Client id ", oauth.GetClientId(c.Request));

		log.Println("Given token does not belong to the owner")
		c.JSON(http.StatusBadRequest, errors.New("Client Id does not match with given token id"))
		return
	}


	result, get_error := services.PartnerService.FindPartnerByOwner(owner_id);
	if get_error!=nil{
		c.JSON(get_error.Code, get_error)
		return
	}
	log.Println("Inside controller checing if this was a public request ")
	log.Println(oauth.IsPublic(c.Request))
	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))

	return
}

func FindByPartnerActive(c *gin.Context) {
	//
	//status, statusError:=getPartnerStatus(c.Param("status"))
	status, statusError:=getPartnerStatus(c.Query("status"))
	log.Println("Searching for ", status)
	if statusError!=nil{
		c.JSON(statusError.Code,statusError)
		return
	}
	statusPartners, errFindPartnerStatus:= services.FindByPartnerActive(status);
	if errFindPartnerStatus!=nil{
		c.JSON(errFindPartnerStatus.Code, errFindPartnerStatus)
		return
	}



	c.JSON(http.StatusOK,statusPartners.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request partners_domains.PartnerLoginRequest

	log.Println("Request to login with Email", request.Email_id);
	log.Println("Request to login with Password ", request.Password);

	if err:= c.ShouldBindJSON(&request); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together

		restError := rest_errors_package.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}
	partner, err := services.PartnerService.LoginUser(request)

	if err!=nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, partner.Marshall(c.GetHeader("X-Public") == "true"))

	return
}
func GetAllPartners(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!")
}
