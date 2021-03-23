package data

import ("time"
		"fmt")

func UpdateBehaviourCRUDOPS (d *CustomerBehaviourRequest) *CustomerBehaviourSuccess{
	//save data to database and return ID
	loc, _ := time.LoadLocation("Asia/Kolkata")
    now := time.Now().In(loc)
    fmt.Println("Location : ", loc, " Time : ", now)
	now.Format("2006-01-02")
	d.Date = now.String()
	d.RecievingScore = 6.98
	res := UpdateCustomerBehaviour(d)
	//sending response
	var response CustomerBehaviourSuccess
	if res == 1 {
		response = CustomerBehaviourSuccess{
			PhoneNumber: d.PhoneNumber,
			Message: "Behaviour added successfully with PhoneNumber",
		}
	} else if res == 0 {
		response = CustomerBehaviourSuccess{
			PhoneNumber: d.PhoneNumber,
			Message: "Behaviour added successfully with PhoneNumber",
		}
	} else {
		response = CustomerBehaviourSuccess{
			PhoneNumber: d.PhoneNumber,
			Message: "Some error occured!",
		}
	}

	return &response
}

func UpdateGeocodeCRUDOPS(d *UpdateGeocodeRequest) *UpdateGeocodeSuccess {
	var response UpdateGeocodeSuccess
	res := UpdateAddressGeocodeES(d)
	if res == "1" {
		response = UpdateGeocodeSuccess{
			DeliveryID: d.DeliveryID,
			Message: "Geocode updated successfully with DeliveryID",
		}
	} else {
		response = UpdateGeocodeSuccess{
			DeliveryID: d.DeliveryID,
			Message: "Some error occured!",
		}
	}
	
	return &response
}