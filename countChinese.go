package main

//source from jimmiysong

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var count int64
var inputFile string

func main() {
	flag.StringVar(&inputFile, "f", "", "filename")
	flag.Parse()
	if inputFile == "" {
		fmt.Printf("You need a filename , bro, just like -f <filename> \n")
		os.Exit(0)
	}

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Readding Error : %s \n", err)
	}

	for _, j := range string(buf) {
		if int64(j) > hex2oct("0x4e00") && int64(j) < hex2oct("0x9fa5") {
			count++
		}
	}

	fmt.Println(inputFile, count)
}

func hex2oct(hex string) int64 {
	oct, _ := strconv.ParseInt(hex, 0, 64)
	return oct
}
