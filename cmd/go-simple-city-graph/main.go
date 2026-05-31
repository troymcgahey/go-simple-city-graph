package main

import (
	"go-simple-city-graph/internal/graph"
	"go-simple-city-graph/utils"
)

func main() {

	var citiesArray [][]string

	utils.ReadCitiesFromFile("cityConnections.csv", &citiesArray)

	cityList := make(map[string]*graph.CityNode)

	graph.CreateGraph(&citiesArray, &cityList)

}
