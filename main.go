package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"web-test/informat"
)

func main() {
	configFilePath := flag.String("f", "", "informat file path")
	flag.Parse()

	if *configFilePath == "" {
		fmt.Println("need to set informat file with -f option")
		os.Exit(1)
	}

	f, e := os.Open(*configFilePath)
	if e != nil {
		fmt.Printf("cannot open file : %s\n", *configFilePath)
		os.Exit(1)
	}
	defer f.Close()

	byteValue, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Printf("cannot open file : %s\n", *configFilePath)
		os.Exit(1)
	}

	var config informat.TestConfig
	e = json.Unmarshal(byteValue, &config)
	if e != nil {
		fmt.Printf("informat informat error : %s\n%s", *configFilePath, e.Error())
		os.Exit(1)
	}

	fmt.Println(config)
}
