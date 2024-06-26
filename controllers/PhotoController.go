package controllers

import (
	"net/http"
	"rakamin-go/app"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController struct {
	DB *gorm.DB
}

func (ctrl *PhotoController) UploadPhoto(c *gin.Context) {
    var photo app.Photo
    if err := c.ShouldBind(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

	user, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := user.(uint)  
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error: User ID assertion failed"})
		return
	}
	photo.UserID = userID


    file, err := c.FormFile("photo")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Photo upload error"})
        return
    }
    filePath := "uploads/" + file.Filename
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the photo"})
        return
    }
    photo.PhotoUrl = filePath

    if result := ctrl.DB.Create(&photo); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Photo uploaded successfully", "photo": photo})
}

func (ctrl *PhotoController) GetPhotos(c *gin.Context) {
    user, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    userID, ok := user.(uint)  
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error: User ID assertion failed"})
        return
    }

    var photos []app.Photo
    if result := ctrl.DB.Where("user_id = ?", userID).Find(&photos); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"photos": photos})
}


func (ctrl *PhotoController) UpdatePhoto(c *gin.Context) {
	photoID := c.Param("id")
	var photoUpdates app.Photo
	if err := c.ShouldBindJSON(&photoUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	result := ctrl.DB.Model(&app.Photo{}).Where("id = ?", photoID).Updates(photoUpdates)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func (ctrl *PhotoController) DeletePhoto(c *gin.Context) {
	photoID := c.Param("id")
	result := ctrl.DB.Delete(&app.Photo{}, photoID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No photo found with provided ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}