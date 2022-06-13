package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Mutant struct {
	dna []string `json:"dna"`
}

func (mutant Mutant) isMutant(result []byte, pat int) bool {

	var isMutant bool = false
	var countPat = 0
	var diagonalPat = pat + 1
	var verticalPat = pat
	var horizontalPat = 1
	var diagonalLimit = len(result) - diagonalPat*3
	var verticalLimit = len(result) - verticalPat*3
	var horizontaLimit = len(result) - horizontalPat*3

	for i := 0; i < len(result); i++ {

		if i < diagonalLimit {
			if (result[i] == result[i+diagonalPat]) &&
				(result[i] == result[i+diagonalPat*2]) &&
				(result[i] == result[i+diagonalPat*3]) {

				fmt.Println("Diagonal Pat")
				countPat++
			}
		}
		if i < verticalLimit {
			if (result[i] == result[i+verticalPat]) &&
				(result[i] == result[i+verticalPat*2]) &&
				(result[i] == result[i+verticalPat*3]) {

				fmt.Println("Vertical Pat")
				countPat++
			}
		}
		if i < horizontaLimit {
			if (result[i] == result[i+horizontalPat]) &&
				(result[i] == result[i+horizontalPat*2]) &&
				(result[i] == result[i+horizontalPat*3]) {

				fmt.Println("Horizontal Pat")
				countPat++
			}
		}
	}

	if countPat > 1 {
		isMutant = true
		fmt.Println("Is Mutant")

	}

	return isMutant
}

func (mutant Mutant) convertToByte() []byte {

	var result []byte
	for _, element := range mutant.dna {
		for _, e := range []byte(element) {
			result = append(result, e)
		}

	}
	return result
}

func (mutant Mutant) searchPat() int {

	var count = 0
	for i := 0; i < len(mutant.dna); i++ {
		count++
	}
	return count
}

func postIsMutant(w http.ResponseWriter, r *http.Request) {

	var mutant1 Mutant

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error")
		return
	}

	json.Unmarshal(reqBody, &mutant1)

	fmt.Println("valor Mutante1", reqBody)
	fmt.Println(mutant1.dna)

	myString := string(reqBody[:])
	fmt.Println("valor string", myString)

	intento := Mutant{dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}

	if Mutant.isMutant(Mutant{}, intento.convertToByte(), intento.searchPat()) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("true")
	} else {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Error")
	}

	w.Header().Set("Content-Type", "application/json")

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/mutant", postIsMutant).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
