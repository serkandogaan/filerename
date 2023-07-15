package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	filePath := "C:/Users/Serkan/Desktop/dosyalar" // Dosya yaolunu buraya girin
	var input string
	fmt.Scanln(&input)

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println("Klasör okunamadi, Dosya yolunu tekrar girin : ", err)
		return
	}

	for i, file := range files {

		oldName := file.Name()
		newName := strconv.Itoa(i + 1)

		if i+1 < 10 {
			newName = "0" + newName
		}
		newName = newName + filepath.Ext(oldName)

		oldPath := filepath.Join(filePath, oldName) // isim değişmeden önceki yol adresi
		newPath := filepath.Join(filePath, newName) // isim değişince oluşacak yol adresi

		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Değiştirilemedi %s", err)
		} else {
			fmt.Printf("Değiştirildi %s \n", newName)
		}
	}

}
