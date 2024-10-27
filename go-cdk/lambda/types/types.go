package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	UserName string `json:"username"`
	PasswordHash string `json:"password"`
}

func NewUser(user RegisterUser) (*User, error) {
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if error != nil {
		return nil, error
	}
	return &User{
		UserName: user.UserName,
		PasswordHash: string(hashedPassword),
	}, nil
}

func ValidatePassword(hashedPassword, plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
	return err == nil
}

func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"user": user.UserName,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := "secret"

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return ""
	}
	return tokenString
}