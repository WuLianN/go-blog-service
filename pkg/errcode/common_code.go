var (
	Success = NewError(200, "成功")
	ServerError = NewError(500, "服务器内部错误")
	TooManyRequests = NewError(40001, "请求频繁")
)