
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/analytics-api/data"
)

// swagger:route GET /analytics/report/customer/{businessID} analytics geoCodeReport
// Get updated excel report of customers by businessID.
//
// responses:
//	200: generateGeocodeReportResponse
//  422: errorValidation
//  501: errorResponse

func (p *Analytics) GenerateGeocodeReport(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> analytics-api Module")
	
	vars := mux.Vars(r)
	id := vars["businessID"]

	lp,file := data.GenerateGeocodeReportCRUDOPS(id)

	err := lp.GenerateGeocodeReportSuccessToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment"; filename="userInputData.xlsx")
	w.Header().Set("File-Name", "userInputData.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err := file.Write(w)
}