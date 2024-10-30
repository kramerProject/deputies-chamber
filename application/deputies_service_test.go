package application_test

import (
	"errors"
	"testing"

	"github.com/kramerProject/deputies-chamber/application"
	mock_application "github.com/kramerProject/deputies-chamber/application/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestProductService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	deputies := application.Deputies{}
	deputiesClient := mock_application.NewMockDeputiesClientInterface(ctrl)
	deputiesClient.EXPECT().GetAll().Return(deputies, errors.New("deu ruim")).AnyTimes()

	service := application.DeputiesService{
		DeputiesClient: deputiesClient,
	}

	_, err := service.GetAll()
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "error returning deputies")
}
