package server

import (
	"context"
	"getir/handlerMessage"
	"getir/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func NewServerImplTest(service Service) ServerImpl {
	serverImpl := ServerImpl{service: service}
	return serverImpl
}

func TestRetrieveDB(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=23&maxCount=19"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl()
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusCreated)
}

func TestRetrieveDBErrorMinCount(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=a&maxCount=19"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl()
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBErrorMaxCount(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=23&maxCount=a"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl()
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBErrorStartDate(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-33&endDate=2023-10-31&minCount=23&maxCount=23"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl()
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBErrorEndDate(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-33&minCount=23&maxCount=23"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl()
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBErrorPathNotFound(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retr"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl()
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestRetrieveDBErrorDB(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=23&maxCount=19"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl()
	serviceMock := NewMockService(storage, insertedValue, "", listData, handlerMessage.ErrRetrieveData)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBMethodNotSupported(t *testing.T) {
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodHead, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=23&maxCount=19"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl()
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusMethodNotAllowed)
}
