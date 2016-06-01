package models
import "github.com/jinzhu/gorm"

type Session struct  {
	gorm.Model
	UserID uint64
	SessionToken string
	IPAddress string
	DeviceType string
}
