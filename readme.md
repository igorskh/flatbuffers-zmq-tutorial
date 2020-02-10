# FlatBuffers + ZeroMQ Sample Sample

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
sudo apt apt-get install libzmq3-dev
```

## Create Go project
For this repository the following path is used - `$GOPATH/src/github.com/igorskh/flatbuffers-zmq-tutorial`. Generally any path in `$GOPATH/src` is fine.
```bash
cd $GOPATH/src
```

In the project folder create two subfolders for writer and reader:
```bash
mkdir reader
mkdir wirte
```