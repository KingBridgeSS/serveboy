package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"serveboy/config"
	"serveboy/handlers"
	"strings"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	// api
	sbapi := r.Group("/sbapi")
	sbapi.POST("upload", handlers.UploadHandler)
	// directory / file
	r.NoRoute(func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "..") {
			c.String(http.StatusBadRequest, "invalid path")
			return
		}
		absolutePath := filepath.Join(config.ServeDirectory, c.Request.URL.Path)
		fileInfo, err := os.Stat(absolutePath)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if fileInfo.IsDir() {
			handlers.DirectoryHandler(c, &absolutePath)
		} else {
			handlers.DownloadHandler(c, &absolutePath)
		}
	})

	return r
}
