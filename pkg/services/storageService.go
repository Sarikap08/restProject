package services

import "github.com/Sarikap08/restProject/pkg/models"

type StorageService struct {
	data []models.Machine
}

func NewStorageService() *StorageService {
	return &StorageService{
		data: []models.Machine{},
	}
}

func (s *StorageService) AddMachine(m *models.Machine) {
	s.data = append(s.data, *m)
}

func (s *StorageService) GetAllMachine() *[]models.Machine {
	return &s.data
}
