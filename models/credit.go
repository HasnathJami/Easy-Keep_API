package models

type Credit struct {
	//CreditId  primitive.ObjectID    `bson:"credit_id"`
	AccountNo uint                  `json:"account_no"`
	UserName  string                `json:"user_name"`
	Amount    float64               `json:"amount"`
	Date      string                `json:"date"`
}

// func CreditMoney(ctx context.Context, credit *Credit) (*mongo.InsertOneResult, error) {
// 	var entryCollection *mongo.Collection = database.DBInstance()
// 	result, insertError := entryCollection.InsertOne(ctx, credit)

// 	return result, insertError
// }