package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Парсинг HTML файла и отправка содержимого клиенту
func HandleGetHTML(w http.ResponseWriter, r *http.Request) {
	// Путь к файлу для парсинга
	// fileName := "index.html"

	// file, err := os.Open(fileName)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// defer file.Close()

	r.ParseMultipartForm(10 << 20) // 10 MB

	// получаем файл из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusConflict)
		return
	}
	// закрываем файл
	defer file.Close()

	// Чтение файла для парсинга
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Передача информации клиенту
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

}

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20) // 10 MB

	// получаем файл из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	// закрываем файл
	defer file.Close()

	// Чтение файла из формы
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Конвертируем текст из файла с помощью пакета текст-Морзе
	conentConv, err := service.Conv(string(content))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Перегоняем конвертированный текст в байты
	contentConvByte := []byte(conentConv)

	// Формируем название файла для хранения конвертированного текста
	nameResultFile := time.Now().UTC().String()

	// Записываем конвертированный текст в файл
	err = os.WriteFile(nameResultFile, contentConvByte, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(contentConvByte)
}
