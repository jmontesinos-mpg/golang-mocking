package mocks

//go:generate mockgen -destination gomock_service.go -package mocks jmontesinos/golang-mocking/test APIServiceA
//go:generate mockgen -destination gomock_other_service.go -package mocks jmontesinos/golang-mocking/test APIServiceX
