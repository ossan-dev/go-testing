package end2end_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestAccounts(t *testing.T) {
	// Setup
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles("../docker-compose.yml"))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal))
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	err = compose.
		WaitForService("api_service", wait.ForListeningPort("8080/tcp")).
		Up(ctx, tc.Wait(true))
	require.NoError(t, err)

	t.Run("Valid HTTP Request return data", func(t *testing.T) {
		client := &http.Client{}
		r, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/accounts?iban=IT77D0300203280543851368733", nil)
		require.NoError(t, err)
		res, err := client.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Not valid HTTP request return err", func(t *testing.T) {
		client := &http.Client{}
		r, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/accounts?iban=abc", nil)
		require.NoError(t, err)
		res, err := client.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
