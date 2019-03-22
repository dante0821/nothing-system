package globals

// CheckErr 检查错误，如果保存则panic
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// NewErrorRsp 新建http错误
func NewErrorRsp(message string, err error, code int) ResponseInfo {
	if err != nil {
		return ResponseInfo{ReturnMsg: message, Data: err.Error(), ReturnCode: code}
	} else {
		return ResponseInfo{ReturnMsg: message, ReturnCode: code}
	}
}
