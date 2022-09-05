// json.go has the methods required to encode a proto.Message to
// JSON and decode JSON to a proto.Message struct.
package main

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func encodeToJSON(pb proto.Message) string {
	options := protojson.MarshalOptions{
		Multiline: true,
	}

	out, err := options.Marshal(pb)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(out)
}

func decodeFromJSON(in string, pb proto.Message) {
	options := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	err := options.Unmarshal([]byte(in), pb)
	if err != nil {
		fmt.Println(err)
		return
	}
}
