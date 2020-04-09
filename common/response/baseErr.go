package response

/****************自定义信息：code 状态码, msg 自定义错误信息***********************/

type BaseError struct {
	code     int
	messages string
}

func NewBaseError(code int, messages string) *BaseError {
	return &BaseError{code: code, messages: messages}
}

func (e *BaseError) Error() string {
	return e.messages
}

func (e *BaseError) Code() int {
	return e.code
}
