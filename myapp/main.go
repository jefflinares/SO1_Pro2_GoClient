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
	"time"
	"math/rand"
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

	fmt.Print("\nCasos totales leídos del archivo: ")
	fmt.Println(len(casosRetorno.Casos))
	var NRequestPerThread int = s.NoRequests / s.NoThreads
	indexInicial := 0

	//ENVIAR POR HILO LA PORCIO DE REQUEST EN CADA UNO
	channels := make(chan string, s.NoThreads)
	for i := 0; i < s.NoThreads; i++ {
		var casosAux Casos
		if (i + 1) == s.NoThreads {
			//El ultimo hilo
			casosAux.Casos = casosRetorno.Casos[indexInicial:len(casosRetorno.Casos)]
			go makePostRequest(i, &casosAux, s.Url, channels)
		} else {
			casosAux.Casos = casosRetorno.Casos[indexInicial : NRequestPerThread+indexInicial]
			go makePostRequest(i, &casosAux, s.Url, channels)
		}
		indexInicial += NRequestPerThread
	}

	count := 0
	for elem := range channels {
		if count == s.NoThreads - 1 {
			close(channels)
		}
		count++
		fmt.Println(elem)
	}

	//makePostRequest(casosRetorno, s.Url)
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

func makePostRequest(index int, s *Casos, url string, channels chan string) {
	// Para simular una carga de trabajo
    // dormimos el programa x cantidad de segundo
    // donde x puede ir de x a 15
	var seconds int
    seconds = rand.Intn(15)
	time.Sleep(time.Duration(seconds) * time.Second)

	// realizar conexion con servidor
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
	cadena := fmt.Sprintf("Hilo No. %d", index + 1)
	channels <- cadena
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
	fmt.Println("Url:        ", s.Url)
	fmt.Println("NoThreads:  ", strconv.Itoa(s.NoThreads))
	fmt.Println("NoRequests: ", strconv.Itoa(s.NoRequests))
	fmt.Println("Path:       ", s.Path)

}
