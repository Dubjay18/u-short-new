package models

import "time"

type Device struct {
	ID                  string    `json:"id" gorm:"primaryKey"`
	DeviceIdentifier    string    `json:"device_identifier" gorm:"unique"`
	UsageCount          int       `json:"usage_count" gorm:"not_null"`
	FirstUsageTimestamp time.Time `json:"first_usage_timestamp" gorm:"not_null"`
	IsBlocked           bool      `json:"is_blocked" gorm:"not_null"`
}

func (d *Device) getDevice() string {
	return d.DeviceIdentifier
}
