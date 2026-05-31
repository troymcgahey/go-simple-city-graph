package graph

type CityNode struct {
	City        string
	Connections []*CityNode
}
