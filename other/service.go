package other

import "jmontesinos/golang-mocking/test"

type OtherService struct {
	serviceA test.APIServiceA
}

func (s *OtherService) AddAllFromServiceA(IDs []string) error {
	var err error

	for _, ID := range IDs {
		err = s.serviceA.Add(ID)
		if err != nil {
			return err
		}
	}

	return nil
}
