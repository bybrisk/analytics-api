package data

import ("time"
		"fmt"
		"strconv"
		"github.com/360EntSecGroup-Skylar/excelize/v2"
		"net/http")

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

func GenerateGeocodeReportCRUDOPS(docID string,status string,date string) (*GenerateGeocodeReportSuccess,*excelize.File){
	var response GenerateGeocodeReportSuccess
	
	res := GetGeocodeDataES(docID,status,date)
	file:=GenerateExcelFile(res)
	//fmt.Println(res)
	response = GenerateGeocodeReportSuccess{
		BybID: docID,
		Message: "Report Successfully generated",
	}
	return &response,file

}

func GenerateExcelFile(res *DeliveryResponseBulk) *excelize.File {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Customer Name")
	f.SetCellValue("Sheet1", "B1", "Phone")
	f.SetCellValue("Sheet1", "C1", "Locality")
	f.SetCellValue("Sheet1", "D1", "Landmark")
	f.SetCellValue("Sheet1", "E1", "City")
	f.SetCellValue("Sheet1", "F1", "Amount")
	f.SetCellValue("Sheet1", "G1", "Item Weight")
	f.SetCellValue("Sheet1", "H1", "Pincode")
	f.SetCellValue("Sheet1", "I1", "Latitude")
	f.SetCellValue("Sheet1", "J1", "Longitude")

	cellAddress:="A1"
	for key,val := range res.Hits.Hits{
		keyString := strconv.Itoa(key+2)
		
		cellAddress="A"+keyString
		f.SetCellValue("Sheet1", cellAddress, val.Source.CustomerName)
	
		cellAddress="B"+keyString
		f.SetCellValue("Sheet1", cellAddress, val.Source.Phone)
		
		cellAddress="C"+keyString
		f.SetCellValue("Sheet1", cellAddress, val.Source.CustomerAddress)
		
		cellAddress="D"+keyString
		f.SetCellValue("Sheet1", cellAddress, ".")

		cellAddress="E"+keyString
		f.SetCellValue("Sheet1", cellAddress, "Bhopal")

		cellAddress="F"+keyString
		f.SetCellValue("Sheet1", cellAddress, val.Source.Amount)

		cellAddress="G"+keyString
		f.SetCellValue("Sheet1", cellAddress, val.Source.ItemWeight)

		cellAddress="H"+keyString
		f.SetCellValue("Sheet1", cellAddress, val.Source.Pincode)

		cellAddress="I"+keyString
		f.SetCellValue("Sheet1", cellAddress, val.Source.Latitude)

		cellAddress="J"+keyString
		f.SetCellValue("Sheet1", cellAddress, val.Source.Longitude)
	}

	//save excel file
	//if err := f.SaveAs("/home/shashank/Downloads/inventions/bybrisk/bybrisk map/customer data/Geocode Auto Generated Report.xlsx"); err != nil {
	//		println(err.Error())
	//	}
	return f
}

func GetUpdatedDeliveryStatusCRUDOPS(docID string) *DeliveryStatusResponseAggregated{
	var response DeliveryStatusResponseAggregated

	resDelivered := FetchDeliveryStatusES(docID,"Delivered")
	resCancelled := FetchDeliveryStatusES(docID,"Cancelled")
	resPending := FetchDeliveryStatusES(docID,"Pending")
	resTransit := FetchDeliveryStatusES(docID,"Transit")

	response = DeliveryStatusResponseAggregated{
		Delivered:resDelivered.Hits.Total.Value ,
		Cancelled:resCancelled.Hits.Total.Value ,
		Pending:resPending.Hits.Total.Value ,
		Transit:resTransit.Hits.Total.Value,
		Message: "Status fetched successfully",
	}

	return &response
}

func PrintGoogleSheetCrudOps (businessID string,actionHandler string, r *http.Request) *GoogleSpreadSheetMetaStruct {
	var response GoogleSpreadSheetMetaStruct
	
	sheetID, sheetLink := PrintOrderToShareGoogleAPI(businessID,actionHandler,r)
	
	response = GoogleSpreadSheetMetaStruct{
		SpreadsheetId: sheetID,
		SpreadsheetUrl: sheetLink,
		Message: "Successfully updated google sheet!",
		Status: 200, 
	}
	
	return &response
}