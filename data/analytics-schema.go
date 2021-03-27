package data

import (
	"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
)

//post request to insert behavioural pattern
type CustomerBehaviourRequest struct{

	// Id of Delivery (BybID of Delivery)
	//
	// required: false
	// max length: 1000
	DeliveryID string `json: "deliveryID"`

	// BusinessID of the business (BybID of business)
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessID" validate:"required"`

	// Phone number of customer
	//
	// required: true
	PhoneNumber string `json:"phoneNumber" validate:"required"`

	// demand of the customer
	//
	// required: true
	// max length: 1000
	Demand int64 `json: "demand" validate:"required"`

	// supply to the customer
	//
	// required: true
	// max length: 1000
	Supply int64 `json: "supply" validate:"required"`

	// Recieving score of customer (no need to provide)
	//
	// required: false
	// max length: 1000
	RecievingScore float64 `json: "recievingScore"`

	// Date of registering the behaviour (no need to provide)
	//
	// required: false
	// max length: 1000
	Date string `json: "date"`
}

//post request to insert behavioural pattern
type UpdateGeocodeRequest struct{

	// Id of Delivery (BybID of Delivery)
	//
	// required: true
	// max length: 1000
	DeliveryID string `json: "deliveryID" validate:"required"`

	// Latitude from GPS
	//
	// required: true
	Latitude float64 `json:"latitude" validate:"required"`

	// Longitude from GPS
	//
	// required: true
	Longitude float64 `json:"longitude" validate:"required"`
}

//post response
type CustomerBehaviourSuccess struct {
	// Phone number of customer
	//
	PhoneNumber string `json:"phoneNumber"`

	// Response message of success or failure
	//
	Message string `json:"message"`
}

//Geocode post response
type UpdateGeocodeSuccess struct {
	// DeliveryID of the address
	//
	DeliveryID string `json: "deliveryID"`

	// Response message of success or failure
	//
	Message string `json:"message"`
}

//Generate Geocode report response
type GenerateGeocodeReportSuccess struct {
	// BybID of the business
	//
	BybID string `json: "bybid"`

	// Response message of success or failure
	//
	Message string `json:"message"`
}

//get all deliveries Response struct
type DeliveryResponseBulk struct {
	Hits struct {
		Hits []struct {
			//Date of delivery
			//
			Index  string `json:"_index"`

			//ID of delivery
			//
			ID     string `json:"_id"`

			//Delivery details
			//
			Source struct {
				//Pincode of delivery location
				//
				Pincode         string  `json:"pincode"`

				//Latitude of delivery location
				//
				Latitude        float64 `json:"latitude"`

				//Phone number of the customer placing delivery
				//
				Phone           string  `json:"phone"`

				//Name of the customer placing delivery
				//
				CustomerName    string  `json:"CustomerName"`

				//Weight of Item delivered
				//
				ItemWeight      float64     `json:"itemWeight"`

				//payment to collect
				//
				Amount   float64    `json:"amount"`

				//Address of delivery
				//
				CustomerAddress string  `json:"CustomerAddress"`

				//Longitude of delivery location
				//
				Longitude       float64 `json:"longitude"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (d *CustomerBehaviourRequest) ValidateCustomerBehaviourRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *UpdateGeocodeRequest) ValidateUpdateGeocodeRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}