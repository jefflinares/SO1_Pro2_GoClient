package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//struct casos
type Casos struct {
	Casos []Caso `json:"casos"`
}

//stuct caso
type Caso struct {
	Nombre         string `json:"nombre"`
	Departamento   string `json:"departamento"`
	Edad           int    `json:"edad"`
	Forma_Contagio string `json:"forma_contagio"`
	Estado         string `json:"estado"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":

		var casos Casos
		err := json.NewDecoder(r.Body).Decode(&casos)
		//datos, err := json.Marshal(data)

		if err != nil {
			log.Fatal(err)
		}

		var s string
		/*
		for i := 0; i < len(casos.Casos); i++ {
			s += "{\n"
			s += "\tNombre:" + casos.Casos[i].Nombre + ",\n"
			s += "\tDepartamento:" + casos.Casos[i].Departamento + ",\n"
			s += "\tEdad:" + strconv.Itoa(casos.Casos[i].Edad) + ",\n"
			s += "\tForma contagio:" + casos.Casos[i].Forma_Contagio + ",\n"
			s += "\tEstado:" + casos.Casos[i].Estado + ",\n"
			s += "}"
		}
		*/
		s = "Casos Recibidos: "+strconv.Itoa(len(casos.Casos))
		//fmt.Println("Json recivido: " + s)
		fmt.Fprintf(w, s)
		//decoder := json.NewDecoder(r.Body)
		/*
			var t test_struct
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			log.Println(t.Test)
		*/

		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		/*
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
			name := r.FormValue("name")
			address := r.FormValue("address")
			fmt.Fprintf(w, "Name = %s\n", name)
			fmt.Fprintf(w, "Address = %s\n", address)
		*/

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
