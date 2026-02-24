package user

type RegisterRequest struct {
	Email string

	Password string

	Username string
}

type LoginRequest struct {
	Email string

	Password string
}

type WhoamiRequest struct {
	AccessToken string
}
