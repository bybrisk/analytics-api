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
	// Phone number of customer
	//
	DeliveryID string `json: "deliveryID"`

	// Response message of success or failure
	//
	Message string `json:"message"`
}

func (d *CustomerBehaviourRequest) ValidateCustomerBehaviourRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *UpdateGeocodeRequest) ValidateUpdateGeocodeRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}