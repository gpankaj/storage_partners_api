package partners_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/storage_partners_api/domains/partners_domains"
	"github.com/gpankaj/storage_partners_api/services"
	"github.com/gpankaj/storage_partners_api/utils/errors"
	"net/http"
	"strconv"
)

func CreatePartner(c *gin.Context) {
	var partner_domain partners_domains.Partner

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

	if err:= c.ShouldBindJSON(&partner_domain); err!= nil {
		//TODO: Handle unmarshal error + request data handling error together

		restError := errors.NewBadRequestError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}

	result, save_error := services.Create_Partner_Service(partner_domain)
	if save_error != nil {
		//TODO: Handle user creation Error
		c.JSON(save_error.Code, save_error)
		return
	}
	//fmt.Println("Partner Domain ",partner_domain)
	//c.String(http.StatusNotImplemented, "Implement Me!")
	c.JSON(http.StatusCreated, result)


}

func GetPartner(c *gin.Context) {
	partner_id, err := strconv.ParseInt(c.Param("partner_id"),10,64)
	if err!= nil {
		err:=errors.NewBadRequestError(fmt.Sprintf("Can not parse input text err: %s", err.Error()))
		c.JSON(err.Code, err)
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
	c.String(http.StatusNotImplemented, "Implement Me!")
}


func GetAllPartners(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!")
}