package db

import (
	"context"
	"testing"

	"github.com/jayantkatia/backend_upcoming_mobiles/util"
	"github.com/stretchr/testify/require"
)

func createRandomDevice(t *testing.T) Device {
	arg := InsertDeviceParams{
		DeviceName: util.RandomDeviceName(),
		Expected:   "Expected Launch: May, 20" + util.RandomYear(),
		Price:      util.RandomPrice(),
		ImgUrl:     "https://github.com",
		SourceUrl:  "https://github.com",
		SpecScore:  util.RandomSpecScore(),
	}
	device, err := testQueries.InsertDevice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, device)
	require.Equal(t, arg.DeviceName, device.DeviceName)
	require.Equal(t, arg.Price, device.Price)
	require.Equal(t, arg.SpecScore, device.SpecScore)
	require.Equal(t, arg.Expected, device.Expected)
	require.Equal(t, arg.ImgUrl, device.ImgUrl)
	require.Equal(t, arg.SourceUrl, device.SourceUrl)
	require.NotZero(t, device.Price)

	return device
}
func TestInsertDevice(t *testing.T) {
	createRandomDevice(t)
}

func TestDeleteDevices(t *testing.T) {
	err := testQueries.DeleteDevices(context.Background())
	require.NoError(t, err)
}

func TestGetDevices(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomDevice(t)
	}
	devices, err := testQueries.GetDevices(context.Background())
	require.NoError(t, err)

	for _, device := range devices {
		require.NotEmpty(t, device)
	}
}
