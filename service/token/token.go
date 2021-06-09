package token

type Token interface {
	GenerateToken(Login) error
	RetrieveToken() string
}

type Login struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}
