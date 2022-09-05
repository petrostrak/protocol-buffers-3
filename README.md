## protocol buffers 3

Protocol Buffers (protobuf) is a fundamental data serialization format. It is leveraged by many top tech companies such as Google and enables micro-services to transfer data in a format that is safe and efficient. 

### Setup protoc
All the info regarding the protocol buffers and Go such as:

* how to install protoc 
* how to generate Go code from a `.proto` file
* and all the `.proto` file structure

can be found [here](https://developers.google.com/protocol-buffers/docs/reference/go-generated).

Find the correct protocol buffers version based on your Linux [distro](https://github.com/protocolbuffers/protobuf/releases).
Example with x64:
```
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/
# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
# Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
```

### Protobuf Scalar Types
There are 15 different scalar types we can use in protocol buffers. 

| Type | Keyword | Default value |
| :---        |    :----:   |          ---: |
|Number|int32,int64,sint32,sint64,uint32,uint64,fixed32,fixed64,sfixed32,sfixed64,float,double|0|
|Boolean|bool|false|
|String|string|empty string|
|Bytes|bytes|empty byte slice|
|Enums|enum|first value of the enum|

<sub>Strings only accept UTF-8 encoded or 7-bit ASCII</sub>

### Enums
Specifically for `Enums` the tag number starts at 0.
```
enum Weekdays {
    WEEKDAY_UNSPECIFIED = 0;
    WEEKDAY_MONDAY = 1;
    WEEKDAY_TUESDAY = 2;
    WEEKDAY_WEDNESDAY = 3;
    WEEKDAY_THURSDAY = 4;
    WEEKDAY_FRIDAY = 5;
    WEEKDAY_SATURDAY = 6;
    WEEKDAY_SUNDAY = 7;
}

message Schedule {
    WeekDays daysOfWeek = 1;
    // ...
}
```
<sub>For enums we generally keep the 0 tag for the first element of the enum which is unspecified.</sub>

### Repeated Fields
`repeated` fields represent lists

```
message Account {
    uint32 id = 1;
    Person person = 2;

    // repeated fields represent lists
    // and is equivalent of var phones []string
    repeated string phones = 3;
}
```

### Nested Messages
We can nest our messages in a `.proto` file as shown below:

```
message City {
    string name = 1;
    uint64 zip_code = 2;
    string country_name = 3;
    
    message Street {
        string street_name = 1;
        City city = 2;

        message Building {
            string building_name = 1;
            uint32 building_number = 2;
            Street street = 3;
        }
    }
}

message Address {
    City city = 1;
    City.Street street = 2;
    City.Street.Building building = 3;
}
```

### Import Messages
We can import messages from different `.proto` files:

```
import "city.proto";
import "street.proto";
import "building.proto";

message Address {
    City city = 1;
    Street street = 2;
    Building building = 3;
}
```

### Import Messages from package
Or we can import messages from different packages:

```
package city;

message City {
    string name = 1;
    uint64 zip_code = 2;
    string country_name = 3;
}
```

```
package street;

import "city.proto";

message Street {
    string street_name = 1;
    city.City city = 2;
}
```

```
package building;

import "street.proto";

message Building {
    string building_name = 1;
    uint32 building_number = 2;
    street.Street street = 3;
}
```

```
package address;

import "city.proto";
import "street.proto";
import "building.proto";

message Address {
    city.City city = 1;
    street.Street street = 2;
    building.Building building = 3;
}
```

## Golang Programming with protobuf
### A simple proto struct example
```
# .proto
message Simple {
    uint32 id = 1;
    bool is_simple = 2;
    string name = 3;
    repeated int32 sample_list = 4;
}

# .go
func getSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "Petros Trak",
		SampleList: []int32{1, 2, 3, 4, 5, 6},
	}
}
```

### A complex proto struct example
```
# .proto
message Dummy {
    int32 id = 1;
    string name = 2;
}

message Complex {
    Dummy one_dummy = 1;
    repeated Dummy multi_dummy = 2;
}

# .go
func getComplex() *pb.Complex {
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
```

### An enum proto struct example
```
# .proto
enum EyeColor {
    EYE_COLOR_UNSPECIFIED = 0;
    EYE_COLOR_GREEN = 1;
    EYE_COLOR_BLUE = 2;
    EYE_COLOR_BROWN = 3;
}

message Enumeration {
    EyeColor eye_color = 1;
}

# .go
func getEnum() *pb.Enumeration {
	return &pb.Enumeration{
		// EyeColor: pb.EyeColor_EYE_COLOR_BROWN,
		EyeColor: 3,
	}
}
```

### OneOfs
```
# .proto

// oneof either takes a message that is string, 
// or a message that is int32.
message Result {
    oneof result {
        string message = 1;
        uint32 id = 2;
    }
}

# .go
func getOneOf(msg any) {
	switch x := msg.(type) {
	case *pb.Result_Id:
		fmt.Println(msg.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(msg.(*pb.Result_Message).Message)
	default:
		fmt.Printf("msg has unexpected type %v\n", x)
	}
}
```

### Maps
```
# .proto

// IdWrapper is the value of the key-value pair in a map.
// So the type of the wrapper will be the type of the value
// of the map.
message IdWrapper {
    uint32 id = 1;
}

// MapExample is a map[string]uint32
message MapExample {
    map<string, IdWrapper> ids = 1;
}

# .go
func getMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"key1": {Id: 1},
			"key2": {Id: 2},
			"key3": {Id: 4},
			"key4": {Id: 5},
		},
	}
}
```

### Reading and Writing to Disk
#### Write to file with proto
```
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
}
```

#### Read from file with proto
```
func readFromFile(filename string, pb proto.Message) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("cannot read file", err)
		return
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("cannot deserialize data", err)
		return
	}
}
```

### Reading and Writing to JSON
#### Encode to JSON
```
func toJSON(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)
	if err != nil {
		log.Fatalln("cannot encode to JSON", err)
		return ""
	}
	return string(out)
}
```

#### Decode from JSON
```
func fromJSON(in string, pb proto.Message) {
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err := option.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("cannot decode JSON to pb", err)
		return
	}
}

func pbfromJSON(json string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(json, message)
	return message
}

{
    // Example, decode JSON to pb.Simple struct.
    msg := pbfromJSON(json, reflect.TypeOf(pb.Simple{}))
    // ...
}
```

### Data Evolution with Protobuf
The most important rule when it comes to updating a `.proto` file is **not to change the number tags** of the existing fields. 
A more detailed guide about updating a message type can be found [here](https://developers.google.com/protocol-buffers/docs/proto3?hl=en#updating).

#### Renaming Fields
We can rename fields freely when we want to update a `.proto` file. Remember that the only thing that matters when serializing or deserializing a proto message is the tags, not the name.

#### Removing Fields
If we want to remove an existing field for our newer `.proto` version, we use the keyword `reserved` before the number tag and optionally before the name. In this way, we prevent the tag and name for future use.
```
message Example {
	reserved 2,3,9 to 11; // 9 and 11 including
	reserved "first_name", "last_name";
	uint32 id = 1;
}
```
### --decode_raw option
The `--decode_raw` command reads an arbitrary protocol message from standard input and writes the raw tag/value pairs in text format to standard output.
```
# .proto
message Simple {
    uint32 id = 1;
    bool is_simple = 2;
    string name = 3;
    repeated int32 sample_list = 4;
}

# .bin
*My name"

$ cat simple.bin | protoc --decode_raw
1: 42
2: 1
3: "My name"
4: "\001\002\003\004\005\006"
```

### --decode option
The `--decode` command reads a binary message of the given type from standard input and writes it in text format to standard output.  The message type must be defined in PROTO_FILES or their imports.. 
```
# .proto
message Simple {
    uint32 id = 1;
    bool is_simple = 2;
    string name = 3;
    repeated int32 sample_list = 4;
}

# .bin
*My name"

$ cat simple.bin | protoc --decode=Simple simple.proto
id: 42
is_simple: true
name: "My name"
sample_list: 1
sample_list: 2
sample_list: 3
sample_list: 4
sample_list: 5
sample_list: 6
```
<sub>In case the `.proto` file is defined in a package e.g. simple, then the --decode parameter changes. E.g. `cat simple.bin | protoc --decode=simple.Simple simple.proto`</sub>

### --encode
The `--encode` command reads a text-format message of the given type from standard input and writes it in binary to standard output.
```
$ cat simple.bin | protoc --decode=Simple simple.proto > simple.txt
id: 42
is_simple: true
name: "My name"
sample_list: 1
sample_list: 2
sample_list: 3
sample_list: 4
sample_list: 5
sample_list: 6

$ cat simple.bin | protoc --decode=Simple simple.proto > simple.txt
$ cat simple.txt | protoc --encode=Simple simple.proto > simple.pb

# simple.pb
*My name"

# To check that the encoded message is the same with the source
$ diff simple.bin simple.pb
```
