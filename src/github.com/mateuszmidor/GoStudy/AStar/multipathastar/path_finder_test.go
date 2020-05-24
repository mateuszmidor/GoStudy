package multipathastar

import (
	"strings"
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

/*
cost not used in multipath algo as all paths are to be found
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
*/

// END PROBLEM SPECIFIC CODE

// test helpful tools
const pathSeparator = ", "

func pathToString(path Path, cities Cities) string {
	if len(path) == 0 {
		return "<empty path>"
	}

	result := cities[path[0]].Name
	for i := 1; i < len(path); i++ {
		result += "-" + cities[path[i]].Name
	}

	return result
}

func pathsToString(paths []Path, cities Cities) string {
	if len(paths) == 0 {
		return ""
	}

	var result string
	for _, p := range paths {
		result += pathToString(p, cities) + pathSeparator
	}

	return result
}

func cutOutSubstring(src, sub string) string {
	i := strings.Index(src, sub)
	if i == -1 {
		return src
	}

	return src[:i] + src[i+len(sub):]
}

func checkExpectedPaths(expected []string, actual []Path, cities Cities, t *testing.T) {
	pathsString := pathsToString(actual, cities)
	for _, p := range expected {
		i := strings.Index(pathsString, p)
		if i == -1 {
			t.Errorf("Missing path: %s. Found: %s", p, pathsString)
		} else {
			pathsString = cutOutSubstring(pathsString, p+pathSeparator)
		}
	}
	if len(pathsString) > 0 {
		t.Errorf("Unexpected paths: %s", pathsString)
	}
}

/*
The general direction is: GDA -> KRK
Loops:
-RZE-WAW-KRK-RZE
-GDA-WAW-BYK-GDA
                          +-----+
                          | GDA |<------
                       /--+-----+\      \------------
                 /-----    /      \                  \-----
+-----+    /-----         /        \                    +-----+
| SZC | <--              v          \                   | BYK |
+-----+              +-----+         |                  +-----+
   \                 | BDG |         \                    -> ^
    \                +-----+          \               ---/  /
     |                  /              \          ---/      |
     \                /-                \     ---/         /
      \              /                   v  -/             |
       \           /-                 +-----+             /
        |         /                   | WAW |             |
        \        /                    +-----+<           /
         \     /-                        /    \          |
          v   <                         /      \        /
        +-----+                        |        \-      |
        | WRO |                        /          \    /
        +-----+--\                    /            \   |
                  ---\               /              \ /
                      ---\          v             +-----+
                          -->+-----+       ------>| RZE |
                             | KRK |------/       +-----+
                             +-----+
*/
type Map struct {
	cities                                      Cities
	neighbors                                   *Neighbors
	costs                                       Costs
	lub, gda, szc, bdg, byk, waw, wro, rze, krk NodeID
}

func makeFlightMap() Map {
	cities := Cities{}
	krk := AppendCity(&cities, City{"KRK", 0, 0})    // 0
	rze := AppendCity(&cities, City{"RZE", 20, 5})   // 1
	wro := AppendCity(&cities, City{"WRO", -20, 10}) // 2
	waw := AppendCity(&cities, City{"WAW", 10, 20})  // 3
	byk := AppendCity(&cities, City{"BYK", 30, 35})  // 4
	bdg := AppendCity(&cities, City{"BDG", -10, 30}) // 5
	szc := AppendCity(&cities, City{"SZC", -30, 30}) // 6
	gda := AppendCity(&cities, City{"GDA", 0, 40})   // 7
	lub := AppendCity(&cities, City{"LUB", 25, 10})  // 8, no connections
	costs := Costs(nil)                              // costs not used in multipath algo //NewCityTravelCosts(cities)

	neighbors := NewNeighbors()
	neighbors.Connect(gda, szc)
	neighbors.Connect(szc, wro)
	neighbors.Connect(gda, bdg)
	neighbors.Connect(bdg, wro)
	neighbors.Connect(wro, krk)
	neighbors.Connect(gda, waw)
	neighbors.Connect(waw, krk)
	neighbors.Connect(krk, rze)
	neighbors.Connect(rze, waw)
	neighbors.Connect(waw, byk)
	neighbors.Connect(byk, gda)
	neighbors.Connect(rze, byk)

	return Map{
		cities:    cities,
		neighbors: neighbors,
		costs:     costs,
		lub:       lub, gda: gda, szc: szc, bdg: bdg, byk: byk, waw: waw, wro: wro, rze: rze, krk: krk,
	}
}

func TestShouldFindSingleSegmentPath(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.wro, m.krk, m.neighbors, m.costs)

	// then
	expected := []string{
		"WRO-KRK",
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}

func TestShouldFindTwoSegmentsPath(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.szc, m.krk, m.neighbors, m.costs)

	// then
	expected := []string{
		"SZC-WRO-KRK",
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}

func TestShouldFindTwoSegments2Paths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.gda, m.wro, m.neighbors, m.costs)

	// then
	expected := []string{
		"GDA-BDG-WRO",
		"GDA-SZC-WRO",
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}

func TestShouldFindDirectAndIndirectPath(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.rze, m.waw, m.neighbors, m.costs)

	// then
	expected := []string{
		"RZE-WAW",
		"RZE-BYK-GDA-WAW",
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}

func TestShouldFindIfSplittingPaths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.gda, m.krk, m.neighbors, m.costs)

	// then
	expected := []string{
		"GDA-WAW-KRK",
		"GDA-BDG-WRO-KRK",
		"GDA-SZC-WRO-KRK",
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}

func TestShouldFindIfMergingPaths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.byk, m.wro, m.neighbors, m.costs)

	// then
	expected := []string{
		"BYK-GDA-SZC-WRO",
		"BYK-GDA-BDG-WRO",
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}

func TestShouldFindMultiLongPaths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.gda, m.byk, m.neighbors, m.costs)

	// then
	expected := []string{
		"GDA-WAW-BYK",
		"GDA-WAW-KRK-RZE-BYK",
		"GDA-SZC-WRO-KRK-RZE-BYK",
		"GDA-BDG-WRO-KRK-RZE-BYK",
		"GDA-SZC-WRO-KRK-RZE-WAW-BYK",
		"GDA-BDG-WRO-KRK-RZE-WAW-BYK",
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}

func BenchmarkFindMultiLongPaths(b *testing.B) {
	m := makeFlightMap()
	for n := 0; n < b.N; n++ {
		_ = findPaths(m.gda, m.byk, m.neighbors, m.costs)
	}
}

func TestShouldHandleCycle(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.waw, m.szc, m.neighbors, m.costs)

	// then
	expected := []string{
		"WAW-BYK-GDA-SZC",
		"WAW-KRK-RZE-BYK-GDA-SZC",
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}

func TestShouldHandleNoSuchPath(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := findPaths(m.waw, m.lub, m.neighbors, m.costs)

	// then
	expected := []string{
		// no such connection
	}
	checkExpectedPaths(expected, paths, m.cities, t)
}
