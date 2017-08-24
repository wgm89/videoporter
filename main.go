package main

import (
	"videoporter/handles"
	"videoporter/middleware"
	//"videoporter/models"
	"videoporter/env"
	"videoporter/util"
	"videoporter/util/gin_html_render"
	"videoporter/util/htmlrender"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func main() {
	env.LoadEnv()
	util.UseJSONLogFormat()

	//dao.InitDb()

	r := gin.New()
	if env.IsProduction {
		r.HTMLRender = htmlrender.New()
	} else {
		htmlRender := GinHTMLRender.New()
		htmlRender.Debug = gin.IsDebugging()
		htmlRender.Layout = "layout/layout.html"
		htmlRender.TemplatesDir = env.TemplateDir
		htmlRender.Ext = ".html"
		r.HTMLRender = htmlRender.Create()
	}

	//r.RedirectTrailingSlash = true

	r.Use(middleware.JSONLogMiddleware())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID(middleware.RequestIDOptions{AllowSetting: false}))
	r.Use(middleware.Auth())
	r.Use(middleware.CORS(middleware.CORSOptions{}))

	handles.InitHandle(r)

	port := env.Port
	log.Info("Service starting on port " + port)

	r.Run(":" + port) // listen and serve
}
