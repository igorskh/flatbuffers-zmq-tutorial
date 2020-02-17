# FlatBuffers + ZeroMQ Sample

This repository is based on [FlatBuffers official Tutorial](https://google.github.io/flatbuffers/flatbuffers_guide_tutorial.html).

[Second part and implementation in C++](https://github.com/igorskh/flatbuffers-zmq-tutorial-cpp).

## Scenario

This example performs the following:
* The writer application sends an object with array of `values` each `value` contains 3 fields `X`, `Y` and `Z`. 
* The wirter specifies on which axis the calculation must be performed, 0-2 for X, Y, Z correspondingly.
* The wirter specifies requested action `mean`, `median`, `sum`.
* The reader receives the request, performs request calculation and outputs result to `STDOUT`. 

The [schema for the scenario](schema.fbs) is the following:
```cpp
namespace MyCalc;

enum CalcType: byte { Sum = 0, Average, Median = 2 }

struct Vec3 {
    x: float;
    y: float;
    z: float;
}

table RawData {
    description: string;
    axis: short = 0;
    calcType: CalcType = Sum; 
    values: [Vec3];
}

root_type RawData;
```

`Description` is just an arbitrary string value, does not affect calculations.

## Installation
### FlatBuffers
[Build from source](https://google.github.io/flatbuffers/flatbuffers_guide_building.html).

And [another tutorial](https://rwinslow.com/posts/how-to-install-flatbuffers/).
```bash
git clone https://github.com/google/flatbuffers.git
cd flatbuffers
mkdir build
cd build
cmake -G "Unix Makefiles" -DCMAKE_BUILD_TYPE=Release
make
./flattests
make install
```

For using with go:
```bash
go get github.com/google/flatbuffers/go
```

### ZeroMQ
Follow the [instructions](https://zeromq.org/download/#linux).

For Ubuntu:
```bash
sudo ap install libzmq3-dev libczmq4 libczmq-dev
go get gopkg.in/zeromq/goczmq.v4
```

Go tutorial for ZMQ: [https://zeromq.org/languages/go/](https://zeromq.org/languages/go/).

## Usage 
### Get the code
Clone the repository to your `GOPATH` to `$GOPATH/src/github.com/igorskh/flatbuffers-zmq-tutorial`
```bash
mkdir -p $GOPATH/src/github.com/igorskh
cd $GOPATH/src/github.com/igorskh
git clone https://github.com/igorskh/flatbuffers-zmq-tutorial.git
```

### Scripts description

Sample script work in two modes:
* Write/Read to/from the file
* Push/Poll to/from the ZeroMQ 

Writer/pusher code is located in `writer` folder, reader/poller in `reader`. Binaries can be compiled with the following code:
```bash
go build -o ./bin/writer ./writer
go build -o ./bin/reader ./reader
```

Argument  | Description | Writer | Reader | Type | Default | Options
--- | --- | --- | --- | ---  | --- | ---
-f | Write/Read to/from file | v | v | bool | false | {true, false}
-i | Path to input file, required only if `-f` is `true` |  | v | string | `data/data.dat` | 
-c | Continuous read from socket, if false reads only the first message  |  | v | bool | false | {true, false}
-o | Path to output file, required only if `-f` is `true` | v |  | string | `data/data.dat` | 
-n | Amount of values to generate | v |  | integer | 5 | 
-a | Axis for calculation over which the calculation is requested, 0 - X, 1 - Y, 2 - Z | v |  | integer | 5 | {0,1,2}
-d | Description field | v |  | string | `do something` | 
-t | Calculation type to request | v |  | string | `mean` |  {mean, median, sum}
-s | Builder buffer size for FlatBuffers | v |  | integer | 1024 |  

### File Read/Write
Write to file (default path: `data/data.dat`):
```bash
go run writer/main.go -f=true
# Wrote 128 byte(s) to file
```

Read from file:
```bash
go run reader/main.go -f=true
```
Result should be the following:
```text
Got request:  do something
Calc axis:  0
Values:  20
[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
Calculating Mean
Result:  9.5
```

### ZeroMQ Sockets
```bash
# start router:
go run reader/main.go
# 2020/02/10 14:32:06 router created and bound
```

Run dealer:
```bash
go run writer/main.go
# 2020/02/10 14:35:43 dealer created and connected
# 2020/02/10 14:35:43 dealer sent message
```

Router outputs:
```text
2020/02/10 14:36:11 router received from '[0 132 127 38 95]' 
Got request:  do something
Calc axis:  0
Values:  5
[0 1 2 3 4]
Calculating Mean
Result:  2
```

## Develop
For this repository the following path is used - `$GOPATH/src/github.com/igorskh/flatbuffers-zmq-tutorial`. Generally any path in `$GOPATH/src` is fine, just make sure the path to generated flatbuffer models are correct.

If you change the schema, generate Go code:
```bash
flatc --go schema.fbs
```
