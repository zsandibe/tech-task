package archiveInfo

import (
	"archive/zip"
	"bytes"
	"doodocsTask/internal/models"
	"doodocsTask/pkg/helpers"
	"errors"
	"mime"
	"path/filepath"
)

func (s *ArchiveInfoService) GetInfoByArchive(fileData []byte, fileName string) (*models.Archive, error) {

	if !helpers.IsZipFile(fileName) {
		return nil, errors.New("Not a valid ZIP file")
	}

	zipReader, err := zip.NewReader(bytes.NewReader(fileData), int64(len(fileData)))
	if err != nil {
		return nil, err
	}

	var totalSize float64
	var files []models.FileInfo

	// Проходим по файлам в архиве
	for _, file := range zipReader.File {
		// Получаем информацию о файле
		info := models.FileInfo{
			FilePath: file.Name,
			Size:     float64(file.UncompressedSize64),
			MimeType: mime.TypeByExtension(filepath.Ext(file.Name)),
		}

		totalSize += info.Size
		files = append(files, info)
	}

	archiveInfo := &models.Archive{
		FileName:    fileName,
		ArchiveSize: float64(len(fileData)),
		TotalSize:   totalSize,
		TotalFiles:  float64(len(files)),
		Files:       files,
	}

	return archiveInfo, nil
}
