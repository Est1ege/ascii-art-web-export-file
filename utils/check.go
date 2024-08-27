package utils

import (
	"errors"
	"fmt"
	"hash/fnv"
	// "os"
)

func Checkhash(style string) error {
	expectedHashes := map[string]uint32{
		"standard.txt":   1766917683,
		"shadow.txt":     4281396044,
		"thinkertoy.txt": 3930937207,
	}
	hasher := fnv.New32a()
	hasher.Write([]byte(style))
	hashValue := hasher.Sum32()
	for _, expectedHash := range expectedHashes {
		if hashValue == expectedHash {
			return nil // Хеш совпадает, возвращаем nil (без ошибки)
		}
	}

	return errors.New("wrong style file") // Если хеш не совпадает, возвращаем ошибку
}

// func Checkhash(style string) {
// 	// 1766917683
// 	// мы получаем хэш файла
// 	hasher := fnv.New32a()
// 	hasher.Write([]byte(style))
// 	hashValue := hasher.Sum32()
// 	if hashValue != 1766917683 && hashValue != 4281396044 && hashValue != 3075161722 {
// 		fmt.Println("ERROR : wrong style file")
// 		os.Exit(0)
// 	}
// }

func CheckInput(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 127 {
			if s[i] == 10 {
				continue
			}
			fmt.Println("ERROR: wrong format")
			return false
		}
	}
	return true
}
