package database

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func TestGetUser(t *testing.T) {
	// 1. Create mock sql.DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// 2. Open GORM with the mock db using a dialector
	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	// 3. Define expectations
	// Note: Use regexp.QuoteMeta to handle special SQL characters
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "John Doe")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	// 4. Execute the GORM method
	var user User
	err = gormDB.First(&user, 1).Error

	// 5. Assert results and expectations
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}
