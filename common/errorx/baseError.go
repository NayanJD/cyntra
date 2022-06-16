package errorx

import "google.golang.org/grpc/codes"

const defaultCode = 1001

type CodeResponse struct {
	Code codes.Code  `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type CodeResponseData struct {
	Code codes.Code  `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewCodeResponse(code codes.Code, msg string, data interface{}) error {
	return &CodeResponse{Code: code, Msg: msg, Data: data}
}

func NewDefaultError(msg string) error {
	return NewCodeResponse(defaultCode, msg, nil)
}

func NewDefaultSuccess(data interface{}) error {
	return NewCodeResponse(0, "success", data)
}

func (e *CodeResponse) Error() string {
	return e.Msg
}

func (e *CodeResponse) Response() *CodeResponseData {
	return &CodeResponseData{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
	}
}

func NewValidationCodeResponse(err error) error {
	return NewCodeResponse(codes.InvalidArgument, codes.InvalidArgument.String(), err.Error())
}
