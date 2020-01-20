package httpProxy

type BaseAuthResult struct {
	IsPassed bool
	Msg      string      // 认证返回信息
	Info     interface{} // 身份信息，未知格式，确定后可以明确定义
}
