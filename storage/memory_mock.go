package storage

type MockMemory struct {
	inserted    InsertInternalDB
	retrieved   string
	listElement []ListOfObject
	err         error
}

func NewMockStorage(inserted InsertInternalDB, retrieved string, listElement []ListOfObject, err error) Storage {
	serviceMock := MockMemory{inserted: inserted, retrieved: retrieved, listElement: listElement, err: err}
	return &serviceMock
}

func (mockServiceImpl *MockMemory) InsertInMemory(key string, value string) (InsertInternalDB, error) {
	return mockServiceImpl.inserted, mockServiceImpl.err
}

func (mockServiceImpl *MockMemory) TakeFromMemory(key string) (string, error) {
	return mockServiceImpl.retrieved, mockServiceImpl.err
}

func (mockServiceImpl *MockMemory) TakeFromDB(startDate []string, endDate []string, minCount string, maxCount string) ([]ListOfObject, error) {
	return mockServiceImpl.listElement, mockServiceImpl.err
}
