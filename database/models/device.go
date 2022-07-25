package models

import (
	"crypto/x509"

	"github.com/google/uuid"
)

const DeviceTableCreationString string = `CREATE TABLE IF NOT EXISTS Devices(
	Id VARCHAR(36) PRIMARY KEY,
	Name VARCHAR(255),
	Domain VARCHAR(255),
	Certificate BLOB
);`

const (
	DeviceInsertionString = "INSERT INTO Devices (Id, Name, Domain, Certificate) VALUES (?,?,?,?)"
	DeviceQueryString     = "SELECT * FROM Devices WHERE Id=? LIMIT 1"
)

type Device struct {
	Id          uuid.UUID
	Name        string
	Domain      string
	Certificate *x509.Certificate
}

func NewEmptyDevice() *Device {
	return &Device{
		Id: uuid.Nil,
	}
}

func (device *Device) IsNullDevice() bool {
	return device.Id == uuid.Nil
}
