package data

import (
	"encoding/json"
	"fmt"
	//"time"
	log "github.com/sirupsen/logrus"
	"bytes"
	//"context"
	"net/http"
	//"github.com/elastic/go-elasticsearch/v8"
	//"github.com/elastic/go-elasticsearch/v8/esutil"
	//"github.com/mitchellh/mapstructure"
	
)

var Elasticurl string = "https://0e38a62eadcb475db131ebdf2152e96d.ap-south-1.aws.elastic-cloud.com:9243"
var UsernameElastic string = "elastic"
var Elasticpassword string = "K8txmFf6hwnGvaNs7HxNcg2w"
var urlAuthenticate string = "https://elastic:K8txmFf6hwnGvaNs7HxNcg2w@0e38a62eadcb475db131ebdf2152e96d.ap-south-1.aws.elastic-cloud.com:9243"

var (
	clusterURLs = []string{Elasticurl}
	username    = UsernameElastic
	password    = Elasticpassword
  )

func UpdateAddressGeocodeES(d *UpdateGeocodeRequest) string{
	var id string
	//Encode the data
	postBody:=`{
		"script" : {
			"source": "ctx._source.latitude=`+fmt.Sprintf("%f", d.Latitude)+`;ctx._source.longitude=`+fmt.Sprintf("%f", d.Longitude)+`;",
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
	 resp, err := http.Post(urlAuthenticate+"/_all/_update_by_query?conflicts=proceed", "application/json", responseBody)
  
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