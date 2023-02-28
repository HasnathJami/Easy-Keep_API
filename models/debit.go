package models

type Debit struct {
	//DebitId   primitive.ObjectID    `bson:"debit_id"`
	AccountNo uint    `json:"account_no"`
	UserName  string  `json:"user_name"`
	Amount    float64 `json:"amount"`
	Date      string  `json:"date"`
}

// func DebitMoney(ctx context.Context, debit *Debit) (*mongo.InsertOneResult, error) {
// 	var entryCollection *mongo.Collection = database.DBInstance()
// 	result, insertError := entryCollection.InsertOne(ctx, debit)

// 	return result, insertError
// }