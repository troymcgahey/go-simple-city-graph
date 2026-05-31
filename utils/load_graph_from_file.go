package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ReadCitiesFromFile(fileName string, cities *[][]string) error {

	fmt.Println(fileName)

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

		fmt.Println(record)
		*cities = append(*cities, record)
	}

	return err
}
