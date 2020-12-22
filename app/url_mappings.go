package app

import (
	"github.com/gpankaj/storage_partners_api/controllers/partners_controller"
	"github.com/gpankaj/storage_partners_api/controllers/pings_controller"
)

func mapUrls(){
	router.GET("/ping", pings_controller.Ping);

	router.POST("/partners", partners_controller.CreatePartner);
	router.GET("/partners/:partner_id", partners_controller.GetPartner);
	router.DELETE("/partners", partners_controller.DeletePartner);

	router.GET("/partners", partners_controller.GetAllPartners);

}
