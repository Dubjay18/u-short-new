package sevices

import (
	"errors"
	"go-react/internal/models"
	"go-react/internal/repositories"
)

type DeviceServiceI interface {
	// GetDeviceById returns a device by its ID.
	GetDeviceById(id string) (*models.Device, error)
	// GetDeviceByIdentifier returns a device by its identifier.
	GetDeviceByIdentifier(identifier string) (*models.Device, error)
	// AddDevice adds a new device to the database.
	AddDevice(device *models.Device) error
	// DeleteDevice deletes a device from the database.
	DeleteDevice(id string) error
}

type DeviceService struct {
	// deviceRepository is the repository that the service uses.
	deviceRepository repositories.DeviceRepositoryI
}

// GetDeviceById returns a device by its ID.
func (s *DeviceService) GetDeviceById(id string) (*models.Device, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	resp, err := s.deviceRepository.GetDeviceById(id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDeviceByIdentifier returns a device by its identifier
func (s *DeviceService) GetDeviceByIdentifier(identifier string) (*models.Device,error) {
	return nil,nil
}


// AddDevice adds a new device to the database
func (s *DeviceService)  AddDevice(device *models.Device) error {
	return nil
}