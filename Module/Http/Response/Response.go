package Response

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

func NewResponse(status int, data interface{}, message string) *Response {
	return &Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
}
