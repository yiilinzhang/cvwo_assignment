package api

type CreateUserInput struct {
	Username string `json:"username"`;
	Password string `json:"password"`;
}