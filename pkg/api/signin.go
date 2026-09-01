package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func jwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

type signinRequest struct {
	Password string `json:"password"`
}

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var req signinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJsonError(w, err)
		return
	}

	pass := os.Getenv("TODO_PASSWORD")
	if req.Password != pass {
		writeJsonError(w, errors.New("неверный пароль"))
		return
	}

	claims := jwt.MapClaims{
		"password": hashPassword(pass),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret())
	if err != nil {
		writeJsonError(w, err)
		return
	}

	writeJson(w, map[string]any{"token": tokenString})
}
