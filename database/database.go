package database

import (
	"github.com/jmg292/G-Net/database/models"

	"github.com/google/uuid"
)

type GNetDatabase interface {
	Open(string) error
	Close() error
	GetConfiguration() (models.Configuration, error)
	CreateOrUpdateConfiguration(models.Configuration) error
	GetUserById(uuid.UUID) (models.User, error)
	CreateOrUpdateUser(models.User) error
	GetDeviceById(uuid.UUID) (models.Device, error)
	CreateOrUpdateDevice(models.Device) error
	GetSessionById(uuid.UUID) (models.Session, error)
	CreateOrUpdateSession(models.Session) error
}
