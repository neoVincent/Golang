package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func main() {
	dataTransform()
}

// data transform between map and entity
func dataTransform() {
	type Entity struct {
		Key    string
		Values []interface{}
	}

	// from map to entity
	mapObj1 := make(map[string]interface{})
	mapObj1["Key"] = "hero"
	mapObj1["Values"] = []string{"batman", "superman"}
	myEntity1 := Entity{}
	mapstructure.Decode(mapObj1, &myEntity1)
	fmt.Println(myEntity1)

	// from entity to map
	mapObj2 := make(map[string]interface{})
	mapstructure.Decode(myEntity1, &mapObj2)
	fmt.Println(mapObj2)
}
