// package utils

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// func Run(Text string, Style string) error {
// 	// Проверяем наличие входных данных
// 	if Text == "" {
// 		return fmt.Errorf("text input is empty")
// 	}

// 	// Создаем файл для записи аски-арта
// 	filename := "sample.txt"
// 	currentDir, err := os.Getwd()
// 	if err != nil {
// 		return fmt.Errorf("failed to get current directory: %v", err)
// 	}
// 	filePath := filepath.Join(currentDir, filename)
// 	f, err := os.Create(filePath)
// 	if err != nil {
// 		return fmt.Errorf("failed to create file: %v", err)
// 	}
// 	defer f.Close() // Закрываем файл в конце функции

// 	// Разбиваем текст на строки
// 	words := strings.Split(Text, "\\n")

// 	// Получаем стили и записываем аски-арт в файл
// 	styleChrs := GetStyle(Style)
// 	Output(styleChrs, words, f)

// 	return nil
// }

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Run(Text string, Style string) {
	// Проверяем количество аргументов командной строки
	filename := "sample.txt"
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir, filename)
	f, e := os.Create(filePath)
	if e != nil {
		fmt.Println("Error creating file:", e)
		os.Exit(1)
	}
	//проверка входных данных есть ли они в аски
	cor := CheckInput(Text)
	if cor == true {
		//получение массива строк разделяя каждую строку
		words := strings.Split(string(Text), "\\n")
		//получение файла c стилями
		styleChrs := GetStyle(Style)
		//вывод
		Output(styleChrs, words, f)
	}
}
