package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"serveboy/config"
	"strings"
)

func UploadHandler(c *gin.Context) {
	if !config.EnableUpload {
		c.String(http.StatusForbidden, "Upload is not enabled")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Prevent path traversal
	filename := filepath.Base(file.Filename)
	if strings.Contains(c.Query("path"), "..") {
		c.String(http.StatusBadRequest, "invalid path")
		return
	}

	savePath := filepath.Join(config.ServeDirectory, c.Query("path"), filename)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "File uploaded successfully")
}
