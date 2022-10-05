package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"sync"
)

// TODO: read form .dat files nad check tools and bases
//
// Variables used for command line parameters
var (
	path   string
	action string
)

// parsing input into var path
func init() {
	flag.StringVar(&path, "p", "", "backups path")
	flag.StringVar(&action, "a", "", "configToExcel")

	flag.Parse()
}

// executing the main process and sending done on the end of the task
func configToExcel(wg *sync.WaitGroup, f os.DirEntry) {
	defer wg.Done()
	obj := *newDataTree()
	obj.getConfigData(path, f)

}

// In build
func colliCheck(wg *sync.WaitGroup) {
	defer wg.Done()
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
	switch strings.ToLower(action) {
	case "excel":
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
	case "colli": // In build
		colliCheck(&wg)
	}
}
