package utils

import (
	"os"
)

func Output(chrs [][]string, words []string, f *os.File) {
	for _, currWord := range words {
		//проверка на /n
		if currWord == "" {
			f.WriteString("\n")
		} else {
			// используем тройной цикл для вывода
			for j := 0; j < 8; j++ {
				for _, ch := range currWord {
					if ch >= 32 && ch <= 127 {
						for k := 0; k < len(chrs[ch-32][j]); k++ {
							f.WriteString(string(chrs[ch-32][j][k]))
						}
					}
				}
				// Записываем перевод строки в конце каждой строки баннера
				f.WriteString("\n")
			}
		}
	}
}
