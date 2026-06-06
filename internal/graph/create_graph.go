package graph

import (
	"strings"
)

//"fmt"

func CreateGraph(citiesArray [][]string, cityList map[string]*CityNode) {

	for rowIndex, row := range citiesArray {

		for colIndex, city := range row {

			//Potentially need a newCity
			var newCity *CityNode

			if city == "" {
				break
			} else {
				//check if city is already in the arraylist if not, createa new node and add it to the list
				if _, exists := cityList[city]; !exists {
					//City does not exist
					//create new node and add it to the list
					newCity = &CityNode{
						City: strings.TrimSpace(city),
					}
					cityList[city] = newCity
				} else {
					newCity = cityList[city]
				}
			}

			if colIndex == 0 {
				continue
			} else {
				//add the city to the Connections city struct
				//originCity.Connections = append(originCity,Connection, currentCity)
				cityList[citiesArray[rowIndex][0]].Connections = append(cityList[citiesArray[rowIndex][0]].Connections, newCity)
			}

		}
	}

	//	for _, node := range cityList {
	//
	//		fmt.Println("City ", node.City, " Connections ", node.Connections)
	//
	//	}

}
