package storage

import (
	"context"
	"fmt"
	"getir/handlerMessage"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"

var (
	dbInternal = make(map[string]string)
)

type InsertInternalDB struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RetrievedFromDB struct {
	Code    int            `json:"code"`
	Message string         `json:"msg"`
	List    []ListOfObject `json:"records"`
}

type ListOfObject struct {
	Key        string `json:"key"`
	CreatedAt  string `json:"createdAt"`
	TotalCount int    `json:"totalCount"`
}

type Storage interface {
	InsertInMemory(key string, value string) (InsertInternalDB, error)
	TakeFromMemory(key string) (string, error)
	TakeFromDB(startDate []string, endDate []string, minCount string, maxCount string) ([]ListOfObject, error)
}

type StorageImpl struct {
}

func NewStorageImpl() Storage {
	storageImpl := StorageImpl{}
	return &storageImpl
}

// service to add a new trainer into the system
func (storageImpl *StorageImpl) InsertInMemory(key string, value string) (InsertInternalDB, error) {
	var insertedValue InsertInternalDB
	if dbInternal[key] == "" {
		dbInternal[key] = value
		insertedValue.Key = key
		insertedValue.Value = value
		return insertedValue, nil
	} else {
		return insertedValue, handlerMessage.ErrDataAlreadyPresent
	}

}
func (storageImpl *StorageImpl) TakeFromMemory(key string) (string, error) {
	retrieved := dbInternal[key]
	if retrieved == "" {
		return "", handlerMessage.ErrDataNotFound
	}
	return retrieved, nil

}
func (storageImpl *StorageImpl) TakeFromDB(startDate []string, endDate []string, minCount string, maxCount string) ([]ListOfObject, error) {
	var listData []ListOfObject
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return listData, handlerMessage.ErrConnection
	}
	defer func() {
		client.Disconnect(context.TODO())
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return listData, handlerMessage.ErrCheckConnection
	}
	fmt.Println("Successfully connected and pinged.")

	collect := client.Database("getir-case-study").Collection("records")

	//filter := bson.D{{"key", "TAKwGc6Jr4i8Z487"}}
	//	fromDate := time.Date(YYYY, MM, DD, 0, 0, 0, 0, time.UTC)
	//toDate := time.Date(YYYY, MM, DD, 0, 0, 0, 0, time.UTC)
	yearStart, _ := strconv.Atoi(startDate[0])
	yearEnd, _ := strconv.Atoi(endDate[0])
	fromDate := time.Date(yearStart, time.November, 4, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(yearEnd, time.November, 5, 0, 0, 0, 0, time.UTC)
	filter := bson.M{
		"createdAt": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		}}

	cursor, err := collect.Find(context.Background(), filter)
	if err != nil {
		return listData, handlerMessage.ErrFindData
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		return listData, handlerMessage.ErrRetrieveData
	}
	if len(results) == 0 {
		return listData, handlerMessage.ErrDataNotFound
	}

	for _, result := range results {
		var data ListOfObject
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
		timeFinalType := time.Unix(timeConverted.Time().Unix(), 0).Format(time.RFC3339)
		fmt.Println(time.Unix(timeConverted.Time().Unix(), 0).Format(time.RFC3339))
		var sum int32
		for _, value := range array {
			sum += value.(int32)
		}
		data.TotalCount = int(sum)
		data.Key = response["key"].(string)
		data.CreatedAt = timeFinalType
		listData = append(listData, data)
		fmt.Println(sum)
	}
	defer cursor.Close(context.TODO())
	return listData, nil
}
