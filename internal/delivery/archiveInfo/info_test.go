package archiveInfo

import (
	"bytes"
	"doodocsTask/internal/models"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockInfo is a mock implementation of the ArchiveInfoProvider interface for testing.
type MockInfo struct{}

func (m *MockInfo) GetInfoByArchive(fileData []byte, fileName string) (*models.Archive, error) {

	return &models.Archive{
		FileName:    fileName,
		ArchiveSize: 4102029.312,
		TotalSize:   6836715.52,
		TotalFiles:  2,
		Files: []models.FileInfo{
			{
				FilePath: "photo.jpg",
				Size:     2516582.4,
				MimeType: "image/jpeg",
			},
			{
				FilePath: "directory/document.docx",
				Size:     4320133.12,
				MimeType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			},
		},
	}, nil
}

func TestArchiveInfoHandler_GetInfoByArchive(t *testing.T) {

	router := gin.Default()

	infoProvider := &MockInfo{}

	handler := NewArchiveInfoHandler(infoProvider)

	router.POST("/api/archive/information", handler.GetInfoByArchive)

	t.Run("Successful request", func(t *testing.T) {

		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		part, _ := w.CreateFormFile("file", "my_archive.zip")
		io.WriteString(part, "mock binary data of ZIP file")
		w.Close()

		req, _ := http.NewRequest("POST", "/api/archive/information", &b)
		req.Header.Set("Content-Type", w.FormDataContentType())

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
		}

		expectedResponse := `{
			"file_name": "my_archive.zip",
			"archive_size": 4102029.312,
			"total_size": 6836715.52,
			"total_files": 2,
			"files": [
				{
					"file_path": "photo.jpg",
					"size": 2516582.4,
					"mime_type": "image/jpeg"
				},
				{
					"file_path": "directory/document.docx",
					"size": 4320133.12,
					"mime_type": "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
				}
			]
		}`

		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})

}
