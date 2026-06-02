package graph

import (
	"fmt"
)

func CreateGraph(citiesArray *[][]string, cityList *map[string]*CityNode) {

	for rowIndex, row := range *citiesArray {

		originCity := *citiesArray[row][0]

		for colIndex, city := range row {
			fmt.Printf(
				"row=%d col=%d city=%s\n",
				rowIndex,
				colIndex,
				city,
			)

			if city == "" {
				break
			} else {
				//check if city is already in the arraylist if not, createa new node and add it to the list
				if node, exists := cityList[city]; !exists {
					//create new node
					newNode := &CityNode{
						City:        city,
						Connections: {},
					}
					cityList[city] = newNode
				}
			}

			//if not an origin city, add this city to the origin cities connection list
			if colIndex > 0 {
				originCity.Connections.append(&newNode)
			}
		}
	}

}
