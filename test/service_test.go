package test

import (
	"jmontesinos/golang-mocking/test/mocks"
	"testing"

	"github.com/ovechkin-dm/mockio/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAllFromServiceA(t *testing.T) {
	s := TestServiceB{
		serviceA: &MockServiceA{},
	}

	strs, err := s.GetAllFromServiceA()

	require.NoError(t, err)
	require.Equal(t, []string{"hello", "world"}, strs)
}

type MockServiceA struct {
}

func (m *MockServiceA) GetAll() ([]string, error) {
	return []string{"hello", "world"}, nil
}

func (m *MockServiceA) Delete(ID string) error {
	return nil
}

func (m *MockServiceA) Add(ID string) error {
	return nil
}

func TestGetAllFromServiceAUsingUberGoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServiceA := mocks.NewMockAPIServiceA(ctrl)

	mockServiceA.EXPECT().GetAll().Return([]string{"hello", "world"}, nil)

	s := TestServiceB{
		serviceA: mockServiceA,
	}

	strs, err := s.GetAllFromServiceA()

	require.NoError(t, err)
	require.Equal(t, []string{"hello", "world"}, strs)
	ctrl.Finish()
}

func TestGetAllFromServiceAUsingMockio(t *testing.T) {
	mock.SetUp(t)
	mockServiceA := mock.Mock[APIServiceA]()

	mock.When(mockServiceA.GetAll()).ThenReturn([]string{"hello", "world"}, nil)

	s := TestServiceB{
		serviceA: mockServiceA,
	}

	strs, err := s.GetAllFromServiceA()

	require.NoError(t, err)
	require.Equal(t, []string{"hello", "world"}, strs)

	mock.Verify(mockServiceA, mock.Once()).GetAll()
}

func TestGetAllFromServiceAUsingTestifyAndMockery(t *testing.T) {
	mockServiceA := mocks.NewMockeryAPIServiceA(t)

	mockServiceA.EXPECT().GetAll().Return([]string{"hello", "world"}, nil)

	s := TestServiceB{
		serviceA: mockServiceA,
	}

	strs, err := s.GetAllFromServiceA()

	require.NoError(t, err)
	require.Equal(t, []string{"hello", "world"}, strs)

	mockServiceA.AssertExpectations(t)
}
