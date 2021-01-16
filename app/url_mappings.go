package app

import (
	"github.com/gpankaj/storage_partners_api/controllers/branches_controller"
	"github.com/gpankaj/storage_partners_api/controllers/customers_controller"
	"github.com/gpankaj/storage_partners_api/controllers/partners_controller"
	"github.com/gpankaj/storage_partners_api/controllers/pings_controller"
)

func mapUrls(){
	router.GET("/ping", pings_controller.Ping);

	router.POST("/partners", partners_controller.CreatePartner);
	router.GET("/partners/:partner_id", partners_controller.GetSinglePartner);
	//router.GET("/partners/:partner_id", partners_controller.GetPartner);


	router.PUT("/partners/:partner_id", partners_controller.UpdatePartner);
	router.PATCH("/partners/:partner_id", partners_controller.UpdatePartner);
	router.DELETE("/partners/:partner_id", partners_controller.DeletePartner);


	router.GET("/internal/partners/search", partners_controller.FindByPartnerActive);

	router.GET("/internal/partners/owner/:owner_id", partners_controller.FindPartnerByOwner);


	router.POST("/partners/login", partners_controller.Login);

	router.GET("/partners", partners_controller.GetAllPartners);

	router.POST("/branches", branches_controller.CreateBranch);
	router.GET("/branches", branches_controller.FindBranches);
	router.PATCH("/branches/:partner_id/:branch_id", branches_controller.UpdateBranch);
	router.DELETE("/branches/:partner_id/:branch_id", branches_controller.DeleteBranch);

	router.GET("/branches/:partner_id", branches_controller.FindPartnerBranches);


	router.POST("/customers",customers_controller.CreateCustomer)
	router.POST("/customers/login", customers_controller.Login);
}
