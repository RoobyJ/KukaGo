package main

type dataTree struct {
	types     [3]datatype
	startTime int64
	endTime   int64
}

type datatype struct {
	name       string
	typeNames  []string
	typeValues [][]float64
}
