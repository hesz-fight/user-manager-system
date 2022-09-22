package service

type LoginRequest struct {
	Username string
	Password string
}

type UpdateNicknameRequest struct {
	Nickname string
}

type UpdateProfileRequest struct {
	Profile string
}
