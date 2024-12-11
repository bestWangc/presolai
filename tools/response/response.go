package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`  // HTTP 状态码
	Message string `json:"message"` // 提示信息
	Data    any    `json:"data"`    // 返回的数据
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

// Error 统一错误响应
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	})
}
