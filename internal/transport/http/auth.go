package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JWTAuth(caller func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if !validateToken(authHeaderParts[1]) {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		caller(w, r)
	}
}

func validateToken(accessToken string) bool {
	var signingKey = []byte("missionimposible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}

		return signingKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
