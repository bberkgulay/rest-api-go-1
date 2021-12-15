package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// API response struct with message.
type API struct {
	Message string "json:message"
}

func main() {
	// a new router is created.
	router := mux.NewRouter()

	// endpoints are defined and mapped with functions.
	router.HandleFunc("/api/sum/{first:[0-9]+}/{second:[0-9]+}", sum)
	router.HandleFunc("/api/subtract/{first:[0-9]+}/{second:[0-9]+}", subtract)
	router.HandleFunc("/api/multiply/{first:[0-9]+}/{second:[0-9]+}", multiply)
	router.HandleFunc("/api/divide/{first:[0-9]+}/{second:[0-9]+}", divide)
	router.HandleFunc("/api/modulo/{first:[0-9]+}/{second:[0-9]+}", modulo)
	router.HandleFunc("/api/isPerfectNumber/{number:[0-9]+}", isPerfectNumber)
	http.Handle("/", router)

	http.ListenAndServe(":8090", nil)
}

// @route       GET /api/v1/sum/:first/:second/
// @example     GET /api/v1/sum/12/9 --> 21
// @access      Public
// @description It takes two parameters and return sum of them.
func sum(w http.ResponseWriter, r *http.Request) {

	//first and second parameters are getted from request.
	urlParams := mux.Vars(r)
	first := urlParams["first"]
	second := urlParams["second"]

	//parameters are converted from string to int to sum.
	firstAsNum, err := strconv.Atoi(first)
	secondAsNum, err := strconv.Atoi(second)

	messageConcat := "Result : " + strconv.Itoa(int(firstAsNum+secondAsNum))

	//created response object as json.
	message := API{messageConcat}
	output, err := json.Marshal(message)

	//error log
	if err != nil {
		fmt.Println("Error.")
	}

	fmt.Fprintf(w, string(output))
}

// @route       GET /api/v1/subtract/:first/:second/
// @example     GET /api/v1/subtract/12/9 --> 3
// @access      Public
// @description It takes two parameters and return subtract of them.
func subtract(w http.ResponseWriter, r *http.Request) {

	//first and second parameters are getted from request.
	urlParams := mux.Vars(r)
	first := urlParams["first"]
	second := urlParams["second"]

	//parameters are converted from string to int to subtract.
	firstAsNum, err := strconv.Atoi(first)
	secondAsNum, err := strconv.Atoi(second)

	messageConcat := "Result : " + strconv.Itoa(int(firstAsNum-secondAsNum))

	//created response object as json.
	message := API{messageConcat}
	output, err := json.Marshal(message)

	//error log
	if err != nil {
		fmt.Println("Error.")
	}

	fmt.Fprintf(w, string(output))
}

// @route       GET /api/v1/multiply/:first/:second/
// @example     GET /api/v1/multiply/10/9 --> 90
// @access      Public
// @description It takes two parameters and return multiply of them.
func multiply(w http.ResponseWriter, r *http.Request) {

	//first and second parameters are getted from request.
	urlParams := mux.Vars(r)
	first := urlParams["first"]
	second := urlParams["second"]

	//parameters are converted from string to int to multiply
	firstAsNum, err := strconv.Atoi(first)
	secondAsNum, err := strconv.Atoi(second)
	messageConcat := "Result : " + strconv.Itoa(int(firstAsNum*secondAsNum))

	//created response object as json.
	message := API{messageConcat}
	output, err := json.Marshal(message)

	//error log
	if err != nil {
		fmt.Println("Error.")
	}

	fmt.Fprintf(w, string(output))
}

// @route       GET /api/v1/divide/:first/:second/
// @example     GET /api/v1/divide/12/4 --> 3
// @access      Public
// @description It takes two parameters and return divide of them.
func divide(w http.ResponseWriter, r *http.Request) {

	//first and second parameters are getted from request.
	urlParams := mux.Vars(r)
	first := urlParams["first"]
	second := urlParams["second"]

	//parameters are converted from string to int to multiply
	firstAsNum, err := strconv.Atoi(first)
	secondAsNum, err := strconv.Atoi(second)
	messageConcat := "Result : " + strconv.FormatFloat(float64(firstAsNum)/float64(secondAsNum), 'f', -1, 64)

	//created response object as json.
	message := API{messageConcat}
	output, err := json.Marshal(message)

	//error log
	if err != nil {
		fmt.Println("Error.")
	}

	fmt.Fprintf(w, string(output))
}

// @route       GET /api/v1/modulo/:first/:second/
// @example     GET /api/v1/modulo/95/10 --> 5
// @access      Public
// @description It takes two parameters and return modulo of them.
func modulo(w http.ResponseWriter, r *http.Request) {

	//first and second parameters are getted from request.
	urlParams := mux.Vars(r)
	first := urlParams["first"]
	second := urlParams["second"]

	//parameters are converted from string to int to modulo
	firstAsNum, err := strconv.Atoi(first)
	secondAsNum, err := strconv.Atoi(second)
	messageConcat := "Result : " + strconv.Itoa(int(firstAsNum%secondAsNum))

	//created response object as json.
	message := API{messageConcat}
	output, err := json.Marshal(message)

	//error log
	if err != nil {
		fmt.Println("Error.")
	}

	fmt.Fprintf(w, string(output))
}

// @route       GET /api/v1/isPerfectNumber/:number/
// @example     GET /api/v1/isPerfectNumber/7 --> The number is not a Perfect Number
// @access      Public
// @description Finds whether it is a perfect number or not.
func isPerfectNumber(w http.ResponseWriter, r *http.Request) {

	//first and second parameters are getted from request.
	urlParams := mux.Vars(r)
	number := urlParams["number"]

	//parameter is converted from string to int to isPerfectNumber
	numberInt, err := strconv.Atoi(number)

	//Sums the divisors of number from zero to the number itself.
	tempSum := 0
	for i := 1; i < numberInt; i++ {
		if numberInt%i == 0 {
			tempSum += i
		}
	}

	messageConcat := ""

	// Check if the sum of the divisors equals the number itself.
	if tempSum == numberInt {
		messageConcat = "The number is a Perfect number!"
	} else {
		messageConcat = "The number is not a Perfect number!"
	}

	//created response object as json.
	message := API{messageConcat}
	output, err := json.Marshal(message)

	//error log
	if err != nil {
		fmt.Println("Error.")
	}
	fmt.Fprintf(w, string(output))
}
