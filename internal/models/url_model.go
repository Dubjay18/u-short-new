package models

import "time"

type Url struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	OriginalUrl string    `json:"original_url" gorm:"not_null"`
	ShortUrl    string    `json:"short_url" gorm:"not_null"`
	DeviceId    string    `json:"device_id" gorm:"not_null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not_null"`
	ExpiresAt   time.Time `json:"expires_at" gorm:"not_null"`
	ClickCount  int       `json:"click_count" gorm:"not_null"`
	IsActive    bool      `json:"is_active" gorm:"not_null"`
}

func (u *Url) getShortUrl() string {
	return u.ShortUrl
}
