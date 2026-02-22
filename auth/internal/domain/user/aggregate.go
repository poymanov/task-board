package user

type NewUser struct {
	Email string

	Password string

	Username string
}

func NewNewUser(email, password, username string) NewUser {
	return NewUser{
		Email:    email,
		Password: password,
		Username: username,
	}
}

type AuthClaims struct {
	UserId int

	Email string

	Username string
}
