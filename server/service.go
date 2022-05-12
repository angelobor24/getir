package server

type TrainerInfo struct {
	Name      string `json:"name" validate:"required,alphaunicode"`
	Surname   string `json:"surname" validate:"required,alphaunicode"`
	Idtrainer int    `json:"idtrainer" validate:"required"`
	Mt        int    `json:"mt" `
}

type Quote struct {
	Pokemon string  `json:"pokemon" validate:"required"`
	Price   float32 `json:"price" `
	Id      int     `json:"id" `
}

type PayedQuote struct {
	Pokemon       string  `json:"pokemon" validate:"required"`
	Price         float32 `json:"price" `
	IdTrainer     int     `json:"idTrainer" `
	IdTransaction int64   `json:"idTransaction" `
	Timestamp     string  `json:"timestamp" `
}
type Service interface {
	insertInMemory() error
	takeFromMemory() error
	takeFromDB() error
}

type ServiceImpl struct {
}

// service implementer of the Service interface. This structure implements all the logicto handle the
// received server request
func NewServiceImpl() Service {
	serviceImpl := ServiceImpl{}
	return &serviceImpl
}

// service to add a new trainer into the system
func (serviceImpl *ServiceImpl) insertInMemory() error {
	return nil
}
func (serviceImpl *ServiceImpl) takeFromMemory() error {
	return nil
}
func (serviceImpl *ServiceImpl) takeFromDB() error {
	return nil
}
