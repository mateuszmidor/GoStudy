package genericastar

import (
	"math"
	"testing"
)

// BEGIN PROBLEM SPECIFIC CODE

type City struct {
	Name string
	X    int16
	Y    int16
}

type Cities = []City

func AppendCity(cities *Cities, node City) NodeID {
	id := NodeID(len(*cities))
	*cities = append(*cities, node)
	return id
}

type CityTravelCosts struct {
	cities Cities
}

func NewCityTravelCosts(cities Cities) *CityTravelCosts {
	return &CityTravelCosts{cities}
}

func distance(from *City, to *City) CostType {
	dx := float64(to.X - from.X)
	dy := float64(to.Y - from.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func (c *CityTravelCosts) D(from NodeID, to NodeID) CostType {
	return distance(&c.cities[from], &c.cities[to])
}

func (c *CityTravelCosts) H(current NodeID, goal NodeID) CostType {
	return distance(&c.cities[current], &c.cities[goal])
}

// END PROBLEM SPECIFIC CODE

// test helper func
func pathToString(path Path, cities Cities) string {
	if len(path) == 0 {
		return "<empty path>"
	}

	result := cities[0].Name
	for i := 1; i < len(path); i++ {
		result += "-" + cities[path[i]].Name
	}

	return result
}

func Test1SegmentPath(t *testing.T) {
	// 1. setup
	var cities Cities
	krk := AppendCity(&cities, City{"KRK", 0, 0})
	waw := AppendCity(&cities, City{"WAW", 0, 25})
	neighbors := NewNeighbors()
	neighbors.Connect(krk, waw)
	costs := NewCityTravelCosts(cities)

	// 2. run
	path := findPath(krk, waw, neighbors, costs)

	// 3. check
	pathString := pathToString(path, cities)
	if pathString != "KRK-WAW" {
		t.Errorf("Invalid path found: %s", pathString)
	}
}

func Test2SegmentPath(t *testing.T) {
	// 1. setup
	var cities Cities
	krk := AppendCity(&cities, City{"KRK", 0, 0})
	waw := AppendCity(&cities, City{"WAW", 0, 25})
	gda := AppendCity(&cities, City{"GDA", 0, 50})
	neighbors := NewNeighbors()
	neighbors.Connect(krk, waw)
	neighbors.Connect(waw, gda)
	costs := NewCityTravelCosts(cities)

	// 2. run
	path := findPath(krk, gda, neighbors, costs)

	// 3. check
	pathString := pathToString(path, cities)
	if pathString != "KRK-WAW-GDA" {
		t.Errorf("Invalid path found: %s", pathString)
	}
}

func Test2SegmentBestPath(t *testing.T) {
	// 1. setup
	var cities Cities
	krk := AppendCity(&cities, City{"KRK", 0, 0})
	waw := AppendCity(&cities, City{"WAW", 0, 25})
	wro := AppendCity(&cities, City{"WRO", -5, 25})
	gda := AppendCity(&cities, City{"GDA", 0, 50})
	neighbors := NewNeighbors()
	neighbors.Connect(krk, waw)
	neighbors.Connect(waw, gda)
	neighbors.Connect(krk, wro)
	neighbors.Connect(wro, gda)
	costs := NewCityTravelCosts(cities)

	// 2. run
	path := findPath(krk, gda, neighbors, costs)

	// 3. check
	pathString := pathToString(path, cities)
	if pathString != "KRK-WAW-GDA" {
		t.Errorf("Invalid path found: %s", pathString)
	}
}
