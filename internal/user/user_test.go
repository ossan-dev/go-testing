//go:build integration

package user_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/ossan-dev/gotesting/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// BEFORE: docker run -d -p 54322:5432 -e POSTGRES_PASSWORD=postgres postgres
type userSuite struct {
	suite.Suite
	gormClient *gorm.DB
	container  testcontainers.Container
}

func (s *userSuite) SetupSuite() {
	// Arrange
	var err error
	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{"5432/tcp"},
		Env:          map[string]string{"POSTGRES_PASSWORD": "postgres"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}
	s.container, err = testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	require.NoError(s.T(), err)
	endpoint, err := s.container.Endpoint(context.Background(), "")
	require.NoError(s.T(), err)
	s.gormClient, err = gorm.Open(postgres.Open(fmt.Sprintf("host=localhost port=%v user=postgres password=postgres dbname=postgres sslmode=disable", strings.Split(endpoint, ":")[1])))
	require.NoError(s.T(), err)
	err = s.gormClient.AutoMigrate(&user.User{})
	require.NoError(s.T(), err)
}

func (s *userSuite) TearDownSuite() {
	if s.container != nil {
		require.NoError(s.T(), s.container.Terminate(context.Background()))
	}
}

func (s *userSuite) SetupTest() {
	require.NoError(s.T(), s.gormClient.Delete(&user.User{}, "1 = 1").Error)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userSuite))
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
