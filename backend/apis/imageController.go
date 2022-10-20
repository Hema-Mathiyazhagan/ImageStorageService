package apis

import (
	"io"
	"net/http"
	"os"
	"log"
	"path/filepath"

	//"log"

	"github.com/Hema-Mathiyazhagan/image-storage-service/lib"
	"github.com/Hema-Mathiyazhagan/image-storage-service/model"

	"github.com/gin-gonic/gin"
)

// ListImages handler returns list of images of the specified album
// @Summary     list of images
// @Description returns list of images of the specified album
// @Tags        images
// @Produce     json
// @Param       albumName path     string true "Album name"
// @Success     200       {object} model.Imglist_json
// @Failure     400       {object} model.Error
// @Failure     500       {object} model.Error
// @Router      /albums/{albumName}/images [get]
func ListImages(c *gin.Context) {
	albumname := c.Param("albumName")
	tempList, listImagesErr := lib.List_images(albumname)
	if listImagesErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": listImagesErr.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"imageList": tempList})
	}
}

// GetImage handler sends the specified image to rest client
// @Summary     get a image
// @Description get the specified image from the server
// @Tags        images
// @Produce     octet-stream
// @Param       albumName path     string true "Album name"
// @Param       imageName path     string true "Image name"
// @Success     200       {file}   octet-stream
// @Failure     400       {object} model.Error
// @Failure     500       {object} model.Error
// @Router      /albums/{albumName}/images/{imageName} [get]
func GetImage(c *gin.Context) {
	albumname := c.Param("albumName")
	imagename := c.Param("imageName")

	lib.Images_info.RLock()
	_, doesAlbumexist := lib.Images_info.Album_map[albumname]
	lib.Images_info.RUnlock()
	if !doesAlbumexist {
		c.JSON(http.StatusBadRequest, gin.H{"error": albumname + " album is not present"})
		return
	}

	lib.Images_info.RLock()
	_, doesImgexist := lib.Images_info.Album_map[albumname][imagename]
	lib.Images_info.RUnlock()
	if !doesImgexist {
		c.JSON(http.StatusBadRequest, gin.H{"error": imagename + " image is not present"})
		return
	}

	_, statErr := os.Stat(lib.Imagestore_path + albumname + "/" + imagename)
	if os.IsNotExist(statErr) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": imagename + " image is present database but not in storage"})
		return
	}

	header := c.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename=" + imagename}

	file, err := os.Open(lib.Imagestore_path + albumname + "/" + imagename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not sent image " + imagename + " for download"})
		return
	}
	defer file.Close()

	io.Copy(c.Writer, file)
}

// AddImage handler adds an image to the specified album
// @Summary     add a image
// @Description adds an image to the specified album
// @Tags        images
// @Accept      multipart/form-data
// @Produce     json
// @Param       albumName path     string true "Album name"
// @Param       file      formData file   true "Add Image"
// @Success     200       {object} model.Result
// @Failure     400       {object} model.Error
// @Failure     500       {object} model.Error
// @Router      /albums/{albumName}/images [post]
func AddImage(c *gin.Context) {
	albumname := c.Param("albumName")
	file, formErr := c.FormFile("file")
	if formErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": formErr.Error()})
		return
	}

	filename := filepath.Base(file.Filename)
	log.Println(filename)

	lib.Images_info.RLock()
	_, doesAlbumexist := lib.Images_info.Album_map[albumname]
	lib.Images_info.RUnlock()
	if !doesAlbumexist {
		c.JSON(http.StatusBadRequest, gin.H{"error": albumname + " album is not present"})
		return
	}

	lib.Images_info.RLock()
	_, doesImgexist := lib.Images_info.Album_map[albumname][filename]
	lib.Images_info.RUnlock()
	if doesImgexist {
		c.JSON(http.StatusBadRequest, gin.H{"error": filename + " image already present"})
		return
	}

	_, statErr := os.Stat(lib.Imagestore_path + albumname + "/" + filename)
	if !os.IsNotExist(statErr) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": filename + " image already present in storage but not present in database"})
		return
	}

	if fileSaveErr := c.SaveUploadedFile(file, lib.Imagestore_path+albumname+"/"+filename); fileSaveErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fileSaveErr.Error()})
		return
	}

	lib.Images_info.Lock()
	lib.Images_info.Album_map[albumname][filename] = model.Image{
		FileName: filename,
		FilePath: lib.Imagestore_path + albumname + "/" + filename,
	}
	lib.Images_info.Unlock()
	c.JSON(http.StatusOK, gin.H{"message": filename + " image added to album " + albumname})
}

// DeleteImage handler deletes a image from the specified album
// @Summary     delete a image
// @Description deletes a image from the specified album
// @Tags        images
// @Produce     json
// @Param       albumName path     string true "Album name"
// @Param       imageName path     string true "Image name"
// @Success     200       {object} model.Result
// @Failure     400       {object} model.Error
// @Failure     500       {object} model.Error
// @Router      /albums/{albumName}/images/{imageName} [delete]
func DeleteImage(c *gin.Context) {
	albumname := c.Param("albumName")
	imagename := c.Param("imageName")
	if respCode, deleteImageErr := lib.Delete_image(albumname, imagename); deleteImageErr != nil {
		c.JSON(respCode, gin.H{"error": deleteImageErr.Error()})
	} else {
		c.JSON(respCode, gin.H{"message": imagename + " image deleted from album " + albumname})
	}
}

