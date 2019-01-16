package controllers

const (
	CODE_SUCCESS = 0
	MSG_SUCCESS  = "success"

	CODE_FAIL = 1
	MSG_FAIL  = "failed"

	CODE_ILLEGAL_PARAM = 2
	MSG_ILLEGAL_PARAM  = "非法参数"
	CODE_CODE_ERROR    = 3
	MSG_CODE_ERROR     = "验证码错误"

	CODE_USER_USERNAME_NOT_EXIST = 100
	MSG_USER_USERNAME_NOT_EXIST  = "当用户不存在"
	CODE_PASSWORD_ERROR          = 101
	MSG_PASSWORD_ERROR           = "用户名或密码错误，请重新输入"
	CODE_MORE_MAXUSER            = 102
	MSG_MORE_MAXUSER             = "当前账户的登录用户量已达上限，请稍后重试"
	CODE_LOGIN_FAILED            = 103
	MSG_LOGIN_FAILED             = "登陆失败"
	CODE_AUTHCODE_TIMEOUT        = 104
	MSG_AUTHOCODE_TIMEOUT        = "验证码已过期"


	CODE_WINCHUSER_PHONE_EXIST = 201
	MSG_WINCHUSER_PHONE_EXIST  = "手机号已被注册"
	CODE_USER_IS_NOR_STATUS    = 202
	MSG_USER_IS_NOR_STATUS     = "用户处于禁用状态"

	CODE_NOT_LOGIN                        = 300
	MSG_NOT_LOGIN                         = "not login in"
	CODE_WINCHUSERINFO_USERNAME_NOT_EXIST = 301
	MSG_WINCHUSERINFO_USERNAME_NOT_EXIST  = "当用户信息不存在"

	MSG_NOT_LINK_USER     = "没有关联的用户id"
	SEND_SMS_MESSAGE_FAIL = "短信发送失败"

	MSG_LOGOUT_FAIL        = "注销失败"
	MSG_UPDATE_FAIL        = "数据更新状态失败，请重试"
	MSG_POSTDATA_FAIL      = "业务方更新状态失败，请重试"
	MSG_ILLEGAL_BUSINESSID = "抱歉,没有找到相关内容"

	DEF_STATUS_DOWN = 0 //默认禁用状态

	DEF_STATUS_UP = 1 //默认启用状态

	REDIS_EXPIRATION = 30 * 60

	MSG_FREQUENCY_LIMIT = "frequency limit"
	DEAFAUD_INFOID      = "-1"
	ADMIN_BID           = "-1"

	MESSAGE_NAME           = "message"
	PAGE_BEAN_NAME         = "pageBean"
	USER_SESSION_NAME      = "winchannel_user"
	USER_CODE_NUMBER       = "code_number"
	DEFAULT_SUCCESS_INFO   = "操作成功！"
	DEFAULT_FAIL_INFO      = "操作失败！"
	ACTION_FAIL_USER_EXIST = "操作失败，用户已存在!"
	SEND_MESSAGE_FAIL      = "消息发送失败！"
	SEND_MESSAGE_SUCCESS   = "消息发送成功！"
	RECIVE_MESSAGE_FAIL    = "消息接收失败！"
	RECIVE_MESSAGE_SUCCESS = "消息接收成功！"

	SMS_PASSWORD = "0b334927e18add010d09"

	FILE_SIZE_UD            = 1024 * 5
	TIME_STAMP_DAY          = 1000 * 60 * 60 * 24
	CODE_NUMBER             = "code_number_"
	REDIS_SESSION_CODE_TIME = 60 * 5
	REDIS_SESSION_USER_TIME = 60 * 30

	MAX_FILE_COUNT = 1024 * 1024
)

const (
	DEFAULT_PAGE = 1 //默认页签

	DEFAULT_PAGE_SIZE = 10 //默认页容量

	DEFAULT_CODE_LENGTH = 6 //默认验证码长度

	DEFAULT_PHONE_LENGTH = 1 //默认电话号码长度
)

//时间相关类型定义
const (
	SimpleDateFormat = "yyyy-MM-dd HH:mm:ss"
	SDF_ymd          = "yyyy-MM-dd"
	SDF_Y            = "yyyy"
	SDF_M            = "MM"
	SDF_D            = "dd"
)

//登录类型
const (
	LOGIN_TYPE_GETAUTHCODE = 1
	LOGIN_TYPE_AUTHCODE    = 2
	LOGIN_TYPE_MOBILE      = 3
)

//sql查询错误
const (
	ERR_NO_ROWS = 1
)

type RetHeader struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Total       int    `json:"total"`
	PageSize    int    `json:"page_size"`
	CurrentPage int    `json:"current_page"`
}

func NewRetHeader(code int, msg string) *RetHeader {
	return &RetHeader{
		Code:        code,
		Message:     msg,
		Total:       0,
		PageSize:    1,
		CurrentPage: 1}
}
