package main

import (
	"fmt"
)

type Location struct {
	x int
	y int
}

func main() {
	fmt.Println("lalalalala")
}

func findShortestPath(source, dest Location, method string) []Location {
	switch method {
	case "car":
		return findShortestPathByCar(source, dest)
	case "walking":
		return findShortestPathByWalking(source, dest)
	case "public":
		return findShortestPathByPublicTransportion(source, dest)
	case "cycling":
		return findShortestPathByCycling(source, dest)
	// default:
	// 	return []Location{}
	}
	return []Location{}
}

// v1
func findShortestPathByCar(source, dest Location) []Location{
	return []Location{}
}

// v2
func findShortestPathByWalking(source, dest Location) []Location{
	return []Location{}
}

// v3
func findShortestPathByPublicTransportion(source, dest Location) []Location{
	return []Location{}
}

// v4
func findShortestPathByCycling(source, dest Location) []Location{
	return []Location{}
}

type Map struct {
	TravelStrategy
}

func (m * Map) setTravelStrategy(travel TravelStrategy) {
	m.TravelStrategy = travel
}

type TravelStrategy interface {
	findShortestPath(source, dest Location) []Location
}