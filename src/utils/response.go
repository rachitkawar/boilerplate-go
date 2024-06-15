package utils

type ApiResponse struct {
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
	Error   interface{}  `json:"error"`
}

type ResponseCode string

var (
	ApiResponseCode = struct {
		SUCCESS ResponseCode
		ERROR   ResponseCode
	}{
		SUCCESS: "99",
		ERROR:   "00",
	}
)

func NewApiResponse(code ResponseCode, message string, data interface{}, error interface{}) *ApiResponse {
	return &ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   error,
	}
}
