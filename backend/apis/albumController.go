package apis

import (
	"net/http"

	"github.com/Hema-Mathiyazhagan/image-storage-service/lib"
	"github.com/Hema-Mathiyazhagan/image-storage-service/model"

	"github.com/gin-gonic/gin"
)

// AddAlbum handler creates a album
// @Summary     create a album
// @Description creates new album if not present
// @Tags        albums
// @Accept      json
// @Produce     json
// @Param       album body     model.Alb_json true "Add Album"
// @Success     200   {object} model.Result
// @Failure     400   {object} model.Error
// @Failure     500   {object} model.Error
// @Router      /albums [post]
func AddAlbum(c *gin.Context) {
	var addAlbum model.Alb_json
	if jsonbinderr := c.ShouldBindJSON(&addAlbum); jsonbinderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": jsonbinderr.Error()})
		return
	}
	if respCode, addAlbumErr := lib.Add_Album(addAlbum.AlbumName); addAlbumErr != nil {
		c.JSON(respCode, gin.H{"error": addAlbumErr.Error()})
	} else {
		c.JSON(respCode, gin.H{"message": addAlbum.AlbumName + " album added"})
	}
}

// DeleteAlbum handler deletes a album
// @Summary     delete a album
// @Description deletes the specified album if present
// @Tags        albums
// @Produce     json
// @Param       albumName path     string true "Delete Album"
// @Success     200       {object} model.Result
// @Failure     400       {object} model.Error
// @Failure     500       {object} model.Error
// @Router      /albums/{albumName} [delete]
func DeleteAlbum(c *gin.Context) {
	albumname := c.Param("albumName")
	if respCode, deleteAlbumErr := lib.Delete_Album(albumname); deleteAlbumErr != nil {
		c.JSON(respCode, gin.H{"error": deleteAlbumErr.Error()})
	} else {
		c.JSON(respCode, gin.H{"message": albumname + " album deleted"})
	}
}

