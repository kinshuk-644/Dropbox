package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/kinshuk-644/Dropbox/backend/connection" // Import database
	"github.com/kinshuk-644/Dropbox/backend/models"     // Import models
)

// --- FileMetadata struct and TableName function are REMOVED from here ---

// --- API Handlers (Gin & GORM version) ---

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	// ... (rest of the error handling)

	uniqueFilename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)
	dstPath := filepath.Join("uploads", uniqueFilename)

	if err := c.SaveUploadedFile(file, dstPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		return
	}

	// Use 'models.FileMetadata'
	meta := models.FileMetadata{
		Filename:     uniqueFilename,
		OriginalName: file.Filename,
		MimeType:     file.Header.Get("Content-Type"),
		Size:         file.Size,
		UploadDate:   time.Now(),
	}

	// Use 'database.DB'
	result := connection.DB.Create(&meta)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file metadata"})
		return
	}

	c.JSON(http.StatusCreated, meta)
}

func GetFiles(c *gin.Context) {
	var files []models.FileMetadata // Use 'models.FileMetadata'

	// Use 'database.DB'
	result := connection.DB.Order("upload_date desc").Find(&files)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve files"})
		return
	}

	c.JSON(http.StatusOK, files)
}

func DownloadFile(c *gin.Context) {
	id := c.Param("id")
	var fileMeta models.FileMetadata // Use 'models.FileMetadata'

	// Use 'database.DB'
	result := connection.DB.First(&fileMeta, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	filePath := filepath.Join("uploads", fileMeta.Filename)
	c.FileAttachment(filePath, fileMeta.OriginalName)
}
