package api

import (
	"fmt"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Starts listening for requests on the given port
func HandleRequests(port int) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// GETs

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Routes

// GETs

// Writes given multipart form data object to the file specified
func writeMultiPartFormDataToDisk(multipartFormData io.ReadCloser, destFile string) error {
	tempFile, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	io.Copy(tempFile, multipartFormData)

	return nil
}
