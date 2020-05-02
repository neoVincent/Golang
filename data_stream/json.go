package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// json is very useful data type during data transition
// for example, in http, rpc.
func main() {
	jsonTransform()
	jsonTransformWithDecoder()
}

// Transform from data structure to json directly
func jsonTransform() {
	fmt.Println("=====================================")
	fmt.Println("jsonTransform")
	fmt.Println("=====================================")
	// to json
	mapStringObj := map[string][]string{"apple": []string{"1", "1"}, "lettuce": []string{"1"}}
	mapStringObjJson, _ := json.Marshal(mapStringObj) // []byte
	fmt.Println(string(mapStringObjJson))

	mapStringInterface := make(map[string]interface{})
	//mapStringInterface["apple"] = make([]interface{},0)
	//mapStringInterface["banana"] = make([]interface{},0)
	mapStringInterface["apple"] = "1"
	mapStringInterface["banana"] = []string{"1", "2"}
	mapStringInterfaceJson, _ := json.Marshal(mapStringInterface)
	fmt.Println(string(mapStringInterfaceJson))

	// from json
	obj1 := make(map[string][]string)
	json.Unmarshal(mapStringObjJson, &obj1) // must use pointer here!
	fmt.Println(obj1)

	obj2 := make(map[string]interface{})
	json.Unmarshal(mapStringInterfaceJson, &obj2)
	fmt.Println(obj2)
}

// json also allow io.writer as encoder
func jsonTransformWithDecoder() {

	fmt.Println("=====================================")
	fmt.Println("jsonTransformWithDecoder")
	fmt.Println("=====================================")

	type KeyValue struct {
		Key   string
		Value string
	}

	var dict = []KeyValue{
		{"apple", "1"},
		{"apple", "1"},
		{"orange", "1"},
	}
	wfile, _ := os.Create("test.dat")
	enc := json.NewEncoder(wfile)

	// write all together
	enc.Encode(&dict)

	// write one by one
	//for _,kv := range dict {
	//	if err := enc.Encode(kv); err!=nil {
	//		log.Print(err.Error())
	//	}
	//}

	rfile, _ := os.OpenFile("test.dat", os.O_RDONLY, 0644)
	var kva []KeyValue
	dec := json.NewDecoder(rfile)

	// decode all together
	dec.Decode(&kva)

	// decode one by one
	//for {
	//	var kv KeyValue
	//	if err := dec.Decode(&kv); err != nil {
	//		break
	//	}
	//	kva = append(kva, kv)
	//}

	for _, kv := range kva {
		fmt.Println(kv)
	}

	os.Remove("test.dat")
}
