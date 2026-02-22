package user

type NewUser struct {
	Login string

	Password string

	Username string

	Email string
}

func NewNewUser(login, password, username, email string) NewUser {
	return NewUser{
		Login:    login,
		Password: password,
		Username: username,
		Email:    email,
	}
}
