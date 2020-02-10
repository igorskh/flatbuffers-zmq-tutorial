# FlatBuffers + ZeroMQ Sample

This repository is based on [FlatBuffers official Tutorial](https://google.github.io/flatbuffers/flatbuffers_guide_tutorial.html).

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
sudo apt apt-get install libzmq3-dev libczmq4 libczmq4-dev
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