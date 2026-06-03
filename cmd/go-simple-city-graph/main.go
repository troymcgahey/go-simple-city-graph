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

	//bfs

	//dfs

}
