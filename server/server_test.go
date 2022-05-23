package server

import (
	"bytes"
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
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=23&maxCount=19"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusCreated)
}

func TestRetrieveDBErrorMinCount(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=a&maxCount=19"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBErrorMaxCount(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=23&maxCount=a"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBErrorStartDate(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-33&endDate=2023-10-31&minCount=23&maxCount=23"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBErrorEndDate(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-33&minCount=23&maxCount=23"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBErrorPathNotFound(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retr"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestRetrieveDBErrorDB(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=23&maxCount=19"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, handlerMessage.ErrRetrieveData)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestRetrieveDBMethodNotSupported(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodHead, ("/database/retrieve?startDate=2023-12-13&endDate=2023-10-31&minCount=23&maxCount=19"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.retrieveDB(w, req)
	assert.Equal(t, w.Code, http.StatusMethodNotAllowed)
}

func TestHandleMemoryNotFound(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/memory?key=test"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "", listData, handlerMessage.ErrDataNotFound)
	server := NewServerImplTest(serviceMock)
	server.handleMemory(w, req)
	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestHandleMemory(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ("/memory?key=test"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "value", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.handleMemory(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestHandleMemoryMethodNotAllowed(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodHead, ("/memory?key=test"), nil)
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "value", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.handleMemory(w, req)
	assert.Equal(t, w.Code, http.StatusMethodNotAllowed)
}

func TestHandleMemoryPostOk(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	insertedValue.Key = "test"
	insertedValue.Value = "test"
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ("/memory"), bytes.NewBuffer([]byte(`{"key":"test","value":"test"}`)))
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "test", listData, nil)
	server := NewServerImplTest(serviceMock)
	server.handleMemory(w, req)
	assert.Equal(t, w.Code, http.StatusCreated)
}

func TestHandleMemoryPostError(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	insertedValue.Key = "test"
	insertedValue.Value = "test"
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ("/memory"), bytes.NewBuffer([]byte(`{"key":"test","value":"test"}`)))
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "test", listData, handlerMessage.ErrDataAlreadyPresent)
	server := NewServerImplTest(serviceMock)
	server.handleMemory(w, req)
	assert.Equal(t, w.Code, http.StatusConflict)
}

func TestHandleMemoryPostErrorPath(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	var insertedValue storage.InsertInternalDB
	insertedValue.Key = "test"
	insertedValue.Value = "test"
	var listData []storage.ListOfObject
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ("/memoryTEST"), bytes.NewBuffer([]byte(`{"key":"test","value":"test"}`)))
	assert.NilError(t, err)
	w := httptest.NewRecorder()
	storage := storage.NewStorageImpl(compareFunc)
	serviceMock := NewMockService(storage, insertedValue, "test", listData, handlerMessage.ErrDataAlreadyPresent)
	server := NewServerImplTest(serviceMock)
	server.handleMemory(w, req)
	assert.Equal(t, w.Code, http.StatusNotFound)
}
