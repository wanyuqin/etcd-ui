package xgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	SuccessCode = "0000"
	FailedCode  = "0001"
)

var (
	SuccessMsg = "success"
	FailedMsg  = "failed"
)

type Result struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewResult() Result {
	return Result{
		Code:    SuccessCode,
		Message: SuccessMsg,
	}
}
func NewFailedResult() Result {
	return Result{
		Code:    FailedCode,
		Message: FailedMsg,
	}
}

func Success(c *gin.Context) {
	result := NewResult()
	c.JSON(http.StatusOK, result)
	c.Abort()
}

func Response(c *gin.Context, data interface{}, err error) {
	result := NewResult()
	if err != nil {
		result.Code = FailedCode
		result.Message = err.Error()
		c.JSON(http.StatusOK, result)
		c.Abort()
		return
	}

	result.Data = data
	c.JSON(http.StatusOK, result)
	c.Abort()
}

func ResponsePage(c *gin.Context, data interface{}, count int64, err error) {
	result := NewResult()
	if err != nil {
		result.Code = FailedCode
		result.Message = err.Error()
		c.JSON(http.StatusOK, result)
		c.Abort()
		return
	}
	m := make(map[string]interface{})
	m["records"] = data
	m["total"] = count
	result.Data = m
	c.JSON(http.StatusOK, result)
	c.Abort()
}

func Failed(c *gin.Context, err error) {
	result := NewFailedResult()
	result.Message = err.Error()
	c.JSON(http.StatusOK, result)
	c.Abort()
}

func ParamInt(c *gin.Context, name string) (int, error) {
	p := c.Param(name)
	return strconv.Atoi(p)
}

func ParamInt64(c *gin.Context, name string) (int64, error) {
	p := c.Param(name)
	return strconv.ParseInt(p, 10, 64)
}

func PostFormInt(c *gin.Context, name string) (int, error) {
	p := c.PostForm(name)
	return strconv.Atoi(p)
}
