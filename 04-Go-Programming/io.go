package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func writeToFile(filename string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("cannot serialize to bytes", err)
		return
	}

	err = ioutil.WriteFile(filename, out, 0644)
	if err != nil {
		log.Fatalln("cannot write to file", err)
		return
	}

	fmt.Println("Data has been written!")
}
