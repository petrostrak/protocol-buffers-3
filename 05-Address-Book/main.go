package main

import (
	pb "address-book/proto"
	"fmt"
	"reflect"
)

var (
	addressBook []*pb.Person
)

// Implement a way to read and write Person in the Address Book.
func main() {
	addressBook = append(addressBook, addPerson("petros", "pt@gmail.com", 1))
	addressBook = append(addressBook, addPerson("maggie", "maggie@gmail.com", 2))
	addressBook = append(addressBook, addPerson("eirini", "eirini@gmail.com", 4))
	addressBook = append(addressBook, addPerson("deppy", "deppy@gmail.com", 3))

	// for _, a := range addressBook {
	// 	fmt.Println(a)
	// }

	json := pbToJSON(addressBook[1])
	msg := pbFromJSON(json, reflect.TypeOf(pb.Person{}))
	fmt.Println(json)
	fmt.Println(msg)

	writeToFile("addressBook", addressBook[2])

	var addr pb.Person
	readFromFile("addressBook", &addr)
}
