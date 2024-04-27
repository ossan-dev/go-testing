//go:build integration

package user_test

import (
	"testing"

	"github.com/ossan-dev/gotesting/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// BEFORE: docker run -d -p 54322:5432 -e POSTGRES_PASSWORD=postgres postgres
func TestAddUser(t *testing.T) {
	// Arrange
	db, err := gorm.Open(postgres.Open("host=localhost port=54322 user=postgres password=postgres dbname=postgres sslmode=disable"))
	require.NoError(t, err)
	err = db.AutoMigrate(&user.User{})
	require.NoError(t, err)
	// Act
	err = user.AddUser(db, &user.User{Name: "john doe"})
	// Assert
	require.NoError(t, err)
	var userInDB user.User
	err = db.Where("Name = ?", "john doe").First(&userInDB).Error
	assert.NoError(t, err)
}

// go test -tags="integration" -v
// AFTER: docker stop <abc> && docker rm -f <abc>
