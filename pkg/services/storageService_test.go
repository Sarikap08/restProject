package services_test

import (
	"testing"
	"time"

	"github.com/Sarikap08/restProject/pkg/models"
	"github.com/Sarikap08/restProject/pkg/services"
	"github.com/ddosify/go-faker/faker"
	"github.com/stretchr/testify/suite"
)

type StorageServiceTestSuite struct {
	suite.Suite
	storageService *services.StorageService
}

func (suite *StorageServiceTestSuite) SetupSuite() {
	suite.storageService = services.NewStorageService()
}

func (suite *StorageServiceTestSuite) TestGetMachine() {
	numberOfMachine := 10

	expectedMachine := getFakeMachine(numberOfMachine)

	for _, machine := range expectedMachine {
		suite.storageService.AddMachine(&machine)
	}

	actualMachines := suite.storageService.GetAllMachine()
	suite.Equal(len(expectedMachine), len(*actualMachines))

}

func getFakeMachine(number int) []models.Machine {
	machines := []models.Machine{}
	faker := faker.NewFaker()

	for i := 0; i < number; i++ {
		machine := models.Machine{
			Id:           faker.RandomUUID(),
			MachineId:    faker.RandomInt(),
			LastLoggedIn: faker.RandomString(),
			SysTime:      time.Time{},
			Stat: models.Stat{
				CpuTemp:  faker.RandomInt(),
				FanSpeed: faker.RandomInt(),
				HDDSpace: faker.RandomInt(),
			},
		}
		machines = append(machines, machine)
	}
	return machines
}

func TestStorageServiceTestSuite(t *testing.T) {
	suite.Run(t, new(StorageServiceTestSuite))
}
