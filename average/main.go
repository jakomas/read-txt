package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"main.go/datafile"
)

func main() {
	//datafile.GetFloats("data.txt")
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// создадим и инициализируем массив
	expArr, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Массив из строк файла:%v\n", expArr)

	var sum float64
	for _, num := range expArr {
		sum += num
	}
	fmt.Printf("Вывод суммы чисел массива:%2.2f \n", sum)
}
