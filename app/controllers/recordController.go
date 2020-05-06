package controllers

import (
	"MyArchiveServer/app/models"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func apiSearchByTitle(ctx *gin.Context) {
	param := ctx.Query("title")
	results := models.FindByTitle(param)
	ctx.JSON(http.StatusOK, results)
}

func apiSearchByAuthor(ctx *gin.Context) {
	param := ctx.Query("name")
	results := models.FindByAuthor(param)
	ctx.JSON(http.StatusOK, results)
}

func apiCreateRecord(ctx *gin.Context) {
	authorId, err := strconv.Atoi(ctx.PostForm("author_id"))
	title := ctx.PostForm("title")
	if err != nil || len(title) == 0 {
		response := models.ApiResponse{
			Status:  "Error",
			Message: "invalid parameter",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	record := models.RECORD{
		Title :      title,
		TitleKana :  ctx.PostForm("title_kana"),
		Evaluation : "0",
		Author :     authorId,
	}
	models.CreateRecord(&record)
	response := models.ApiResponse{
		Status:  "OK",
		Message: "success created record",
	}
	ctx.JSON(http.StatusOK, response)
}

func apiFindRecord(ctx *gin.Context) {
	recordId := ctx.Query("id")
	results := models.FindById(recordId)
	ctx.JSON(http.StatusOK, results)
}

func apiUpdateRecord(ctx *gin.Context) {
	id := ctx.PostForm("id")
	authorId, err := strconv.Atoi(ctx.PostForm("author_id"))
	title := ctx.PostForm("title")
	if err != nil || len(title) == 0 {
		response := models.ApiResponse{
			Status:  "Error",
			Message: "invalid parameter",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	record := models.RECORD{
		Title :      title,
		TitleKana :  ctx.PostForm("title_kana"),
		Evaluation : ctx.PostForm("evaluation"),
		Author :     authorId,
	}
	models.UpdateRecord(id, &record)
	response := models.ApiResponse{
		Status:  "OK",
		Message: "success updated record",
	}
	ctx.JSON(http.StatusOK, response)
}

func apiDeleteRecord(ctx *gin.Context) {
	recordId := ctx.Query("id")
	if len(recordId) == 0 {
		response := models.ApiResponse{
			Status:  "Error",
			Message: "invalid parameter",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	models.DeleteRecord(recordId)
	response := models.ApiResponse{
		Status:  "OK",
		Message: "success delete record",
	}
	ctx.JSON(http.StatusOK, response)
}

func webRecordCsvUpload(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		showUploadError(ctx, err.Error())
		return
	}
	defer file.Close()
	go func() {
		var records []models.RECORD
		reader := csv.NewReader(file)
		var line []string
		for {
			line, err = reader.Read()
			if err != nil {
				break
			}
			authorId := line[2]
			author, err := strconv.Atoi(authorId)
			if err != nil {
				author = 0
			}
			record := models.RECORD{
				Title: line[0],
				TitleKana: line[3],
				Evaluation: line[1],
				Author: author,
			}
			records = append(records, record)
		}
		fmt.Println(records)
		for _, record := range records {
			models.CreateRecord(&record)
		}
	}()// TODO 非同期管理テーブル作成
	ctx.HTML(http.StatusOK, "index.html", gin.H{"message": "success csv file upload"})
}