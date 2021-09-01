package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ward-go/common"
	"ward-go/router"
)

//go:embed static
var static embed.FS

//go:embed  view
var view embed.FS

func main() {

	//gin.SetMode(gin.ReleaseMode)

	engine := router.Router()

	engine.StaticFS("/public", http.FS(static))
	t, _ := template.ParseFS(view, "view/**/*")
	engine.SetHTMLTemplate(t)
	
	fmt.Println("Start server...")
	err := engine.Run(common.Config.Addr)
	if err != nil {
		log.Fatalln(err)
	}

}
