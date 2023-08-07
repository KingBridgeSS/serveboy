package main

import (
	"embed"
	"flag"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"serveboy/config"
	"serveboy/routes"
)

//go:embed templates/*
var templates embed.FS

var (
	port           = flag.String("p", "8088", "port")
	serveDirectory = flag.String("d", ".", "directory to serve")
	enableUpload   = flag.Bool("u", false, "enable upload")
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()
	config.Port = *port
	config.ServeDirectory = *serveDirectory
	config.EnableUpload = *enableUpload

	r := routes.InitRoutes()
	r.SetHTMLTemplate(template.Must(template.ParseFS(templates, "templates/*")))
	log.Print("Serving HTTP on port: " + config.Port)
	r.Run(":" + config.Port)

}
