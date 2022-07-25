package models

import "github.com/google/uuid"

const SessionTableCreationString string = `CREATE TABLE IF NOT EXISTS Sessions(
	Id VARCHAR(36) PRIMARY KEY,
	User VARCHAR(36),
	Device VARCHAR(36),
	StartTime UNSIGNED BIGINT,
	EndTime UNSIGNED BIGINT
);`

const (
	SessionInsertionString = "INSERT INTO Sessions (Id, User, Device, StartTime, EndTime) VALUES (?,?,?,?,?)"
	SessionQueryString     = "SELECT * FROM Sessions WHERE Id=? LIMIT 1"
)

type Session struct {
	Id        uuid.UUID
	Device    uuid.UUID
	User      uuid.UUID
	StartTime uint64
	EndTime   uint64
}
