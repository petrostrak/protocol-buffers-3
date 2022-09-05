package main

import (
	pb "address-book/proto"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func addPerson(name, email string, id int32) *pb.Person {
	return &pb.Person{
		Name:  name,
		Id:    id,
		Email: email,
		Phones: []*pb.Person_PhoneNumber{
			{Number: "000-1234", Type: pb.Person_HOME},
			{Number: "000-4567", Type: pb.Person_MOBILE},
			{Number: "000-7890", Type: pb.Person_WORK},
		},
		LastUpdated: timestamppb.Now(),
	}
}

func pbToJSON(p proto.Message) string {
	return encodeToJSON(p)
}

func pbFromJSON(json string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	decodeFromJSON(json, message)

	return message
}
