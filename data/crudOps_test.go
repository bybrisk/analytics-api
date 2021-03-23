package data_test

import (
	"testing"
	"fmt"
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	"github.com/bybrisk/analytics-api/data"
)

/*func TestUpdateBehaviourCRUDOPS(t *testing.T) {
	agentRequest := &data.CustomerBehaviourRequest{
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
		DeliveryID: "bybrisk123",
		Demand: 3,
		Supply: 5,
		PhoneNumber: "93402122380",
	}
	res:= data.UpdateBehaviourCRUDOPS(agentRequest) 
	fmt.Println(res)

}*/

func TestUpdateGeocodeCRUDOPS (t *testing.T) {
	agentRequest := &data.UpdateGeocodeRequest{
		DeliveryID: "kqP2LHgB6eN9jSTIZGHi",
		Latitude: 23.25059567,
		Longitude: 77.4618567,
	}
	res:= data.UpdateGeocodeCRUDOPS(agentRequest) 
	fmt.Println(res)
}