package main

import (
	"fmt"
	"github.com/labstack/echo"
	"main.go/config"
	"main.go/pkg/db/controller"
	"net/http"
	"os"
)

func main() {
	config.LoadEnvironments()
	fmt.Println("Application started successfully. :)")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to blog APP developed with Golang")
	})
	e.POST("/add-file", controller.UploadFile)
	//e.GET("/download-file", controller.DownloadFile)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
