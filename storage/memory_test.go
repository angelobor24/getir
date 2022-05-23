package storage

import (
	"getir/handlerMessage"
	"testing"

	"gotest.tools/assert"
)

func TestInsertInMemory(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	storage := NewStorageImpl(compareFunc)
	insertedValue, err := storage.InsertInMemory("key", "value")
	assert.Equal(t, insertedValue.Key, "key")
	assert.Equal(t, insertedValue.Value, "value")
	assert.Equal(t, err == nil, true)
	insertedValue, err = storage.InsertInMemory("key", "value")
	assert.Equal(t, err, handlerMessage.ErrDataAlreadyPresent)
}

func TestTakeFromMemory(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	storage := NewStorageImpl(compareFunc)
	storage.InsertInMemory("key", "value")
	returnedValue, err := storage.TakeFromMemory("key")
	assert.Equal(t, returnedValue, "value")
	assert.Equal(t, err == nil, true)
	returnedValue, err = storage.TakeFromMemory("keyTest")
	assert.Equal(t, returnedValue, "")
	assert.Equal(t, err, handlerMessage.ErrDataNotFound)
}

func TestTakeFromDB(t *testing.T) {
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	storage := NewStorageImpl(compareFunc)
	_, err := storage.TakeFromDB([]string{"99999999", "12", "11"}, []string{"99999999", "12", "11"}, "23", "2000")
	assert.Equal(t, err, handlerMessage.ErrDataNotFound)
	_, err = storage.TakeFromDB([]string{"2010", "12", "11"}, []string{"2018", "12", "11"}, "23", "2000")
	assert.Equal(t, err == nil, true)
}
