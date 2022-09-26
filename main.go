package main

import (
	"flag"
	"log"
	"os"
	"sync"
)

// Variables used for command line parameters
var (
	path string
)

// parsing input into var path
func init() {
	flag.StringVar(&path, "p", "", "backups path")
	flag.Parse()
}

// executing the main process and sending done on the end of the task
func configToExcel(wg *sync.WaitGroup, f os.DirEntry) {
	defer wg.Done()
	obj := *newDataTree()
	obj.getConfigData(path, f)

}

// counting all folders in path
func countDirs(f []os.DirEntry) int {
	ctn := 0
	for _, file := range f {
		if file.IsDir() {
			ctn += 1
		}
	}
	return ctn
}

// checking if an error has happened
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	wg := sync.WaitGroup{}
	files, err := os.ReadDir(path) // all folders in entered dir
	checkError(err)
	count := countDirs(files) // amount of folders in input dir
	wg.Add(count)
	for _, f := range files {
		if f.IsDir() {
			go configToExcel(&wg, f)
		}
	}
	wg.Wait()
}
