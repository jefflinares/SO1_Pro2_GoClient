package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

type CaseInsensitiveReplacer struct {
	toReplace   *regexp.Regexp
	replaceWith string
}

func NewCaseInsensitiveReplacer(toReplace, replaceWith string) *CaseInsensitiveReplacer {
	return &CaseInsensitiveReplacer{
		toReplace:   regexp.MustCompile("(?i)" + toReplace),
		replaceWith: replaceWith,
	}
}

func (cir *CaseInsensitiveReplacer) Replace(str string) string {
	return cir.toReplace.ReplaceAllString(str, cir.replaceWith)
}

type Casos struct {
	Casos []Caso `json:"casos"`
}

type Caso struct {
	Nombre         string `json:"nombre"`
	Departamento   string `json:"departamento"`
	Edad           int    `json:"edad"`
	Forma_contagio string `json:"forma_contagio"`
	Estado         string `json:"estado"`
}

func reaData(casos *Casos, path string) {

	//Read de file of specifications
	casos_data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//file_data := strings.Replace(strings.ToLower(string(casos_data)), "\"forma de contagio\"", "\"Contagio\"", -1)

	r := NewCaseInsensitiveReplacer("forma de contagio", "Forma_Contagio")
	//fmt.Println()

	var ap_ string = "{\n\t\"casos\":" + r.Replace(string(casos_data)) + "\n}"
	//Conver to a Go Struct (Spec)

	json.Unmarshal([]byte(ap_), &casos)

	//fmt.Println("imprimir los casos: " + file_data)
	//fmt.Println(ap_)
	//Print all the information
	for i := 0; i < len(casos.Casos); i++ {
		fmt.Println(casos.Casos[i].Nombre)
		fmt.Println(casos.Casos[i].Departamento)
		fmt.Println(strconv.Itoa(casos.Casos[i].Edad))
		fmt.Println(casos.Casos[i].Forma_contagio)
		fmt.Println(casos.Casos[i].Estado)
	}

}

func getRandomValues(n int) string {
	s := "{\n"

	nombre := []string{"Luis", "Dulce", "Antonia", "Pedro", "Pilar", "Fernando", "Jose", "Pablo", "Anna", "Camila", "Susana", "Erick", "Ronald", "Jeff"}
	apellido := []string{" Caceres", " Perez", " Juarez", " Setino", " Ortega", " Girón", " Hernández", " Gomez", " Flores", " Polanco", " López", " González"}
	departamento := []string{"Guatemala", "Escuintla", "Sacatepequez", "Chiquimula", "Huehuetenango", "Quetzaltenango", "Peten", "Zacapa", "Santa Rosa"}
	formas_contagio := []string{"Comunitario", "viral"}
	estado := []string{"Activo", "Recuperado", "Fallecido"}

	for i := 0; i < n; i++ {
		s += "\t\"Nombre\":\"" + nombre[rand.Intn(len(nombre)+1)] + "\"\n"
		s += "\t\"Departamento\":\"" + departamento[rand.Intn(len(departamento)+1)] + "\"\n"
		s += "\t\"Edad\":" + string(rand.Intn(80-15+1)) + "\n"
		s += "\t\"Forma de Contagio\":\"" + formas_contagio[rand.Intn(len(formas_contagio)+1)] + "\"\n"
		s += "\t\"Estado\":\"" + estado[rand.Intn(len(estado)+1)] + "\"\n"
	}

	return s
}
