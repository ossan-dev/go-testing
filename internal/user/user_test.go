//go:build integration

package user_test

import (
	"testing"

	"github.com/ossan-dev/gotesting/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// BEFORE: docker run -d -p 54322:5432 -e POSTGRES_PASSWORD=postgres postgres
type userSuite struct {
	suite.Suite
	gormClient *gorm.DB
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userSuite))
}

func (s *userSuite) SetupSuite() {
	// Arrange
	var err error
	s.gormClient, err = gorm.Open(postgres.Open("host=localhost port=54322 user=postgres password=postgres dbname=postgres sslmode=disable"))
	require.NoError(s.T(), err)
	err = s.gormClient.AutoMigrate(&user.User{})
	require.NoError(s.T(), err)
}

func (s *userSuite) SetupTest() {
	require.NoError(s.T(), s.gormClient.Delete(&user.User{}, "1 = 1").Error)
}

func (s *userSuite) TestAddUser() {
	// Act
	err := user.AddUser(s.gormClient, &user.User{Name: "john doe"})
	// Assert
	require.NoError(s.T(), err)
	var userInDB user.User
	err = s.gormClient.Where("Name = ?", "john doe").First(&userInDB).Error
	assert.NoError(s.T(), err)
}

// go test -tags="integration" -v
// AFTER: docker stop <abc> && docker rm -f <abc>
