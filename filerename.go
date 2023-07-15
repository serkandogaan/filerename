package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "rename [dosya_yolu]",
	Short: "Dosya adi degistirme",
	Long:  "dosya adi degistirir",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		// filePath := "C:/Users/Serkan/Desktop/dosyalar" // Dosya yaolunu buraya girin
		filePath := os.Args[2]
		// var input string
		// fmt.Scanln(&input)
		// filePath = input

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

	},
}

func main() {

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
