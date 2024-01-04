package other

import (
	"jmontesinos/golang-mocking/test"
	"jmontesinos/golang-mocking/test/mocks"
	"testing"

	mockio "github.com/ovechkin-dm/mockio/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAllFromServiceAUsingUberGoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServiceA := mocks.NewMockAPIServiceA(ctrl)

	mockServiceA.EXPECT().Add("TEST").Times(1)
	mockServiceA.EXPECT().Add("OTHER")

	s := OtherService{
		serviceA: mockServiceA,
	}

	err := s.AddAllFromServiceA([]string{"TEST", "OTHER"})

	require.NoError(t, err)
	ctrl.Finish()
}

func TestGetAllFromServiceAUsingMockio(t *testing.T) {
	mockio.SetUp(t)
	mockServiceA := mockio.Mock[test.APIServiceA]()

	mockio.When(mockServiceA.Add(mockio.AnyString()))

	mockServiceX := mockio.Mock[test.APIServiceX]()
	mockServiceX.IncrementCounterBy(mockio.AnyInt())

	s := OtherService{
		serviceA: mockServiceA,
	}

	err := s.AddAllFromServiceA([]string{"TEST", "OTHER"})

	require.NoError(t, err)

	mockio.Verify(mockServiceA, mockio.Once()).Add("TEST")
	mockio.Verify(mockServiceA, mockio.Once()).Add("OTHER")
	// mock.VerifyNoMoreInteractions(mockServiceA) -- It does not work as expected
}

func TestGetAllFromServiceAUsingTestifyAndMockery(t *testing.T) {
	mockServiceA := mocks.NewMockeryAPIServiceA(t)

	mockServiceA.EXPECT().Add(mock.AnythingOfType("string")).Return(nil).Once()
	mockServiceA.EXPECT().Add("OTHER").Return(nil).Once()

	s := OtherService{
		serviceA: mockServiceA,
	}

	err := s.AddAllFromServiceA([]string{"TEST", "OTHER"})

	require.NoError(t, err)

	mockServiceA.AssertExpectations(t)
}
