package proto

type LoginRequestParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseParam struct {
	ErrCode  int    `json:"err_code"`
	Msg      string `json:"msg"`
	UserInfo User   `json:"user_info"`
	Cookie   string `json:"cookie"`
}

type User struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
}

type UpdateNicknameRequestParam struct {
	Uid      string `json:"uid"`
	Cookie   string `json:"cookie"`
	Nickname string `json:"nickname"`
}

type UpdateNicknameResponseParam struct {
	ErrCode int    `json:"err_code"`
	Msg     string `json:"msg"`
}

type UpdateProfileRequestParam struct {
	Uid     string `json:"uid"`
	Cookie  string `json:"cookie"`
	Profile string `json:"profile"`
}

type UpdateProfileResponseParam struct {
	ErrCode int    `json:"err_code"`
	Msg     string `json:"msg"`
}
