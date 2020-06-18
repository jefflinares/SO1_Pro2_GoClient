package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Spec struct {
	Url        string `json:"url"`
	NoThreads  int    `json:"noThreads"`
	NoRequests int    `json:"noRequests"`
	Path       string `json:"path"`
}

func main() {
	//Read de file of specifications
	configFile, err := ioutil.ReadFile("./specs.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//Conver to a Go Struct (Spec)
	var s Spec
	json.Unmarshal(configFile, &s)

	//Print all the information
	fmt.Println(s.Url)
	fmt.Println(strconv.Itoa(s.NoThreads))
	fmt.Println(strconv.Itoa(s.NoRequests))
	fmt.Println(s.Path)
	return
}
