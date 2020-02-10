package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"

	mycalc "github.com/igorskh/flatbuffers-zmq-tutorial/MyCalc"
)

func sortedValues(dat *mycalc.RawData) []float64 {
	var vec mycalc.Vec3

	values := make([]float64, dat.ValuesLength())
	for i := 0; i < dat.ValuesLength(); i++ {
		ret := dat.Values(&vec, i)
		if !ret {
			fmt.Printf("Error on index = %d\n", i)
			break
		}

		switch axis := dat.Axis(); axis {
		case 1:
			values[i] = float64(vec.Y())
		case 2:
			values[i] = float64(vec.Z())
		default:
			values[i] = float64(vec.X())
		}
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}

func calcMedian(values []float64) float64 {
	medianPos := len(values) / 2

	if medianPos%2 != 0 {
		return values[medianPos]
	}
	return (values[medianPos-1] + values[medianPos]) / 2
}

func calcSum(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum
}

func calcMean(values []float64) float64 {
	sum := calcSum(values)
	return sum / float64(len(values))
}

func readObject(dat []byte) {
	calcReq := mycalc.GetRootAsRawData(dat, 0)
	fmt.Println("Got request: ", string(calcReq.Description()))
	fmt.Println("Calc axis: ", calcReq.Axis())
	fmt.Println("Values: ", calcReq.ValuesLength())

	values := sortedValues(calcReq)
	fmt.Println(values)

	res := 0.0
	switch reqType := calcReq.CalcType(); reqType {
	case mycalc.CalcTypeAverage:
		{
			fmt.Println("Calculating Mean")
			res = calcMean(values)
		}
	case mycalc.CalcTypeSum:
		{
			fmt.Println("Calculating Sum")
			res = calcSum(values)
		}
	case mycalc.CalcTypeMedian:
		{
			fmt.Println("Calculating Median")
			res = calcMedian(values)
		}
	}
	fmt.Println("Result: ", res)
}

func main() {
	var path = flag.String("i", "data/data.dat", "path to save")
	flag.Parse()

	dat, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	readObject(dat)
}
