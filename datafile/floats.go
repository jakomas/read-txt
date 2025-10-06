package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64           // объявление возвращаемого массива
	file, err := os.Open("data.txt") //откр файл для чтения, проверить операцию на ошибку
	if err != nil {
		return numbers, err
	}

	i := 0 // счётчик проходов по файлу - читает построчно,
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) //строку в число, полож в массив
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close() // закрыть файл, если при закр есть ошибка вернуть её
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
