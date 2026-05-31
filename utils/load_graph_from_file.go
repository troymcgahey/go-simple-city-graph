package utils

import (
	"bufio"
	"os"
)

func ReadCitiesFromFile(fileName string, cities *[]string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		city := scanner.Text()
		*cities = append(*cities, city)
	}
	return scanner.Err()
}
