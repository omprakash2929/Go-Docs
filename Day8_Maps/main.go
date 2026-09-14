package main

import (
	"fmt"
	"maps"
)

// * Maps -> it's like objects, dict
func main() {

	//? creating maps

	// m := make(map[string]string)

	// //? setting an elmenet
	// m["name"] = "golang"
	// m["area"] = "backend"
	// m["ext"] = ".go"

	// get element
	// fmt.Println(m["name"])
	// fmt.Println(m["name11"])  //! If key does not exists in the map then return zero value

	// print maps
	// fmt.Println(m)

	// m := make(map[string]int)

	// m["age"] = 21
	// m["id"] = 01
	// m["price"] = 500

	// fmt.Println(m["phon"])
	// fmt.Println(len(m))
	// delete(m, "price") //? delete the element "maps name": key of element
	// clear(m) //? clear the maps
	// fmt.Println(m)

	//! without useing make keyword

	// m := map[string]int{"id": 01, "price": 400, "phone": 3}
	// fmt.Println(m)

	// s := map[string]string{}

	// s["name"] = "om"

	// fmt.Println(s)

	// m := map[string]int{"id": 01, "price": 400, "phone": 3}

	// k, ok := m["price"] //? If you not want to k then use _ "underscore"
	// fmt.Println(k)      //? K give the value of key
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("Not ok")
	// }

	m1 := map[string]int{"id": 01, "price": 400, "phone": 3}
	m2 := map[string]int{"id": 01, "price": 400, "phone": 3}

	fmt.Println(maps.Equal(m1, m2))

	// m3 := maps.Clone(m2) //? It's shallow copy
	// maps.Copy(dst,src) //? use to copy

}
