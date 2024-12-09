package repositories

import (
	"go-react/internal/models"
	"gorm.io/gorm"
)

type DeviceRepositoryI interface {
	// GetDeviceById returns a device by its ID.
	GetDeviceById(id string) (*models.Device, error)
	// GetDeviceByIdentifier returns a device by its identifier.
	GetDeviceByIdentifier(identifier string) (*models.Device, error)
	// AddDevice adds a new device to the database.
	AddDevice(device *models.Device) error
	// DeleteDevice deletes a device from the database.
	DeleteDevice(id string) error
}

type DeviceRepository struct {
	// db is the database connection.
	db *gorm.DB
}

// GetDeviceById returns a device by its ID.
func (r *DeviceRepository) GetDeviceById(id string) (*models.Device, error) {
	var device models.Device
	if err := r.db.Where("id = ?", id).First(&device).Error; err != nil {
		return nil, err
	}
	return &device, nil
}

// GetDeviceByIdentifier returns a device by its identifier.
func (r *DeviceRepository) GetDeviceByIdentifier(identifier string) (*models.Device, error) {
	var device models.Device
	if err := r.db.Where("device_identifier = ?", identifier).First(&device).Error; err != nil {
		return nil, err
	}
	return &device, nil
}

// AddDevice adds a new device to the database.
func (r *DeviceRepository) AddDevice(device *models.Device) error {
	return r.db.Create(device).Error
}

// DeleteDevice deletes a device from the database.
func (r *DeviceRepository) DeleteDevice(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Device{}).Error
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{
		db: db,
	}
}
