package graph

import "fmt"

func SearchGraphBFS(cityGraph map[string]*CityNode, startCity string, destinationCity string) bool {

	type connectionCityNode struct {
		cityNode      *CityNode
		connectedFrom *CityNode
	}

	var newConnectionCity *connectionCityNode

	newConnectionCity = &connectionCityNode{
		cityNode: cityGraph[startCity],
	}

	searchCityList := []*connectionCityNode{newConnectionCity}

	//for lineage, need to a reference for each CityNode, where it was connected from

	for len(searchCityList) > 0 {

		//set the currentCity to the first element in the searchCityList Slice
		currentCity := searchCityList[0]

		//Remove the currentCity from the searchCityList
		searchCityList = searchCityList[1:]

		if currentCity.cityNode.City == "" {
			continue
		}

		if currentCity.cityNode.City == destinationCity {

			fmt.Println("Connected from ", currentCity.connectedFrom.City)
			return true
		}

		for i, _ := range currentCity.cityNode.Connections {
			//for len(currentCity.cityNode.Connections) > 0 {

			newConnectionCity = &connectionCityNode{
				cityNode:      currentCity.cityNode.Connections[i],
				connectedFrom: currentCity.cityNode,
			}

			searchCityList = append(searchCityList, newConnectionCity)

			i++
		}
	}

	return false
}
