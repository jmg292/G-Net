package models

import (
	"crypto/x509"
	roles "gnet/access_control/user_roles"

	"github.com/google/uuid"
)

const UserTableCreationString string = `CREATE TABLE IF NOT EXISTS Users(
	Id VARCHAR(36) PRIMARY KEY,
	Roles INTEGER,
	Name TEXT,
	Certificate BLOB
);`

const (
	UserInsertionString = "INSERT INTO Users (Id, Roles, Name, Certificate) VALUES (?,?,?,?)"
	UserQueryString     = "SELECT * FROM Users WHERE Id=? LIMIT 1"
)

type User struct {
	Id          uuid.UUID
	Name        string
	Certificate *x509.Certificate
	Roles       roles.UserRole
}

func NewEmptyUser() *User {
	return &User{
		Id: uuid.Nil,
	}
}

func (user *User) IsEmptyUser() bool {
	return user.Id == uuid.Nil
}
