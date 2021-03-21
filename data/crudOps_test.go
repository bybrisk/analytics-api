package data_test

import (
	"testing"
	"fmt"
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	"github.com/bybrisk/analytics-api/data"
)

func TestUpdateBehaviourCRUDOPS(t *testing.T) {
	agentRequest := &data.CustomerBehaviourRequest{
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
		DeliveryID: "bybrisk123",
		Demand: 3,
		Supply: 5,
		PhoneNumber: "93402122380",
	}
	res:= data.UpdateBehaviourCRUDOPS(agentRequest) 
	fmt.Println(res)

}