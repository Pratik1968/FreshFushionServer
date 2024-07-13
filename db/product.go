package db

import (
	"context"
	"log"
	"github.com/PratikShekhar/FreshFusionServer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProduct(id string) model.Product{
col:= client.Database("database").Collection("products")
objectId, err := primitive.ObjectIDFromHex(id)
if err != nil{
    log.Println("Invalid id")
}
query:= bson.D{{ Key: "_id", Value: objectId}}
var result model.Product
if err:=col.FindOne(context.TODO(),query).Decode(&result); err!=nil{
	panic(err)

}
return result

}
func GetSearchProduct(value string) [] model.Product{
	// TODO: use fuzzy search
	col:= client.Database("database").Collection("products")
	IndexModel := mongo.IndexModel{Keys: bson.D{{"name", "text"}}}
name, err := col.Indexes().CreateOne(context.TODO(), IndexModel)
if err != nil {
	panic(err)
}
log.Println("Name of index created: " + name)
	query  := bson.D{{"$text", bson.D{{"$search",value}}}}
var results [] model.Product
cursor, err := col.Find(context.TODO(), query)
if err != nil {
   panic(err)
}

if err = cursor.All(context.TODO(), &results); err != nil {
   panic(err)
}

return results
}
func GetTrendingProducts() []model.Product{
	col:= client.Database("database").Collection("products")
opts := options.Find().SetSort(bson.D{{"rating",-1}}).SetLimit(5) 
query := bson.D{{}}
cursor,err := col.Find(context.TODO(),query,opts)
if err!=nil{
	panic(err)
}
var results []model.Product
if err = cursor.All(context.TODO(), &results); err != nil {
	panic(err)
 }
 
 return results
 

}
func GetByCategory(category string)[] model.Product{
	col:= client.Database("database").Collection("products")
query:= bson.D{
				{Key: "category",Value: category},

			}
	cursor,err:= col.Find(context.TODO(),query)
	if err!=nil{
		panic(err)
	}
	var results [] model.Product
	if err= cursor.All(context.TODO(),&results); err!=nil{
		panic(err)
	}
return results

}
