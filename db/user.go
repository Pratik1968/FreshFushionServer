package db

import (
	"context"
	"log"

	"github.com/PratikShekhar/FreshFusionServer/model"
)

func  InsertUser(userInfo *model.UserInfo) error{
 col  := client.Database("database").Collection("user")
 result, err := col.InsertOne(context.TODO(),userInfo)
 if err!=nil{
	return err
 }
log.Printf("Inserted document with _id: %v\n", result.InsertedID)	
return nil
}

func InsertUID(uid string) error{
	col  := client.Database("database").Collection("user_id")
	doc := model.UserID{User_id: uid}
	result, err := col.InsertOne(context.TODO(),doc)
	if err!=nil{
		return err
	 }
	log.Printf("Inserted document with _id: %v\n", result.InsertedID)	
	return nil

}