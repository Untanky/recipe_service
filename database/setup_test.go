package database_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB gorm.DB
var Mock sqlmock.Sqlmock

func SetupDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()

	if err != nil {
		t.Error(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	DB = *gormDB
	Mock = mock

	return gormDB, mock
}

func CleanUpDB(t *testing.T) {
	err := Mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

	// sqlDB, err := DB.DB()
	// if err != nil {
	// 	t.Error(err)
	// }

	// err = sqlDB.Close()
	// if err != nil {
	// 	t.Error(err)
	// }
}
