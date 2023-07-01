package main

import (
	"os"
	"strings"
)

func renameFile(dir string) {
	files, _ := os.ReadDir(dir)
	str := "【海量资源：】"
	for _, file := range files {
		if file.IsDir() {
			renameFile(dir + "\\" + file.Name())
		} else {
			if strings.Contains(file.Name(), str) {
				_ = os.Rename(dir+"\\"+file.Name(), dir+"\\"+strings.Join(strings.Split(file.Name(), str), ""))
			}
		}
	}
}

func main() {
	dir := "D:\\"

	renameFile(dir)
}
