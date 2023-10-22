package utils

import (
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, file *multipart.FileHeader) error {
	currentTime := time.Now().Format("01-02-06  15-04-.999")

	err := c.SaveUploadedFile(file, "assets/images"+currentTime+file.Filename)
	return err
}
