package response

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

type iJsonContext interface {
    JSON(code int, i interface{}) error
}

func SendResponseJson(c iJsonContext, status string, message string, data interface{}) error {
    response := Response{
        Status:  status,
        Message: message,
        Data:    data,
    }

    return c.JSON(http.StatusOK, response)
}

func GetResponseJson(status string, message string, data interface{}) string {
    response := Response{
        Status:  status,
        Message: message,
        Data:    data,
    }

    result, _ := json.Marshal(response)

    return string(result)
}
