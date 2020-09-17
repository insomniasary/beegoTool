/**
 * 自定义错误类型及代码
 * Author: yansheng
 * RegTime: 2019/5/18
 */
package services

type CskError interface {
	Error() (string, string)
}
type CskErrors struct {
	Code string
	Msg  string
}

var ErrCode = map[string]string{

	"50000": "登录超时",
}

func NewError(errcode string) CskError {
	if errcode == "2000" {
		return nil
	}
	errMsg := ErrCode[errcode]
	return &CskErrors{
		//retError:retError
		Code: errcode,
		Msg:  errMsg,
	}
}

func NewCommonError(err error) CskError {
	errMsg := err.Error()
	return &CskErrors{
		Code: "50000",
		Msg:  errMsg,
	}
}

func NewMsgError(errMsg string) CskError {
	return &CskErrors{
		Code: "40000",
		Msg:  errMsg,
	}
}

func (this *CskErrors) Error() (string, string) {
	return this.Code, this.Msg
}
