package types

import "golang.org/x/crypto/bcrypt"

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