package handlers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type FileAnchor struct {
	FileName string
	Href     string
}
type DirectoryAnchor struct {
	DirectoryName string
	Href          string
}

func linuxCleanPath(path string) string {
	cleanPath := filepath.Clean(path)
	if filepath.Separator == '/' {
		return cleanPath
	}

	return strings.ReplaceAll(cleanPath, string(filepath.Separator), "/")
}

func DirectoryHandler(c *gin.Context, absolutePath *string) {
	files, _ := ioutil.ReadDir(*absolutePath)
	directoryAnchorList := make([]DirectoryAnchor, 0)
	fileAnchorList := make([]FileAnchor, 0)
	// previous folder
	directoryAnchorList = append(directoryAnchorList, DirectoryAnchor{
		DirectoryName: "..",
		Href:          linuxCleanPath(c.Request.URL.Path + "/.."),
	})
	for _, file := range files {
		href := linuxCleanPath(c.Request.URL.Path + "/" + file.Name())
		if file.IsDir() {
			directoryAnchorList = append(directoryAnchorList, DirectoryAnchor{
				DirectoryName: file.Name() + string(filepath.Separator),
				Href:          href,
			})

		} else {
			fileAnchorList = append(fileAnchorList, FileAnchor{
				FileName: file.Name(),
				Href:     href,
			})
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"directoryAnchorList": directoryAnchorList,
		"fileAnchorList":      fileAnchorList,
		"currentDirectory":    c.Request.URL.Path,
	})

}
