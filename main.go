package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(storage []byte, tmp byte) bool {
	size := len(storage)
	for i := 0; i < size; i++ {
		if storage[i] == tmp {
			return true
		}
	}
	return false
}
func main() {
	fmt.Printf("Идет выполнение..\n")
	file, err := os.ReadFile("Sourse.txt")
	if err != nil {
		fmt.Printf("Ошибка чтения из файла [%s]", err)
		return
	}
	size := len(file) - 1

	storage := make([]byte, size)
	var g int = 0
	test := string(file)
	_ = test
	for i := size; i != 0; i-- {
		if file[i] == 10 {
			i = i - 1
		}
		tmp := file[i]
		if check(storage, tmp) && tmp != 32 {
			file[i] = 32
		} else {
			storage[g] = file[i]
			g = g + 1
		}
	}
	fileName := "Sourse.txt"
	eror := ioutil.WriteFile(fileName, file, 0666)
	fmt.Printf("Конец")
	if eror != nil {
		fmt.Printf("Ошибка записи в файл [%s]", err)
		return
	}
}
