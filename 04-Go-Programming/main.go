package main

import (
	"fmt"
	pb "proto-go-programming/proto"
	"reflect"

	"google.golang.org/protobuf/proto"
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

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		// EyeColor: pb.EyeColor_EYE_COLOR_BROWN,
		EyeColor: 3,
	}
}

func doOneOf(msg any) {
	switch x := msg.(type) {
	case *pb.Result_Id:
		fmt.Println(msg.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(msg.(*pb.Result_Message).Message)
	default:
		fmt.Printf("msg has unexpected type %v\n", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"key1": {Id: 1},
			"key2": {Id: 2},
			"key3": {Id: 4},
			"key4": {Id: 5},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	msg := &pb.Simple{}
	readFromFile(path, msg)
	fmt.Println(msg)
}

func pbToJSON(p proto.Message) string {
	return toJSON(p)
}

func pbFromJSON(json string, t reflect.Type) proto.Message {
	msg := reflect.New(t).Interface().(proto.Message)
	fromJSON(json, msg)

	return msg
}

func main() {
	// fmt.Println(doSimple())
	// fmt.Println(doComplex())
	// fmt.Println(doEnum())
	// doOneOf(&pb.Result_Id{Id: 1})
	// doOneOf(&pb.Result_Message{Message: "hello"})
	// fmt.Println(doMap())
	// doFile(doSimple())

	json := pbToJSON(doSimple())
	msg := pbFromJSON(json, reflect.TypeOf(pb.Simple{}))
	fmt.Println(json)
	fmt.Println(msg)

	json = pbToJSON(doComplex())
	msg = pbFromJSON(json, reflect.TypeOf(pb.Complex{}))
	fmt.Println(json)
	fmt.Println(msg)
}
