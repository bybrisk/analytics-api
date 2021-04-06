package data

import (
	"encoding/json"
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
	"bytes"
	//"context"
	"io/ioutil"
	"net/http"
	
)

var Elasticurl string = "https://search-krayfin-ewgnw3zevhpuznh2kvo5hbbmtq.us-east-2.es.amazonaws.com"
var UsernameElastic string = "elastic"
var Elasticpassword string = "K8txmFf6hwnGvaNs7HxNcg2w$"
var urlAuthenticate string = "https://elastic:K8txmFf6hwnGvaNs7HxNcg2w$@search-krayfin-ewgnw3zevhpuznh2kvo5hbbmtq.us-east-2.es.amazonaws.com"
var awsYearIndex string = "/*2021"

var (
	clusterURLs = []string{Elasticurl}
	username    = UsernameElastic
	password    = Elasticpassword
  )

func UpdateAddressGeocodeES(d *UpdateGeocodeRequest) string{
	var id string

	//get todays date
	loc, _ := time.LoadLocation("Asia/Kolkata")
    now := time.Now().In(loc)
    fmt.Println("Location : ", loc, " Date And Time of delivery : ", now)
	t:=now.Format("01-02-2006")
	dodATime := t

	//Encode the data
	postBody:=`{
		"script" : {
			"source": "ctx._source.latitude=`+fmt.Sprintf("%f", d.Latitude)+`;ctx._source.longitude=`+fmt.Sprintf("%f", d.Longitude)+`;ctx._source.dateOfDelivery='`+dodATime+`';",
			"lang": "painless"  
		  },
		  "query": {
			  "ids" : {
			"values" : "`+d.DeliveryID+`"
			}
		  }
	  }`

	 responseBody := bytes.NewBufferString(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(urlAuthenticate+awsYearIndex+"/_update_by_query?conflicts=proceed", "application/json", responseBody)
  
	 //Handle Error
	 if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()

	var r map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
    	log.Printf("Error parsing the response body: %s", err)
    } else {
    	// Print the response status and indexed document version.
		id=fmt.Sprintf("%v", r["updated"])
    }

	fmt.Println(id)
	return id	
}  

func GetGeocodeDataES(docID string) *DeliveryResponseBulk{
	var deliveries DeliveryResponseBulk

	postBody:=`{
		"query": {
		  "bool": {
			"filter": [
			  {"term": {"BybID": "`+docID+`"}},
			  {"term" : {"deliveryStatus.keyword" : "Delivered" }}
			]
		  }
		}
	  }`

	 responseBody := bytes.NewBufferString(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(urlAuthenticate+"/_all/_search?size=5000", "application/json", responseBody)
  
	 //Handle Error
	 if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()

	 body, err := ioutil.ReadAll(resp.Body)
	 if err != nil {
		log.Error("ReadAll ERROR : ")
		log.Error(err)
	 }
	 
	 err = json.Unmarshal(body, &deliveries)
	 if err != nil {
		log.Error("json.Unmarshal ERROR : ")
		log.Error(err)
    	}	
	return &deliveries
}