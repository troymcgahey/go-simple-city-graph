package utils

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadCitiesFromFile(fileName string, cities *[][]string) error {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		*cities = append(*cities, record)
	}

	return err
}
