package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/analytics-api/data"
)

// swagger:route POST /analytics/track/address analytics updateGeocode
// Update geocode of customer address using GPS from App.
//
// responses:
//	200: updateGeocodeResponse
//  422: errorValidation
//  501: errorResponse

func (p *Analytics) UpdateGeocode (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> analytics-api Module")
	customer := &data.UpdateGeocodeRequest{}

	err:=customer.FromJSONToUpdateGeocodeRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = customer.ValidateUpdateGeocodeRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> analytics-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//update customer behaviour
	response := data.UpdateGeocodeCRUDOPS(customer)

	//writing to the io.Writer
	err = response.UpdateGeocodeSuccessToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}