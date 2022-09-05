// io.go has the methods required to read and write a
// proto.Message to a file
package main

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

func writeToFile(filename string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(filename, out, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readFromFile(filename string, pb proto.Message) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		fmt.Println(err)
		return
	}
}
