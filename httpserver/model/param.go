package model

type LoginHttpReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// for show
type LoginHttpRsp struct {
	Code     int       `json:"code"`
	Msg      int       `json:"msg"`
	UserInfo *UserInfo `json:"user_info"`
}

type UserInfo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
}

type UpdateNicknameReq struct {
	Nickname string `json:"nickname"`
}

type UpdateNicknameRsp struct {
	Code int `json:"code"`
	Msg  int `json:"msg"`
}

type UpdateProfileReq struct {
	Profile string `json:"profile"`
}

type UpdateProfileRsp struct {
	Code int `json:"code"`
	Msg  int `json:"msg"`
}
