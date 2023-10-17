package sendEmail

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSendEmailHandler_SendEmail(t *testing.T) {
	// Создаем новый инстанс Gin для тестирования
	r := gin.New()

	// Добавляем маршрут для обработки запросов
	sendEmailHandler := &SendEmailHandler{} // Замените на вашу реализацию
	r.POST("/api/mail/file", sendEmailHandler.SendEmail)

	// Тест 1: отправка правильного файла и списка почт
	t.Run("ValidRequest", func(t *testing.T) {
		// Создаем тестовый запрос
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fileWriter, _ := writer.CreateFormFile("file", "document.pdf")
		fileWriter.Write([]byte("binary_data_of_file"))
		writer.WriteField("emails", "elonmusk@x.com,jeffbezos@amazon.com,zuckerberg@meta.com")
		writer.Close()

		req := httptest.NewRequest("POST", "/api/mail/file", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// Создаем тестовый Recorder и выполняем запрос
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		// Проверяем статус-код и ожидаемый ответ
		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"Email sent successfully"}`
		if w.Body.String() != expectedResponse {
			t.Errorf("Expected response %s, but got %s", expectedResponse, w.Body.String())
		}
	})

	// Тест 2: отправка файла с недопустимым MIME-типом
	t.Run("InvalidMimeType", func(t *testing.T) {
		// Создаем тестовый запрос с недопустимым MIME-типом
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fileWriter, _ := writer.CreateFormFile("file", "invalid.jpg")
		fileWriter.Write([]byte("binary_data_of_file"))
		writer.WriteField("emails", "elonmusk@x.com,jeffbezos@amazon.com,zuckerberg@meta.com")
		writer.Close()

		req := httptest.NewRequest("POST", "/api/mail/file", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// Создаем тестовый Recorder и выполняем запрос
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Проверяем статус-код и ожидаемый ответ
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, but got %d", http.StatusBadRequest, w.Code)
		}
	})

	// Добавьте другие тесты для обработки других сценариев
}
