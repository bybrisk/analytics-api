package data

import (
	//"github.com/bybrisk/structs"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
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