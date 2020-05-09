package controllers

import (
	"MyArchiveServer/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func webSearchAsync(ctx *gin.Context) {
	asyncManages := models.FindAsyncManage()
	ctx.HTML(http.StatusOK, "asyncManage.html", gin.H{"data": asyncManages})
}
