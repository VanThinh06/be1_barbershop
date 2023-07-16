package api

import (
	"errors"
	"fmt"
	"image"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) uploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	dts := "./assets/upload/"
	// Upload the file to specific dst.
	err = ctx.SaveUploadedFile(file, dts+"/"+file.Filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	 
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	
	file.Open()

}

func (server *Server) loadImageFromURL(URL string) (image.Image, error) {

	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
