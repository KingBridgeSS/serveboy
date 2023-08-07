package handlers

import (
	"github.com/gin-gonic/gin"
)

func DownloadHandler(c *gin.Context, absolutePath *string) {
	c.File(*absolutePath)
}
