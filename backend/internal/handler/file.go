package handler

import (
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {

}

type CertificateFile struct {
	Kind int `json:"kind" form:"kind"`
}

func UploadCertificate(c *gin.Context) {

}
