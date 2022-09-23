package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")
	NotFound      = NewError(10000002, "找不到")
)

// 业务层错误码
var (
	ErrorLoginFail      = NewError(20010001, "用户登录失败")
	ErrorGetUserFail    = NewError(20010002, "获取用户失败")
	ErrorUpdateUserFail = NewError(20010003, "更新用户失败")
)
