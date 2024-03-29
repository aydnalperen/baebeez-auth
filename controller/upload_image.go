package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadImage(ctx *gin.Context) {

	uid, err := utils.ExtractTokenUID(ctx)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Request.ParseMultipartForm(10 << 20)

	file, handler, err := ctx.Request.FormFile("file")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	path := "images"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}
	}

	dst, err := os.Create("images/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	models.DB.Model(&models.User{}).Where("uid = ?", uid).Update("photo", "images/"+handler.Filename)

	MakeCompleted(ctx)
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func MakeCompleted(ctx *gin.Context) { //
	uid, _ := utils.ExtractTokenUID(ctx)
	if result := models.DB.Model(&models.User{}).Where("uid=?", uid).Update("iscomp", 1); result != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}
