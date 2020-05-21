package controllers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.LoadHTMLGlob("app/views/*.html")
	router.Use(cors.Default())
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
		router.GET("/web/async/console", webSearchAsync)
	}

	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/healthCheck", healthCheck)
		{
			recordRoute := apiRoute.Group("/record")
			recordRoute.GET("/search-title", apiSearchByTitle)
			recordRoute.GET("/search-author", apiSearchByAuthor)
			recordRoute.POST("/new", apiCreateRecord)
			recordRoute.GET("/edit", apiFindRecord)
			recordRoute.POST("/update", apiUpdateRecord)
			recordRoute.POST("/delete", apiDeleteRecord)
		}
		{
			authorRoute := apiRoute.Group("/author")
			authorRoute.GET("/search", apiFindAuthorByName)
			authorRoute.POST("/new", apiCreateAuthor)
			authorRoute.GET("/edit", apiFindAuthor)
			authorRoute.POST("/update", apiUpdateAuthor)
			authorRoute.POST("/delete", apiDeleteAuthor)
			authorRoute.GET("/count", ApiFindCountByAuthor)
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

func showUploadError(ctx *gin.Context, message string) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": message})
}

type healthCheckStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func healthCheck(ctx *gin.Context) {
	var status healthCheckStatus
	status.Status = "OK"
	status.Message = "health check is ok"
	ctx.JSON(http.StatusOK, status)
}
