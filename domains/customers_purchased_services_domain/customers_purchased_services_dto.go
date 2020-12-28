package customers_purchased_services_domain

type Customer_Purchased_Service struct {
	Id 									int64
	Service_name		 				string

	Service_rates						int64
	Service_discount_offered			int64
	Service_customer_paid_amouunt		string

	Service_agreement_terms				string

	Service_start_date					string
	Service_end_date					string
	Service_renew_date					string
	Service_duration					string

	Service_recurring					bool
	Service_recurring_period			string
	Service_recurring_cost				string

	Service_terminated					string

	Storage_service_location			string
	Storage_service_location_id			string

	Transport_service_from_location		string
	Transport_service_to_location		string

	Insurance_service_signed_agreement	string

	Service_customer_accepted			bool

	Service_customer_paid				int64

	Date_created 						string

	Comments							string
}