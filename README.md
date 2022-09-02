## protocol buffers 3

Protocol Buffers (protobuf) is a fundamental data serialization format. It is leveraged by many top tech companies such as Google and enables micro-services to transfer data in a format that is safe and efficient. 

All the info regarding the protocol buffers and Go such as:

* how to install protoc 
* how to generate Go code from a `.proto` file
* and all the `.proto` file structure

can be found [here](https://developers.google.com/protocol-buffers/docs/reference/go-generated).

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
```
<sub>For enums we generally keep the 0 tag for the first element of the enum which is unspecified.</sub>