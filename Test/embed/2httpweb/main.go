package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
)

//go:embed templates
var tmpl embed.FS

//go:embed static
var static embed.FS

func main() {
	r := gin.Default()
	t, _ := template.ParseFS(tmpl, "templates/*.tmpl")
	r.SetHTMLTemplate(t)
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.tmpl", gin.H{"title": "Golang Embed 测试"})
	})
	r.Run(":8080")
}
