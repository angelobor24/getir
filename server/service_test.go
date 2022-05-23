package server

import (
	"getir/storage"
	"testing"

	"gotest.tools/assert"
)

func TestInsertInMemory(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	storageMock := storage.NewMockStorage(insertedValue, "", listData, nil)
	service := NewServiceImpl(storageMock)
	_, err := service.insertInMemory("", "")
	assert.Equal(t, err == nil, true)
}

func TestTakeFromMemory(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	storageMock := storage.NewMockStorage(insertedValue, "", listData, nil)
	service := NewServiceImpl(storageMock)
	_, err := service.takeFromMemory("")
	assert.Equal(t, err == nil, true)
}

func TestTakeFromDB(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	storageMock := storage.NewMockStorage(insertedValue, "", listData, nil)
	service := NewServiceImpl(storageMock)
	_, err := service.takeFromDB([]string{"2020", "12", "11"}, []string{"2022", "12", "11"}, "23", "2000")
	assert.Equal(t, err == nil, true)
}
