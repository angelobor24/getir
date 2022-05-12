package main

import (
	"context"
	"fmt"
	"getir/server"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"

func main() {
	testConnect()
	server := server.NewServerImpl(server.NewServiceImpl())
	server.StartServer()
}

func testConnect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	collect := client.Database("getir-case-study").Collection("records")

	//filter := bson.D{{"key", "TAKwGc6Jr4i8Z487"}}
	//	fromDate := time.Date(YYYY, MM, DD, 0, 0, 0, 0, time.UTC)
	//toDate := time.Date(YYYY, MM, DD, 0, 0, 0, 0, time.UTC)
	fromDate := time.Date(2014, time.November, 4, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2020, time.November, 5, 0, 0, 0, 0, time.UTC)
	filter := bson.M{
		"createdAt": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		}}

	cursor, err := collect.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	/*for cursor.Next(context.TODO()) {
		var result bson.D
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}*/

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	/*for _, result := range results {
		fmt.Println(result)
	}*/
	for _, result := range results {
		fmt.Println(result)
		response := result.Map()
		fmt.Println(response["key"])
		fmt.Println(response["createdAt"])
		fmt.Println(response["counts"])
		fmt.Println(response["value"])
		interfaceArray := response["counts"]
		array := interfaceArray.(primitive.A)
		timeRetrieved := response["createdAt"]
		timeConverted := timeRetrieved.(primitive.DateTime)
		//fmt.Println(timeConverted.Time().Date())
		//fmt.Println(timeConverted.Time().Clock())
		//2017-01-28T02:22:14+01:00
		fmt.Println(time.Unix(timeConverted.Time().Unix(), 0).Format(time.RFC3339))
		var sum int32
		for _, value := range array {
			sum += value.(int32)
		}
		fmt.Println(sum)
		//timeConverted.Time().After()
	}
	defer cursor.Close(context.TODO())

}
