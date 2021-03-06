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

// local DB handled with a map
var (
	dbInternal = make(map[string]string)
)

type Convert func(int32, int32, int32) bool

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
	TotalCount int32  `json:"totalCount"`
}

type Storage interface {
	InsertInMemory(key string, value string) (InsertInternalDB, error)
	TakeFromMemory(key string) (string, error)
	TakeFromDB(startDate []string, endDate []string, minCount string, maxCount string) ([]ListOfObject, error)
}

type StorageImpl struct {
	acceptValue Convert
}

func NewStorageImpl(inputFunc func(int32, int32, int32) bool) Storage {
	storageImpl := StorageImpl{acceptValue: inputFunc}
	return &storageImpl
}

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

	yearStart, _ := strconv.Atoi(startDate[0])
	yearEnd, _ := strconv.Atoi(endDate[0])
	monthStart, _ := strconv.Atoi(startDate[1])
	monthEnd, _ := strconv.Atoi(endDate[1])
	dayStart, _ := strconv.Atoi(startDate[2])
	dayEnd, _ := strconv.Atoi(endDate[2])
	fromDate := time.Date(yearStart, time.Month(monthStart), dayStart, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(yearEnd, time.Month(monthEnd), dayEnd, 0, 0, 0, 0, time.UTC)
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
		response := result.Map()
		interfaceArray := response["counts"]
		array := interfaceArray.(primitive.A)
		timeRetrieved := response["createdAt"]
		timeConverted := timeRetrieved.(primitive.DateTime)
		timeFinalType := time.Unix(timeConverted.Time().Unix(), 0).Format(time.RFC3339)
		var sum int32
		for _, value := range array {
			sum += value.(int32)
		}
		timeConvertedMin, _ := strconv.Atoi(minCount)
		timeConvertedMax, _ := strconv.Atoi(maxCount)
		if storageImpl.acceptValue(int32(timeConvertedMin), int32(timeConvertedMax), sum) {
			data.TotalCount = sum
			data.Key = response["key"].(string)
			data.CreatedAt = timeFinalType
			listData = append(listData, data)
		}

	}
	defer cursor.Close(context.TODO())
	return listData, nil
}
