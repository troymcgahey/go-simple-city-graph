package graph

import "strings"

//"fmt"

func CreateGraph(citiesArray [][]string, cityList map[string]*CityNode) {

	for rowIndex, row := range citiesArray {

		for colIndex, city := range row {

			if city == "" {
				break
			} else {
				//check if city is already in the arraylist if not, createa new node and add it to the list
				if _, exists := cityList[city]; !exists {
					//create new node and add it to the list
					cityList[city] = &CityNode{
						City: strings.TrimSpace(city),
					}
				}
			}

			if colIndex == 0 {
				continue
			} else {
				//add the city to the Connections city struct
				//originCity.Connections = append(originCity,Connection, currentCity)
				cityList[citiesArray[rowIndex][0]].Connections = append(cityList[citiesArray[rowIndex][0]].Connections, cityList[city])
			}

		}
	}

}
