package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := "./"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for index, f := range files {
		fmt.Println(f.Name())
		os.Rename(path+f.Name(), path+fmt.Sprintf("20200208-%02d.jpg", index))
	}
}
