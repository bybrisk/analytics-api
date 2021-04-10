
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/analytics-api/data"
)

// swagger:route GET /analytics/delivery/status/{businessID} analytics getDeliveryStatus
// Get updated delivery status of an account by businessID.
//
// responses:
//	200: getDeliveryStatusResponse
//  422: errorValidation
//  501: errorResponse

func (p *Analytics) GetUpdatedDeliveryStatus(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> analytics-api Module")
	
	vars := mux.Vars(r)
	id := vars["businessID"]

	lp := data.GetUpdatedDeliveryStatusCRUDOPS(id)

	err := lp.DeliveryStatusResponseAggregatedToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}