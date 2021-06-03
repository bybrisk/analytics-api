package data

import (
	//"github.com/bybrisk/structs"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var resultID string 

//Database Funcs
func UpdateCustomerBehaviour (payload *CustomerBehaviourRequest) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("analytics")
	//id, _ := primitive.ObjectIDFromHex(agents.BybID)
	options := options.Update()
	options.SetUpsert(true)
	filter := bson.M{"phoneNumber": payload.PhoneNumber}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.D{{Key: "$set", Value: payload}},options)
	if err != nil {
		log.Error("UpdateCustomerBehaviour ERROR:")
		log.Error(err)
	}
	fmt.Println(updateResult.ModifiedCount)
	return updateResult.ModifiedCount
}

func GetSheetIdAndURLMongo(docID string) (string,string){
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}

	type SheetIdAndLinkStruct struct{
		SheetID string `json:"sheetID"`
		SheetLink string `json:"sheetLink"`
	}

	var document SheetIdAndLinkStruct

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetGeocodes ERROR:")
		log.Error(err)
	}
	return document.SheetID, document.SheetLink
}