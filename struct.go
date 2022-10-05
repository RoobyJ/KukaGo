package main

type dataTree struct {
	types     [3]datatype
	startTime int64
	endTime   int64
}

// This is representing each data type in $config.dat file
// for example Bases, Tools, Loads, all of them has a list of names and list of
// rows with data values
type datatype struct {
	name       string
	typeNames  []string
	typeValues [][]float64
}
