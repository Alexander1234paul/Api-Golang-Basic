package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Abono struct {
	Id          string  `json:id`
	Cantidad    float32 `json:cantidad`
	Precio      int     `json:precio`
	Observacion string  `json:observacion`
}

// Array global de productos:
var Abonos []Abono

// Array global de producto
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h2>Servidor funcionando!</h2></body></html>")
	fmt.Println("Solicitud atendida: homePage")
}
func findAllAbonos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findAllAbonos")
	json.NewEncoder(w).Encode(Abonos)
}
func findAbonoById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findAbonoById")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, abono := range Abonos {
		if abono.Id == key {
			json.NewEncoder(w).Encode(abono)
		}
	}
}

func deleteAbono(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: deleteAbono")
	vars := mux.Vars(r)
	key := vars["id"]
	// buscar el producto a eliminar:
	for index, abono := range Abonos {
		if abono.Id == key {
			// borrar del array:
			Abonos = append(Abonos[:index], Abonos[index+1:]...)
		}
	}
}

func updateAbono(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: updateAbono")
	// Se obtiene el body desde el request y
	// se deserializa en una variable producto:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var abono Abono
	json.Unmarshal(reqBody, &abono)
	key := abono.Id
	// buscar el producto a actualizar:
	for index, p := range Abonos {
		if p.Id == key {
			// actualizar el array:
			Abonos[index] = abono
			break
		}
	}
	json.NewEncoder(w).Encode(abono)
}

func createNewAbono(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el body desde el request y
	// se deserializa en una variable producto:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var abono Abono
	json.Unmarshal(reqBody, &abono)
	// adicionamos en el array el nuevo producto:
	Abonos = append(Abonos, abono)
	json.NewEncoder(w).Encode(abono)
}

func iniciarServidor() {
	fmt.Println("API REST simple con lenguaje go.")
	ruteador := mux.NewRouter().StrictSlash(true)
	ruteador.HandleFunc("/", homePage)
	ruteador.HandleFunc("/abonos", findAllAbonos)
	//el orden de definicion es importante en el manejo de rutas:
	ruteador.HandleFunc("/abono", createNewAbono).Methods("POST")
	ruteador.HandleFunc("/abono/{id}", deleteAbono).Methods("DELETE")
	ruteador.HandleFunc("/abono/{id}", findAbonoById)
	ruteador.HandleFunc("/abono", updateAbono).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5050", ruteador))
}

func main() {
	Abonos = []Abono{
		Abono{Id: "1", Cantidad: 12, Precio: 2.0, Observacion: "NA"},
		Abono{Id: "2", Cantidad: 2, Precio: 23.0, Observacion: "NA"},
	}
	iniciarServidor()
}
