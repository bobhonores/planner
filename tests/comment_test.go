//go:build e2e

package tests

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("missionimposible"))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return tokenString
}

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug":"/","author":"Elliot","body":"hello"}`).
			SetHeader("Authorization", "Bearer "+createToken()).
			Post("http://localhost:8080/api/v1/comments")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("can't post comment without JWT", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug":"/","author":"Elliot","body":"hello"}`).
			Post("http://localhost:8080/api/v1/comments")
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})
}
