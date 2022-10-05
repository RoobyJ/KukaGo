package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Init func passing names of structs datatype into their place
func newDataTree() *dataTree {
	return &dataTree{types: [3]datatype{{name: "bases", typeNames: nil, typeValues: nil},
		{name: "tools", typeNames: nil, typeValues: nil},
		{name: "loads", typeNames: nil, typeValues: nil}}}

}

// the 'main' func of class dataTree, the $config.dat is opened and proceed, and then
// it's closed
func (data *dataTree) getConfigData(path string, fileName os.DirEntry) {
	file, err := os.Open(path + "\\" + fileName.Name() + "\\KRC\\R1\\System\\$config.dat")
	resp := os.IsNotExist(err) // check if the given path has this file
	defer func() {
		if err = file.Close(); err != nil && resp == false {
			log.Fatal(err)
		}
	}()
	b, err := io.ReadAll(file)

	data.processData(string(b))
	if b != nil {
		data.createExcel(path, fileName.Name())
		fmt.Printf("%s: %d \n", file.Name(), data.endTime-data.startTime)
	}
}

// splitting the data into rows and checking the prefix of each which length is 9 or more
func (data *dataTree) processData(dataString string) {
	for _, line := range strings.Split(dataString, "\n") {
		if len(line) >= 9 {
			switch line[:9] {
			case "BASE_DATA":
				data.addData(line, 0, false)
			case "TOOL_DATA":
				data.addData(line, 1, false)
			case "LOAD_DATA":
				data.addData(line, 2, false)
			case "BASE_NAME":
				data.addData(line, 0, true)
			case "TOOL_NAME":
				data.addData(line, 1, true)
			case "LOAD_NAME":
				data.addData(line, 2, true)
			}
		}

	}

}

// this function replaces specified chars and return a string to the next process
func replaceChars(s string) string {
	replacer := strings.NewReplacer("X", "", "Y", "", "Z", "", "A", "", "B", "", "C", "", "M", "", "{", "",
		"}", "", "M", "", "J", "")
	return replacer.Replace(s)
}

// adding names and values to their destination in object dataTree with changing for values the string to float
func (data *dataTree) addData(s string, idType int, typeOfData bool) {
	if typeOfData {
		data.types[idType].typeNames = append(data.types[idType].typeNames, s[strings.Index(s, string('"'))+1:len(s)-2]) // adding the name

	} else {
		// adding list of floats to data structure
		data.types[idType].typeValues = append(data.types[idType].typeValues, clearNumbers(s))
	}
}

// function for values processing, separating the values by comma then putting them into a slice
// and returning a slice of float64
func clearNumbers(s string) []float64 {
	var numbers []float64
	// convert s to float and check if all fine
	for _, n := range strings.Split(strings.Replace(s[strings.Index(s, "=")+2:len(s)-3], " ", "", -1), ",") {
		if newN, err := strconv.ParseFloat(replaceChars(n), 32); err == nil {
			numbers = append(numbers, newN)
		} else {
			log.Fatal(err)
		}
	}
	return numbers
}
