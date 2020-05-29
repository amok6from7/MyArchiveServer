package controllers

import (
	"MyArchiveServer/app/models"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func apiCreateAuthor(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if len(name) == 0 {
		response := models.ApiResponse{
			Status:  "Error",
			Message: "invalid parameter",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	author := models.Author{
		Name:     name,
		NameKana: ctx.PostForm("name_kana"),
	}
	models.CreateAuthor(&author)
	response := models.ApiResponse{
		Status:  "OK",
		Message: "success created author",
	}
	ctx.JSON(http.StatusOK, response)
}

func apiFindAuthor(ctx *gin.Context) {
	authorId := ctx.Query("id")
	author := models.ApiFindAuthor(authorId)
	ctx.JSON(http.StatusOK, author)
}

func apiFindAuthorByName(ctx *gin.Context) {
	name := ctx.Query("name")
	author := models.ApiFindAuthorByName(name)
	ctx.JSON(http.StatusOK, author)
}

func apiUpdateAuthor(ctx *gin.Context) {
	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	if len(id) == 0 || len(name) == 0 {
		response := models.ApiResponse{
			Status:  "Error",
			Message: "invalid parameter",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	author := models.Author{
		Name:     name,
		NameKana: ctx.PostForm("name_kana"),
	}
	models.UpdateAuthor(id, &author)
	response := models.ApiResponse{
		Status:  "OK",
		Message: "success updated author",
	}
	ctx.JSON(http.StatusOK, response)
}

func apiDeleteAuthor(ctx *gin.Context) {
	authorId := ctx.PostForm("id")
	if len(authorId) == 0 {
		response := models.ApiResponse{
			Status:  "Error",
			Message: "invalid parameter",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	models.DeleteAuthor(authorId)
	response := models.ApiResponse{
		Status:  "OK",
		Message: "success delete author",
	}
	ctx.JSON(http.StatusOK, response)
}

func webAuthorCsvUpload(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		showUploadError(ctx, err.Error())
		return
	}
	defer file.Close()
	id := models.CreateAsyncManage("webAuthorCsvUpload")

	var authors []models.Author
	reader := csv.NewReader(file)
	var line []string
	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		author := models.Author{
			Name:     line[0],
			NameKana: line[1],
		}
		authors = append(authors, author)
	}

	var wg sync.WaitGroup
	for _, author := range authors {
		wg.Add(1)
		go func(author models.Author) {
			defer wg.Done()
			models.CreateAuthor(&author)
		}(author)
	}
	wg.Wait()
	models.UpdateAsyncManage(id)
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": "success csv file upload"})
}

func ApiFindCountByAuthor(ctx *gin.Context) {
	count := models.ApiFindCountByAuthor()
	ctx.JSON(http.StatusOK, count)
}

func webAuthorCsvFile(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "authorsUpload.html", gin.H{"message": ""})
}

func truncateAuthor(ctx *gin.Context) {
	models.TruncateAuthor()
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": "truncate Author done"})
}
