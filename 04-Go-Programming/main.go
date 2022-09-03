package main

import (
	"fmt"
	pb "proto-go-programming/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "Petros Trak",
		SampleList: []int32{1, 2, 3, 4, 5, 6},
	}
}

func main() {
	fmt.Println(doSimple())
}
