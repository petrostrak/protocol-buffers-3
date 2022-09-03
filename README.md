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