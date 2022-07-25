package database

import (
	"crypto/x509"
	"database/sql"
	"fmt"
	userroles "gnet/access_control/user_roles"
	"gnet/database/models"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDatabase struct {
	handle *sql.DB
}

func (database *SqliteDatabase) ensureDatabaseIsOpen() error {
	if database.handle == nil {
		return fmt.Errorf("must call Open() before interacting with the database")
	}
	return nil
}

func (database *SqliteDatabase) createTables() error {
	tableCreationStrings := [3]string{
		models.UserTableCreationString,
		models.DeviceTableCreationString,
		models.SessionTableCreationString,
	}
	for _, tableCreationString := range tableCreationStrings {
		_, err := database.handle.Exec(tableCreationString)
		if err != nil {
			return err
		}
	}
	return nil
}

func (database *SqliteDatabase) Open(dbPath string) error {
	_, err := os.Stat(dbPath)
	shouldPrepareDatabase := os.IsNotExist(err)
	database.handle, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	if shouldPrepareDatabase {
		err = database.createTables()
		if err != nil {
			return err
		}
	}
	return nil
}

func (database *SqliteDatabase) Close() error {

}

func (database *SqliteDatabase) GetConfiguration() models.Configuration {

}

func (database *SqliteDatabase) CreateOrUpdateConfiguration(configuration models.Configuration) error {

}

func (database *SqliteDatabase) GetUserById(userId string) (*models.User, error) {
	if err := database.ensureDatabaseIsOpen(); err != nil {
		return models.NewEmptyUser(), err
	}
	statement, err := database.handle.Prepare(models.UserQueryString)
	if err != nil {
		return models.NewEmptyUser(), err
	}
	rows, err := statement.Query(userId)
	var userIdString, userName string
	var userRoles uint16
	var userCertificateBytes []byte
	for rows.Next() {
		err = rows.Scan(&userIdString, &userName, &userRoles, &userCertificateBytes)
		if err != nil {
			return models.NewEmptyUser(), err
		}
	}
	var parsedUserId uuid.UUID
	parsedUserId, err = uuid.Parse(userIdString)
	if err != nil {
		return models.NewEmptyUser(), err
	}
	var userCertificate *x509.Certificate
	userCertificate, err = x509.ParseCertificate(userCertificateBytes)
	if err != nil {
		return models.NewEmptyUser(), err
	}
	return &models.User{
		Id:          parsedUserId,
		Name:        userName,
		Roles:       userroles.UserRole(userRoles),
		Certificate: userCertificate,
	}, nil
}

func (database *SqliteDatabase) CreateOrUpdateUser(user *models.User) error {
	if err := database.ensureDatabaseIsOpen(); err != nil {
		return nil
	}
}

func (database *SqliteDatabase) GetDeviceById(deviceId string) *models.Device {

}

func (database *SqliteDatabase) CreateOrUpdateDevice(deviceId *models.Device) error {

}

func (database *SqliteDatabase) GetSessionById(sessionId string) *models.Session {

}

func (database *SqliteDatabase) CreateOrUpdateSession(session *models.Session) error {

}
