package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetStyle(style string) [][]string {
	var chrs [][]string
	var currRow []string
	var standard []byte
	var err error
	if style == "standard" {
		standard, err = os.ReadFile("styles/standard.txt")
	} else if style == "shadow" {
		standard, err = os.ReadFile("styles/shadow.txt")
	} else if style == "thinkertoy" {
		standard, err = os.ReadFile("styles/thinkertoy.txt")
		var correct string
		correct = strings.ReplaceAll(string(standard), string(13), "")
		standard = []byte(correct)
	} else {
		fmt.Println("Wrong style")
	}
	if err != nil {
		fmt.Println(err)
	}
	Checkhash(string(standard))
	standard = standard[1:]
	tempChrs := strings.Split(string(standard), "\n\n")
	for j := 0; j < len(tempChrs); j++ {
		currRow = strings.Split(tempChrs[j], "\n")
		chrs = append(chrs, currRow)
	}
	return chrs
}
