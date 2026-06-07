package graph

import "fmt"

func SearchGraphBFS(cityGraph map[string]*CityNode, startCity string, destinationCity string) bool {

	type connectionCityNode struct {
		connectionCity string
		connectedFrom  string
	}

	var newConnectionCity connectionCityNode

	newConnectionCity = connectionCityNode{
		connectionCity: startCity,
	}

	searchCityList := []connectionCityNode{newConnectionCity}

	x := 0

	for x < len(searchCityList) {

		//set the currentCity to the first element in the searchCityList Slice
		currentCityNode := searchCityList[x]

		fmt.Println("CurrentCity ", currentCityNode.connectionCity)

		if currentCityNode.connectionCity == "" {
			continue
			x++
		}

		if currentCityNode.connectionCity == destinationCity {

			fmt.Println("Connection Route Found from ", startCity, " to ", destinationCity)

			for currentCityNode.connectedFrom != startCity {

				fmt.Println("Connect from ", currentCityNode.connectionCity)

			}
			return true
		}

		for i := range cityGraph[currentCityNode.connectionCity].Connections {

			newConnectionCity := connectionCityNode{
				connectionCity: cityGraph[currentCityNode.connectionCity].Connections[i].City,
				connectedFrom:  currentCityNode.connectionCity,
			}

			fmt.Println("New Connection City ", newConnectionCity.connectionCity)

			searchCityList = append(searchCityList, newConnectionCity)

		}

		x++
	}

	return false
}
