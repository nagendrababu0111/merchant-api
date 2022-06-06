package main

import (
	"merchant-api/routers"
	"merchant-api/utils/config"

	// _ "go.mongodb.org/mongo-driver/bson"
	// "merchant-api/routers"
	// "merchant-api/utils/config"
	"merchant-api/utils/commons"
)

// type Merchant struct {
// 	Id          primitive.ObjectID `bson:"_id,omitempty`
// 	Code        string             `json:"code"`
// 	Category    string             `json:"category"`
// 	Description string             `json:"description"`
// 	VersionDate time.Time          `json:"version_date"`
// }

// // Team Member Schema
// type Member struct {
// 	EmailId      string    `json:"email_id"`
// 	MerchantCode string    `json:"merchant_code"`
// 	FirstName    string    `json:"first_name"`
// 	LastName     string    `json:"last_name"`
// 	VersionDate  time.Time `json:"version_date"`
// }

func main() {
	config.LoadConfigSettings()
	commons.InitIndexes()
	commons.SeedUsers()

	// bson.I
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Disconnect(ctx)

	// database := client.Database("quickstart")
	// podcastsCollection := database.Collection("member")
	// episodesCollection := database.Collection("merchant")
	// podcastsCollection.W
	routingEngine := routers.Init()
	routingEngine.Run()
}
