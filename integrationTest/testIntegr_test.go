package integrationtest

import (
	"bytes"
	"encoding/json"
	"getir/server"
	"getir/storage"
	"net/http"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestFT1(t *testing.T) {
	//initialize compare function
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	//initialize server data structure
	server := server.NewServerImpl(server.NewServiceImpl(storage.NewStorageImpl(compareFunc)))
	//start server
	go server.StartServer()
	time.Sleep(2 * time.Second)
	newInternalElement := storage.InsertInternalDB{Key: "test", Value: "valueTest"}
	postBody, _ := json.Marshal(newInternalElement)
	responseBody := bytes.NewBuffer(postBody)
	// check insert into local DB
	_, err := http.Post("http://127.0.0.1:8080/memory", "application/json", responseBody)
	assert.Equal(t, err, nil)
	time.Sleep(2 * time.Second)
	postBody, _ = json.Marshal(newInternalElement)
	responseBody = bytes.NewBuffer(postBody)
	// check retrieve data
	resp, err := http.Get("http://127.0.0.1:8080/memory?key=test")
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	resp, err = http.Get("http://127.0.0.1:8080/memory?key=testTEST")
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusNotFound)

	resp, err = http.Get("http://127.0.0.1:8080/database/retrieve?startDate=2015-12-12&endDate=2022-12-12&minCount=10&maxCount=100")
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	resp, err = http.Get("http://127.0.0.1:8080/database/retrieve?startDate=9999-12-12&endDate=9999-12-12&minCount=10&maxCount=100")
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)

	resp, err = http.Get("http://127.0.0.1:8080/database/retrieve?startDate=2050-12-12&endDate=2050-12-12&minCount=10&maxCount=100")
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusNotFound)
}
