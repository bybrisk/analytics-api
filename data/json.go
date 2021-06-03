package data

import (
	"encoding/json"
	"io"
)	

func (d *CustomerBehaviourRequest) FromJSON (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}


func (d *CustomerBehaviourSuccess) ResultToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *GenerateGeocodeReportSuccess) GenerateGeocodeReportSuccessToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *UpdateGeocodeRequest) FromJSONToUpdateGeocodeRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *UpdateGeocodeSuccess) UpdateGeocodeSuccessToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *DeliveryStatusResponseAggregated) DeliveryStatusResponseAggregatedToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *GoogleSpreadSheetMetaStruct) GoogleSpreadSheetMetaStructToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}