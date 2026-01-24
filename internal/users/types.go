package users

type CreateUserInput struct {
	Username string `json:"username"`;
	Password string `json:"password"`;
}

type UserAuth struct {
	UserID string;
	PasswordHash string;
}