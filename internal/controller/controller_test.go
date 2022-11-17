package controller_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"balance/internal/controller"
	"balance/internal/mocks"
)

const (
	testUserID uint = 1
)

var (
	nilError error
)

func TestGetBalance(t *testing.T) {
	mockDB := &mocks.Db{}

	expectedUser := controller.User{
		ID:      testUserID,
		Balance: 22,
		Reserve: 33,
	}

	mockDB.On("ReadUser",
		context.Background(),
		testUserID,
	).Return(
		expectedUser,
		nilError,
	)

	ctrl := controller.New(mockDB)

	actualUser, err := ctrl.GetBalance(context.Background(), controller.RequestBalance{
		UserID: 1,
	})

	require.NoError(t, err)
	require.Equal(t, expectedUser, actualUser)

}
