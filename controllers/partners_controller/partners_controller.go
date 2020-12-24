package partners_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/storage_partners_api/domains/partners_domains"
	"github.com/gpankaj/storage_partners_api/services"
	"github.com/gpankaj/storage_partners_api/utils/errors"
	"log"
	"net/http"
	"strconv"
)


func getPartnerId(partnerIdParams string) (int64, *errors.RestErr) {
	partner_id, partnerIdError := strconv.ParseInt(partnerIdParams,10,64)
	if partnerIdError!= nil {
		return 0, errors.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", partnerIdError.Error()))

	}
	return partner_id, nil
}

func getPartnerStatus(partnerActiveParams string) (bool, *errors.RestErr) {
	fmt.Println("Inside getPartnerStatus ",partnerActiveParams)
	status, statusError := strconv.ParseBool(partnerActiveParams)
	if statusError!= nil {
		//True is the default status.
		return true, errors.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", statusError.Error()))

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

		restError := errors.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}
	fmt.Println("Password is ", partner_domain)
	result, save_error := services.Create_Partner_Service(*partner_domain)
	if save_error != nil {
		//TODO: Handle user creation Error
		c.JSON(save_error.Code, save_error)
		return
	}
	//fmt.Println("Partner Domain ",partner_domain)
	//c.String(http.StatusNotImplemented, "Implement Me!")
	c.JSON(http.StatusCreated, result)


}

func UpdatePartner(c *gin.Context) {
	//create a partner domain.

	partner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}

	partner_domain := partners_domains.NewPartner()

	//Populate the partner with given user
	if err:= c.ShouldBindJSON(&partner_domain); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together

		restError := errors.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}
	partner_domain.Id = partner_id

	log.Println(partner_domain)

	isPartial := c.Request.Method == http.MethodPatch


	result, update_error := services.Update_Partner_Service(isPartial,*partner_domain)
	if update_error != nil {
		//TODO: Handle user creation Error
		c.JSON(update_error.Code, update_error)
		return
	}
	//fmt.Println("Partner Domain ",partner_domain)
	//c.String(http.StatusNotImplemented, "Implement Me!")
	c.JSON(http.StatusOK, result)
}
func GetPartner(c *gin.Context) {
	partner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}
	result, get_error := services.Get_Partner_Service(partner_id)
	if get_error!= nil {
		c.JSON(get_error.Code,get_error)
		return
	}

	c.JSON(http.StatusOK,result)

	//c.String(http.StatusNotImplemented, "Implement Me!")
}


func DeletePartner(c *gin.Context) {

	partner_id, idError:=getPartnerId(c.Param("partner_id"))
	if idError!=nil{
		c.JSON(idError.Code,idError)
		return
	}
	if errDelete:= services.Delete_Partner_Service(partner_id); errDelete!=nil {
		c.JSON(errDelete.Code, errDelete)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": fmt.Sprintf("deleted partner with id %d", partner_id)})
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
	c.JSON(http.StatusOK,statusPartners)
}

func GetAllPartners(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!")
}