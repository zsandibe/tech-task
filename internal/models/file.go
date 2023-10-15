package models

type FileInfo struct {
	FilePath string `json:"file_path"`
	Size float64 `json:"size"`
	MimeType string `json:"mime_type"`
}