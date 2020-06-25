package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
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

//struct casos
type Casos struct {
	Casos []*Caso `json:"casos"`
}

//stuct caso
type Caso struct {
	Nombre         string `json:"nombre"`
	Departamento   string `json:"departamento"`
	Edad           int    `json:"edad"`
	Forma_Contagio string `json:"forma_contagio"`
	Estado         string `json:"estado"`
}

func reaData(casos *Casos, path string, n int) *Casos {

	//Read de file of specifications
	casosData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//file_data := strings.Replace(strings.ToLower(string(casos_data)), "\"forma de contagio\"", "\"Contagio\"", -1)

	r := NewCaseInsensitiveReplacer("forma de contagio", "Forma_Contagio")
	//fmt.Println()

	var ap string = "{\n\t\"Casos\":" + r.Replace(string(casosData)) + "\n}"
	//Conver to a Go Struct (Spec)

	json.Unmarshal([]byte(ap), &casos)

	// fmt.Println("Imprimir casos")
	//Print all the information

	var casosRetorno Casos

	if len(casos.Casos) == 0 {
		fmt.Println("No leyo los datos")

	} else {

		if len(casos.Casos) < n {

			for i := 0; i < len(casos.Casos); i++ {
				casosRetorno.Casos = append(casosRetorno.Casos, casos.Casos[i])
			}

			for i := len(casos.Casos); i < n; i++ {
				casosRetorno.Casos = append(casosRetorno.Casos, getNewRandomValue())
			}
			return &casosRetorno
		}

		for i := 0; i < n && i < len(casos.Casos); i++ {

			casosRetorno.Casos = append(casosRetorno.Casos, casos.Casos[i])
			/*fmt.Println(casos.Casos[i].Nombre)
			fmt.Println(casos.Casos[i].Departamento)
			fmt.Println(strconv.Itoa(casos.Casos[i].Edad))
			fmt.Println(casos.Casos[i].Forma_Contagio)
			fmt.Println(casos.Casos[i].Estado)*/
		}
	}
	return &casosRetorno

}

func getRandomValues(n int) string {
	s := "[\n"

	nombre := []string{"Luis", "Dulce", "Antonia", "Pedro", "Pilar", "Fernando", "Jose", "Pablo", "Anna", "Camila", "Susana", "Erick", "Ronald", "Jeff"}
	apellido := []string{" Caceres", " Perez", " Juarez", " Setino", " Ortega", " Girón", " Hernández", " Gomez", " Flores", " Polanco", " López", " González"}
	departamento := []string{"Guatemala", "Escuintla", "Sacatepequez", "Chiquimula", "Huehuetenango", "Quetzaltenango", "Peten", "Zacapa", "Santa Rosa"}
	formasContagio := []string{"Comunitario", "viral"}
	estado := []string{"Activo", "Recuperado", "Fallecido"}

	for i := 0; i < n; i++ {
		if i != 0 {
			s += "\t\t,\n"
		}
		s += "\t\t{\n"
		s += "\t\t\t\"Nombre\":\"" + nombre[rand.Intn(len(nombre))] + apellido[rand.Intn(len(apellido))] + "\",\n"
		s += "\t\t\t\"Departamento\":\"" + departamento[rand.Intn(len(departamento))] + "\",\n"
		s += "\t\t\t\"Edad\":" + fmt.Sprintf("%v", (rand.Intn(80-15+1))) + ",\n"
		s += "\t\t\t\"Forma de Contagio\":\"" + formasContagio[rand.Intn(len(formasContagio))] + "\",\n"
		s += "\t\t\t\"Estado\":\"" + estado[rand.Intn(len(estado))] + "\"\n"
		s += "\t\t}\n"
	}

	s += "\t]"
	return s
}

func getNewRandomValue() *Caso {
	nombre := []string{"Luis", "Dulce", "Antonia", "Pedro", "Pilar", "Fernando", "Jose", "Pablo", "Anna", "Camila", "Susana", "Erick", "Ronald", "Jeff"}
	apellido := []string{" Caceres", " Perez", " Juarez", " Setino", " Ortega", " Girón", " Hernández", " Gomez", " Flores", " Polanco", " López", " González"}
	departamento := []string{"Guatemala", "Escuintla", "Sacatepequez", "Chiquimula", "Huehuetenango", "Quetzaltenango", "Peten", "Zacapa", "Santa Rosa"}
	formasContagio := []string{"Comunitario", "viral"}
	estado := []string{"Activo", "Recuperado", "Fallecido"}

	return &Caso{Nombre: nombre[rand.Intn(len(nombre))] + apellido[rand.Intn(len(apellido))],
		Departamento:   departamento[rand.Intn(len(departamento))],
		Forma_Contagio: formasContagio[rand.Intn(len(formasContagio))],
		Edad:           (rand.Intn(80 - 15 + 1)), Estado: estado[rand.Intn(len(estado))]}

}

func writeData(n int) {
	d1 := []byte(getRandomValues(n))
	err := ioutil.WriteFile("./Files/casos.json", d1, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
