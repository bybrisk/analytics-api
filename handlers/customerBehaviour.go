
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/analytics-api/data"
)

// swagger:route POST /analytics/track/behaviour analytics updateBehaviour
// Update behaviour of customer based on demand and supply towards a particular business.
//
// responses:
//	200: customerBehaviourResponse
//  422: errorValidation
//  501: errorResponse

func (p *Analytics) UpdateCustomerBehaviour (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> analytics-api Module")
	customer := &data.CustomerBehaviourRequest{}

	err:=customer.FromJSON(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = customer.ValidateCustomerBehaviourRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> analytics-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//update customer behaviour
	response := data.UpdateBehaviourCRUDOPS(customer)

	//writing to the io.Writer
	err = response.ResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}