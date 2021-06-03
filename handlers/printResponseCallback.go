
package handlers

import (
	"net/http"
	"github.com/bybrisk/analytics-api/data"
	"encoding/base64"
	"encoding/json"
)

// swagger:route GET /analytics/callback analytics analyseResponseOnSheetCallback
// Analyse all response of business on google sheet callback function.
//
// responses:
//	200: googleSheetAnalysisResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) AnalyseResponseOnSheet (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> analytics-api Module")

	state := r.URL.Query().Get("state")

	sDec, _ := base64.StdEncoding.DecodeString(state)

	request := data.GoogleSheetStructDir{}

	json.Unmarshal(sDec, &request)

	lp := data.PrintGoogleSheetCrudOps(request.Id,request.ActionHandler,r)

	err := lp.GoogleSpreadSheetMetaStructToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}