package proto

const (
	LoginRequest = iota + 1
	LoginResponse
	UpdateNicknameRequest
	UpdateNicknameResponse
	UpdateProfileRequest
	UpdateProfileResponse
)

type LoginRequestParam struct {
	Username string
	Password string
}

type LoginResponseParam struct {
	ErrCode int
	Msg     string
}

type UpdateNicknameRequestParam struct {
	Uid      string
	Nickname string
}

type UpdateNicknameResponseParam struct {
	ErrCode int
	Msg     string
}

type UpdateProfileRequestParam struct {
	Uid     string
	Profile string
}

type UpdateProfileResponseParam struct {
	ErrCode int
	Msg     string
}
