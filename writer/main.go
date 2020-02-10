package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	mycalc "github.com/igorskh/flatbuffers-zmq-tutorial/MyCalc"
	"gopkg.in/zeromq/goczmq.v4"
)

func writeBytes(buf []byte, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.Write(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d byte(s) to file\n", n)
}

func createObject(builder *flatbuffers.Builder, n int, axis int, action string, desc string) []byte {
	builder.Reset()

	description := builder.CreateString(desc)

	mycalc.RawDataStartValuesVector(builder, n)
	for i := n; i >= 0; i-- {
		val := float32(i)
		mycalc.CreateVec3(builder, val, val+1, val+2)
	}
	values := builder.EndVector(n)

	mycalc.RawDataStart(builder)
	mycalc.RawDataAddDescription(builder, description)
	mycalc.RawDataAddValues(builder, values)
	mycalc.RawDataAddAxis(builder, int16(axis))
	switch action {
	case "sum":
		mycalc.RawDataAddCalcType(builder, mycalc.CalcTypeSum)
	case "median":
		mycalc.RawDataAddCalcType(builder, mycalc.CalcTypeMedian)
	default:
		mycalc.RawDataAddCalcType(builder, mycalc.CalcTypeAverage)
	}
	calcReq := mycalc.RawDataEnd(builder)

	builder.Finish(calcReq)
	buf := builder.FinishedBytes()
	return buf
}

func zeroMQDealer(dat []byte) {
	dealer, err := goczmq.NewDealer("tcp://127.0.0.1:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Destroy()
	dealer.SetOption(goczmq.SockSetSndtimeo(100))
	log.Println("dealer created and connected")

	err = dealer.SendFrame(dat, goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("dealer sent message")

	time.Sleep(1 * time.Millisecond)
}

func main() {
	var n = flag.Int("n", 5, "number of elements")
	var axis = flag.Int("a", 0, "axis 0 - X, 1 - Y, 2 - Z")
	var toFile = flag.Bool("f", false, "save to file")
	var path = flag.String("o", "data/data.dat", "path to save")
	var desc = flag.String("d", "do something", "description")
	var action = flag.String("t", "mean", "type of action (sum, mean, median)")
	var bufSize = flag.Int("s", 1024, "buffer size")
	flag.Parse()

	builder := flatbuffers.NewBuilder(*bufSize)

	buf := createObject(builder, *n, *axis, *action, *desc)

	if *toFile {
		writeBytes(buf, *path)
		return
	}
	zeroMQDealer(buf)
}
