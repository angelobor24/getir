package server

import (
	"getir/storage"
	"strconv"
)

type Service interface {
	insertInMemory(key string, value string) (storage.InsertInternalDB, error)
	takeFromMemory(key string) (string, error)
	takeFromDB(startDate []string, endDate []string, minCount string, maxCount string) ([]storage.ListOfObject, error)
}

type ServiceImpl struct {
	Storage storage.Storage
}

// service implementer of the Service interface. This structure implements all the logicto handle the
// received server request
func NewServiceImpl(storage storage.Storage) Service {
	serviceImpl := ServiceImpl{Storage: storage}
	return &serviceImpl
}

func ValidateDate(date []string) bool {
	if len(date) != 3 {
		return false
	}
	year, err := strconv.Atoi(date[0])
	if err != nil {
		return false
	}
	if year < 1900 || year > 2999 {
		return false
	}
	month, err := strconv.Atoi(date[1])
	if err != nil {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	day, err := strconv.Atoi(date[2])
	if err != nil {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}
	return true
}

func validateCount(value string) bool {
	if value == "" {
		return false
	}
	_, err := strconv.Atoi(value)
	return err == nil

}

// service to add a new trainer into the system
func (serviceImpl *ServiceImpl) insertInMemory(key string, value string) (storage.InsertInternalDB, error) {
	return serviceImpl.Storage.InsertInMemory(key, value)
}
func (serviceImpl *ServiceImpl) takeFromMemory(key string) (string, error) {
	return serviceImpl.Storage.TakeFromMemory(key)
}
func (serviceImpl *ServiceImpl) takeFromDB(startDate []string, endDate []string, minCount string, maxCount string) ([]storage.ListOfObject, error) {
	return serviceImpl.Storage.TakeFromDB(startDate, endDate, minCount, maxCount)
}
