package models

import (
	"context"
	"fmt"
	"log"

	"github.com/HasnathJami/Easy-Keep/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Report struct {
	ReportId    primitive.ObjectID `bson:"report_id"`
	//UserId      string             `json:"user_id"`
	//UserName    string             `json:"user_name"`
	//AccountNo   string             `json:"account_no"`
	TotalAmount   float64               `json:"total_amount"`
	//Date        string             `json:"date"`
}

func UpdateReport(ctx context.Context, report *Report) (*mongo.UpdateResult, error) {
	var entryCollection *mongo.Collection = database.DBInstance()
	fmt.Println("report data: ",report)
	//result, updateError := entryCollection.InsertOne(ctx, report)
	//result, updateError := entryCollection.UpdateOne(ctx, bson.M{"_id": docID},
	result, updateError := entryCollection.ReplaceOne(ctx,
		bson.M{},
		bson.M{
			"total_amount" : report.TotalAmount,
		},
	)

	return result, updateError
}

func GetReport(ctx context.Context) ([]primitive.M, error,error){
	var entryCollection *mongo.Collection = database.DBInstance()
	
    var reports []bson.M
	cursor, err1 := entryCollection.Find(ctx, bson.M{})

	//utils.CheckError(err,err.Error(), http.StatusInternalServerError, w)
	if err1 != nil{
		log.Fatal(err1) 
	}

	err2 :=  cursor.All(ctx, &reports)
	if err2 != nil{
		log.Fatal(err2)
	}
	//utils.CheckError(err,err.Error(), http.StatusInternalServerError, w)
	return reports, err1, err2

}

func CreateReport(ctx context.Context, report *Report) (*mongo.InsertOneResult, error) {
	var entryCollection *mongo.Collection = database.DBInstance()
	result, insertError := entryCollection.InsertOne(ctx, report)

	return result, insertError
}