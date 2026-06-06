package graph

func SearchGraphBFS(cityList map[string]*CityNode, startCity string, destinationCity string) bool {

	//Initilize the searchCityList to be the startCity
	searchCityList := []*CityNode{cityList[startCity]}

	for len(searchCityList) > 0 {

		//set the currentCity to the first element in the searchCityList Slice
		currentCity := searchCityList[0]

		//Remove the currentCity from the searchCityList
		searchCityList = searchCityList[1:]

		if currentCity.City == "" {
			continue
		}

		if currentCity.City == destinationCity {
			return true
		}

		//Add the currentCities Connections to the searchCityList
		searchCityList = append(searchCityList, currentCity.Connections...)

	}

	return false
}
