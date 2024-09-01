package main 

import(
	"fmt"
)

//compare if the two maps have the same elements
func compareMaps(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		value2, exist := map2[key]
		if !exist || value1 != value2 {
			return false
		}
	}
	return true
}

func main() {
    map1 := map[string]int{"a": 1, "b": 2, "c": 3}
    map2 := map[string]int{"a": 1, "b": 2, "c": 3}
    map3 := map[string]int{"a": 1, "b": 2, "c": 4}

    fmt.Println(compareMaps(map1, map2)) 
    fmt.Println(compareMaps(map1, map3)) 
}