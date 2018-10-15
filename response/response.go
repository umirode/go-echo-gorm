package response

type Response struct {
	Data interface{} `json:"data"`
}

type iJsonContext interface {
	JSON(code int, i interface{}) error
}

func SendResponseJson(c iJsonContext, status int, data interface{}) error {
	response := Response{
		Data: data,
	}

	return c.JSON(status, response)
}
