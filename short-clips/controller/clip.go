package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l33t3mr/go-projects/shortclips/model"
)

var GetClipRoute string = "/clip/:id"
var GetAllCLIPRoute string = "/clips"
var PostClipRoute string = "/clip"

func GetClip(c *gin.Context) {
	id := c.Param("id")
	var clip model.Clip
	result := model.DB.Preload("Directors").Find(&clip, id).Error

	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Clip with this ID not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, clip)
}

func GetAllClips(c *gin.Context) {
	var clips []model.Clip
	result := model.DB.Find(&clips).Error
	if result != nil {
		c.JSON(http.StatusInternalServerError, "No Clip found")
		return
	}
	c.IndentedJSON(http.StatusOK, clips)
}

func PostClip(c *gin.Context) {
	isbn := c.PostForm("isbn")
	title := c.PostForm("title")
	directorFirstName := c.PostForm("director_firstname")
	directorLastName := c.PostForm("director_lastname")
	clip := model.Clip{
		ISBN:  isbn,
		Title: title,
	}
	director := model.Director{
		Firstname: directorFirstName,
		Lastname:  directorLastName,
	}
	director.Clips = append(director.Clips, &clip)
	clip.Directors = append(clip.Directors, &director)
	result := model.DB.Create(&clip)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to save: %s", result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("ID: %d, ", clip.ID),
	})
}
