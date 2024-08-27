package main

import (
	"ascii-art-web-export-file/utils"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Открываем файл sample.txt
	file, err := os.Open("sample.txt")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	// Устанавливаем заголовок Content-Disposition для указания имени файла
	w.Header().Set("Content-Disposition", "attachment; filename=sample.txt")

	// Копируем содержимое файла в тело ответа
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error copying file contents: %v", err)
		return
	}
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	var tmplPath string
	// Определяем путь к шаблону в зависимости от запрошенного URL
	if r.URL.Path == "/ascii-art" {
		tmplPath = filepath.Join("templates", "asciiart.html")
	} else if r.URL.Path == "/" {
		tmplPath = filepath.Join("templates", "main.html")
	} else {
		tmplPath = filepath.Join("templates", "404.html")
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}

func generateASCIIArtHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mainPageHandler(w, r)
	case "POST":
		// Проверяем текст на правильный формат
		text := r.FormValue("text")
		if !utils.CheckInput(text) {
			http.Redirect(w, r, "/400", http.StatusBadRequest)
			return
		}

		// Получаем имя баннера
		banner := r.FormValue("banner")

		// Проверяем стиль файла
		styleFile := filepath.Join("styles", banner+".txt")
		style, err := os.ReadFile(styleFile)
		if err != nil {
			if os.IsNotExist(err) { // Проверяем, была ли ошибка о том, что файл не найден
				http.Redirect(w, r, "/500", http.StatusInternalServerError)
				tmplPath := filepath.Join("templates", "500.html")
				tmpl, err := template.ParseFiles(tmplPath)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					log.Printf("Error parsing template: %v", err)
					return
				}
				if err := tmpl.Execute(w, nil); err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					log.Printf("Error executing template: %v", err)
					return
				}
				return
			}
			http.Redirect(w, r, "/500", http.StatusInternalServerError)
			return
		}
		if err := utils.Checkhash(string(style)); err != nil {
			http.Redirect(w, r, "/500", http.StatusInternalServerError)
			return
		}

		// Генерируем ASCII Art
		utils.Run(text, banner)

		// Читаем результат из файла
		result, err := os.ReadFile("sample.txt")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error reading file: %v", err)
			return
		}
		asciiArt := string(result)
		// Отображаем результат в шаблоне
		tmpl, err := template.ParseFiles("templates/main.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}
		tmpl.Execute(w, asciiArt)
	default:
		http.Redirect(w, r, "/400", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", generateASCIIArtHandler)
	http.HandleFunc("/download", downloadHandler)
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	})

	http.HandleFunc("/400", func(w http.ResponseWriter, r *http.Request) {
		tmplPath := filepath.Join("templates", "400.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusBadRequest)
			log.Printf("Error parsing template: %v", err)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusBadRequest)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		tmplPath := filepath.Join("templates", "500.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
