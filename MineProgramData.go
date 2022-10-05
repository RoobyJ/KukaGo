package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	walkFolders = [...]string{"R1", "BMW_App"}
)

func getPathTree() {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
			return err
		}
		if info.IsDir() && contains(info.Name()) {
			fmt.Printf("File Name: %s\n", info.Name())
			return nil
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func contains(name string) bool {
	for _, str := range walkFolders {
		if str == name {
			return true
		}
	}
	return false
}
