package server

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
