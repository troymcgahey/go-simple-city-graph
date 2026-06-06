package main

import (
	"fmt"
	"go-simple-city-graph/internal/graph"
	"go-simple-city-graph/utils"
	"os"
)

func main() {

	var citiesArray [][]string

	utils.ReadCitiesFromFile("cityConnections.csv", &citiesArray)

	cityList := make(map[string]*graph.CityNode)

	graph.CreateGraph(citiesArray, cityList)

	//prompt for a starting city and find a path to a destination city

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go Source_City Destination_City")
		return
	}

	//validate source and destination exists in CityList
	//if not, errror and exist

	if _, exists := cityList[os.Args[1]]; !exists {
		fmt.Println("Source city does not exist")
		return
	}

	if _, exists := cityList[os.Args[2]]; !exists {
		fmt.Println("Destination city does not exist")
	}

	if graph.SearchGraphBFS(cityList, os.Args[1], os.Args[2]) {
		fmt.Println("Connection route found")
	} else {
		fmt.Println("Connection route not found")
	}

}
