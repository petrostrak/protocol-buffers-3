package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("cannot encode to JSON", err)
		return ""
	}

	return string(out)
}

func fromJSON(in string, pb proto.Message) {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("cannot decode from JSON", err)
		return
	}
}
