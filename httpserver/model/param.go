package model

type LoginHttpReq struct {
	Username string
	Password string
}

// for show
type LoginHttpRsp struct {
	Username string
	Nickname string
	Profile  string
}

type UpdateNicknameReq struct {
	Nickname string
}

type UpdateProfileReq struct {
	Profile string
}
