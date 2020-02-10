package main

import (
	"flag"
	"fmt"
	"os"

	flatbuffers "github.com/google/flatbuffers/go"
	mycalc "github.com/igorskh/flatbuffers-zmq-tutorial/MyCalc"
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

func fillData(builder *flatbuffers.Builder, n int) {

}

func createObject(builder *flatbuffers.Builder, n int, axis int, desc string) []byte {
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
	mycalc.RawDataAddCalcType(builder, mycalc.CalcTypeSum)
	calcReq := mycalc.RawDataEnd(builder)

	builder.Finish(calcReq)
	buf := builder.FinishedBytes()
	return buf
}

func main() {
	var n = flag.Int("n", 5, "number of elements")
	var axis = flag.Int("a", 0, "axis 0 - X, 1 - Y, 2 - Z")
	var path = flag.String("o", "data/data.dat", "path to save")
	var desc = flag.String("d", "do something", "description")
	var bufSize = flag.Int("s", 1024, "buffer size")
	flag.Parse()

	builder := flatbuffers.NewBuilder(*bufSize)

	buf := createObject(builder, *n, *axis, *desc)
	writeBytes(buf, *path)
}
