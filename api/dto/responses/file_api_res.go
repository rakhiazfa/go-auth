package responses

import "github.com/google/uuid"

type FileAPIRes struct {
	Message string  `json:"message"`
	File    FileRes `json:"file"`
}

type FileRes struct {
	ID       uuid.UUID `json:"id"`
	FileName string    `json:"fileName"`
	FileType string    `json:"fileType"`
	FileSize int64     `json:"fileSize"`
	MimeType string    `json:"mimeType"`
}
