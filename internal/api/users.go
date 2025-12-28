package api

type CreateUserInput struct {
	Username string `json:"username"`;
	Password string `json:"password"`;
}

//Used in authentication JWT token res
//TODO depreciated delate
type LoginResponse struct {
	Token string `json:"token"`
}