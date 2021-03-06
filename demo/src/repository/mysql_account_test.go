package repository_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	accountRepo "repository"
	"testing"
)

func TestGetByName(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "Stub Name 1", "Stun Email 2")
	sql := `SELECT id, name, email FROM account WHERE name=?`

	// Stub/Mocking
	mock.ExpectQuery(sql).WillReturnRows(rows)
	ar := accountRepo.NewMySQLAccountRepository(db)

	// Act
	account, err := ar.GetByName(context.TODO(), "Stub Name 1")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, account)
}

func TestGetByNameWithDataNotFound(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "email"})
	sql := `SELECT id, name, email FROM account WHERE name=?`

	// Stub/Mocking
	mock.ExpectQuery(sql).WillReturnRows(rows)
	ar := accountRepo.NewMySQLAccountRepository(db)

	// Act
	account, err := ar.GetByName(context.TODO(), "Stub Name 1")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestGetByNameWithFail(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sql := `SELECT id, name, email FROM account WHERE name=?`

	// Stub/Mocking
	mock.ExpectQuery(sql).WillReturnError(fmt.Errorf("some error"))
	ar := accountRepo.NewMySQLAccountRepository(db)

	// Act
	account, err := ar.GetByName(context.TODO(), "Stub Name 1")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestGetByNameWithFailAndNoColumeName(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "email"}).AddRow(1, "Stun Email 2")
	sql := `SELECT id, name, email FROM account WHERE name=?`

	// Stub/Mocking
	mock.ExpectQuery(sql).WillReturnRows(rows)
	ar := accountRepo.NewMySQLAccountRepository(db)

	// Act
	account, err := ar.GetByName(context.TODO(), "Stub Name 1")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, account)
}
