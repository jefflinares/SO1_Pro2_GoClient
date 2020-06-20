package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	var s Spec
	var casos Casos

	//escribir el archivo de casos
	//writeData(20000)
	readSpecifications(&s)

	casosRetorno := reaData(&casos, s.Path, s.NoRequests)

	makePostRequest(casosRetorno, s.Url)
	//makeGetRequest(&s, s.Url)
	return
}

func makeGetRequest(s *Spec, url string) {
	//realizar conexion con servidor nginx
	resp, err := http.Get(s.Url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func makePostRequest(s *Casos, url string) {
	//realizar conexion con servidor nginx
	jData, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jData))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}

func readSpecifications(s *Spec) {

	//Read de file of specifications
	configFile, err := ioutil.ReadFile("./Files/specs.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//Conver to a Go Struct (Spec)

	json.Unmarshal(configFile, &s)

	//Print all the information
	fmt.Println(s.Url)
	fmt.Println(strconv.Itoa(s.NoThreads))
	fmt.Println(strconv.Itoa(s.NoRequests))
	fmt.Println(s.Path)

}
