package controllers

import (
	"MyArchiveServer/app/models"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
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
	author := models.AUTHOR{
		Name: name,
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
	author := models.AUTHOR{
		Name: name,
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
	authorId := ctx.Param("id")
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
	go func() {
		var authors []models.AUTHOR
		reader := csv.NewReader(file)
		var line []string
		for {
			line, err = reader.Read()
			if err != nil {
				break
			}
			author := models.AUTHOR{
				Name: line[0],
				NameKana: line[1],
			}
			authors = append(authors, author)
		}
		for _, author := range authors {
			models.CreateAuthor(&author)
		}
	}()// TODO 非同期管理テーブル作成
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": "success csv file upload"})
}