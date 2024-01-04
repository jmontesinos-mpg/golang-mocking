package test

type APIServiceA interface {
	GetAll() ([]string, error)
	Delete(ID string) error
	Add(ID string) error
}

type TestServiceB struct {
	serviceA APIServiceA
}

func (s *TestServiceB) GetAllFromServiceA() ([]string, error) {
	return s.serviceA.GetAll()
}

func (s *TestServiceB) DeleteFromServiceA(ID string) error {
	return s.serviceA.Delete(ID)
}
