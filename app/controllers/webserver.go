package controllers

import (
	"MyArchiveServer/app/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.LoadHTMLGlob("app/views/*.html")
}

func StartWebServer() {
	router.GET("/", viewTopHandler)

	truncateRoute := router.Group("/truncate")
	{
		truncateRoute.GET("/author", truncateAuthor)
		truncateRoute.GET("/record", truncateRecord)
	}

	webRoute := router.Group("/web")
	{
		authorRoute := webRoute.Group("/author")
		{
			authorRoute.GET("/file", webAuthorCsvFile)
			authorRoute.POST("/upload", webAuthorCsvUpload)
		}
		recordRoute := webRoute.Group("/record")
		{
			recordRoute.GET("/file", webRecordCsvFile)
			recordRoute.POST("/upload", webRecordCsvUpload)
		}
	}

	apiRoute := router.Group("/api")
	{
		recordRoute := apiRoute.Group("/record")
		{
			recordRoute.GET("/search-title", apiSearchByTitle)
			recordRoute.GET("/search-author", apiSearchByAuthor)
			recordRoute.POST("/new", apiCreateRecord)
			recordRoute.GET("/edit", apiFindRecord)
			recordRoute.POST("/update", apiUpdateRecord)
			recordRoute.POST("/delete", apiDeleteRecord)
		}
		{
			authorRoute := apiRoute.Group("/author")
			authorRoute.POST("/new", apiCreateAuthor)
			authorRoute.GET("/edit/:id", apiFindAuthor)
			authorRoute.POST("/update", apiUpdateAuthor)
			authorRoute.POST("/delete/:id", apiDeleteAuthor)
		}
	}

	err := router.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func viewTopHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": ""})
}

func webRecordCsvFile(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "recordsUpload.html", gin.H{"message": ""})
}

func webAuthorCsvFile(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "authorsUpload.html", gin.H{"message": ""})
}

func truncateAuthor(ctx *gin.Context) {
	models.TruncateAuthor()
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": "truncate Author done"})
}

func truncateRecord(ctx *gin.Context) {
	models.TruncateRecord()
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": "truncate Record done"})
}

func showUploadError(ctx *gin.Context, message string) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": message})
}