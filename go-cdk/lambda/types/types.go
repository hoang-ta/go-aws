package types

type RegisterUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}