package helpers

import "path/filepath"

func IsZipFile(filename string) bool {
	// Проверяем расширение файла, чтобы определить, что это ZIP-архив
	ext := filepath.Ext(filename)
	return ext == ".zip"
}
