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

func doComplex() *pb.Complex {
	return &pb.Complex{
		OnDummy: &pb.Dummy{
			Id:   1,
			Name: "Petros Trak",
		},
		MultiDummy: []*pb.Dummy{
			{Id: 2, Name: "Eirini Tour"},
			{Id: 3, Name: "Deppy Bou"},
			{Id: 4, Name: "Giannis Lio"},
		},
	}
}

func main() {
	// fmt.Println(doSimple())
	fmt.Println(doComplex())
}
