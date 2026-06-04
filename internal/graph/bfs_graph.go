package graph

import (
	"fmt"
)

func SearchGraphBFS(cityList map[string]*CityNode, startCity string, destinationCity string) bool {

	visitedCity := make(map[string]bool)

	searchCityList := []*CityNode{cityList[startCity]}

	for len(searchCityList) > 0 {

		currentCity := searchCityList[0]

		fmt.Println("currentCity ", currentCity.City, " destinationCity ", destinationCity)
		searchCityList = searchCityList[1:]

		if currentCity.City == destinationCity {
			return true
		}

		visitedCity[currentCity.City] = true

		searchCityList = append(searchCityList, currentCity.Connections...)

	}

	return false
}
