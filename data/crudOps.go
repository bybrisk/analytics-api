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