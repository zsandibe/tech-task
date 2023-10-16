package models

type Archive struct {
	FileName    string     `json:"file_name"`
	ArchiveSize float64    `json:"archive_size"`
	TotalSize   float64    `json:"total_size"`
	TotalFiles  float64    `json:"total_files"`
	Files       []FileInfo `json:"files"`
}
