package responses

type FileAPIRes struct {
	File FileRes `json:"file"`
}

type FileRes struct {
	ID       string `json:"id"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	FileSize int64  `json:"fileSize"`
	MimeType string `json:"mimeType"`
}
