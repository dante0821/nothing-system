package globals

type ResponseInfo struct {
	Data       interface{} `json:"data"`
	ReturnCode int         `json:"code"`
	ReturnMsg  string      `json:"msg"`
}
