package main

import (
	"fmt"
	//"go-simple-city-graph/internal/graph"
	"go-simple-city-graph/utils"
)

//"fmt"

func main() {

	//cityIndex := make(map[string]*graph.CityNode)

	//Build graph

	var cities [][]string

	utils.ReadCitiesFromFile("cityConnections.csv", &cities)

	fmt.Println("Loaded")
	fmt.Println(cities)
}
