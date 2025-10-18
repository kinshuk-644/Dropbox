package models

import "time"

// FileMetadata struct with GORM model definitions
type FileMetadata struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Filename     string    `json:"filename"`
	OriginalName string    `json:"original_name"`
	MimeType     string    `json:"mimetype"`
	Size         int64     `json:"size"`
	UploadDate   time.Time `json:"upload_date"`
}

// TableName explicitly tells GORM what the table's name is
func (FileMetadata) TableName() string {
	return "files"
}
