// Package classification of Analytics API
//
// Documentation for Analytics API
//
//	Schemes: https
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/analytics-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// Success message on updating a customer behaviour
// swagger:response customerBehaviourResponse
type accountPostResponseWrapper struct {
	// Success message on updating customer behaviour towards a business
	// in: body
	Body data.CustomerBehaviourSuccess
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateBehaviour
type updateBehaviourParamsWrapper struct {
	// data structure to update customer behaviour.
	// Note: the id field is ignored by create operations
	// in: body
	// required: true
	Body data.CustomerBehaviourRequest
}