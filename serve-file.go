package main

import (
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"os"

	"io"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	router := gin.Default()
	router.GET("/1MB.pdf", pdf)
	router.GET("/1MB.docx", docx)
	router.Run(":8092")
}

func pdf(c *gin.Context) {
	c.File("1MB.pdf")

}

func docx(c *gin.Context) {
	c.File("1MB.docx")

}

