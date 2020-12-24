package app

import (
	"github.com/gpankaj/storage_partners_api/controllers/partners_controller"
	"github.com/gpankaj/storage_partners_api/controllers/pings_controller"
)

func mapUrls(){
	router.GET("/ping", pings_controller.Ping);

	router.POST("/partners", partners_controller.CreatePartner);
	router.GET("/partners/:partner_id", partners_controller.GetPartner);
	router.PUT("/partners/:partner_id", partners_controller.UpdatePartner);
	router.PATCH("/partners/:partner_id", partners_controller.UpdatePartner);
	router.DELETE("/partners/:partner_id", partners_controller.DeletePartner);

	router.GET("/internal/partners/search", partners_controller.FindByPartnerActive);

	router.GET("/partners", partners_controller.GetAllPartners);

}
