package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"log"
	"main.go/config"
	"main.go/pkg/db/model"
	"net/http"
	"os"
)

func UploadFile(c echo.Context) error {
	file := &model.File{}
	if err := c.Bind(file); err != nil {
		return err
	}

	fmt.Println(file)
	client, _ := config.LoadDB("")
	bucket, err := gridfs.NewBucket(
		client.Database("myfiles"),
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	uploadStream, err := bucket.OpenUploadStream(
		file.FileName,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(file.File)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
	return c.String(http.StatusCreated, "file upload done")
}
