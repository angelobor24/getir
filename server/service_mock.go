package server

import "getir/storage"

type MockService struct {
	inserted    storage.InsertInternalDB
	retrieved   string
	listElement []storage.ListOfObject
	err         error
}

func NewMockService(storage storage.Storage, inserted storage.InsertInternalDB, retrieved string, listElement []storage.ListOfObject, err error) Service {
	serviceMock := MockService{inserted: inserted, retrieved: retrieved, listElement: listElement, err: err}
	return &serviceMock
}

func (mockServiceImpl *MockService) insertInMemory(key string, value string) (storage.InsertInternalDB, error) {
	return mockServiceImpl.inserted, mockServiceImpl.err
}

func (mockServiceImpl *MockService) takeFromMemory(key string) (string, error) {
	return mockServiceImpl.retrieved, mockServiceImpl.err
}

func (mockServiceImpl *MockService) takeFromDB(startDate []string, endDate []string, minCount string, maxCount string) ([]storage.ListOfObject, error) {
	return mockServiceImpl.listElement, mockServiceImpl.err
}
