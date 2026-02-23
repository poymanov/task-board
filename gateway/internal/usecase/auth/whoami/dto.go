package whoami

type WhoamiDTO struct {
	AccessToken string
}

type WhoamiUserDTO struct {
	UserId int

	Email string

	Username string
}
