package models

import (
	"fmt"
	"context"
	"time"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	

)

type User struct{
	Name string
	Age int
	Address string
}


func getClient() (*mongo.Client, context.Context){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}


func InsertOne(filter interface{}){
	client,ctx := getClient()
	collection := client.Database("signoi").Collection("users")

	// user := User{Name, Age, Address}

	insertResult, err := collection.InsertOne(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func ReadOne(filter interface{}) User{
	var result User
	client,ctx := getClient()
	collection := client.Database("signoi").Collection("users")

	fmt.Println(reflect.TypeOf(filter))
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result

}

func ReadAll(filter interface{}) []User{
	var results []User
	client,ctx := getClient()
	collection := client.Database("signoi").Collection("users")
	findOptions := options.Find()

	cur,err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
	
		results = append(results, elem)
	}
	
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	return results

}

func DeleteOne(filter interface{}){
	client,ctx := getClient()
	collection := client.Database("signoi").Collection("users")
	
	deleteOptions := options.Delete()
	_,err := collection.DeleteOne(ctx, filter,deleteOptions)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateOne(filter interface{}, update interface{}){
	client,ctx := getClient()
	collection := client.Database("signoi").Collection("users")
	
	_, err := collection.UpdateOne(
        ctx,
        filter,
        update,
	)
	if err != nil {
		log.Fatal(err)
	}
}