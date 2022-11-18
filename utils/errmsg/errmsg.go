package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//1000...用户模块的错误
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USER_NOT_EXIT      = 1003
	ERROR_TOKEN_EXIT         = 1004
	ERROR_TOKEN_RUNTIME      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_FORMAT_RRONG = 1007
	ERROR_USER_NORIGHT       = 1008
	//2000...文章模块的错误
	ERROR_ARTICLE_NOTEXIT = 2001

	//3000...分类模块的错误
	ERROR_CATEGORYNAME_USED    = 3001
	ERROR_CATEGORYNAME_NOTEXIT = 3002
)

var codeMsg = map[int]string{
	SUCCESS:                    "OK",
	ERROR:                      "Fail",
	ERROR_USERNAME_USED:        "用户名已经存在",
	ERROR_PASSWORD_WRONG:       "密码错误",
	ERROR_USER_NOT_EXIT:        "用户不存在",
	ERROR_TOKEN_EXIT:           "token不存在",
	ERROR_TOKEN_RUNTIME:        "token已过期",
	ERROR_TOKEN_WRONG:          "token错误",
	ERROR_TOKEN_FORMAT_RRONG:   "token格式错误",
	ERROR_CATEGORYNAME_USED:    "分类已经存在",
	ERROR_ARTICLE_NOTEXIT:      "文章不存在",
	ERROR_CATEGORYNAME_NOTEXIT: "分类不存在",
	ERROR_USER_NORIGHT:         "用户没有权限",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
