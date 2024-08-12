package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Step 1
// URL Structure Created where we goan save data
type URL struct {
	ID            string    `json:"id"`
	OriginalURL   string    `json:"Original_url"`
	Short_URL     string    `json:"short_URL"`
	Creation_date time.Time `json:"Creation_dater"`
}

// Step 2
// Created urlDB map to store data in Key value pair

/*
Hum data ko map krne wale hai

	d7557f5{
			ID:"d7557f5"
			OriginalURL :" Long URL"
			Short_URL:"Short ULR"
			creation_date : "time.now"
	}
*/

var urlDB = make(map[string]URL) // create map

// Step 3
// created Function
func GenerateShort_URL(OriginalURL string) string {

	//Step 3.1
	// hum yaha per hasing algo use krne wale hai long URL ko short krne keliye
	hasher := md5.New()

	//Step 3.2
	//Convert String data to bytes
	hasher.Write([]byte(OriginalURL)) // Its convert OriginalURL  string into byte Slices
	println("String Data", OriginalURL)
	println()
	fmt.Println("Converted String data into Bytes data", hasher)
	println()

	//Step 3.3
	data := hasher.Sum(nil) //sum of bytes slices
	fmt.Println("Hash data", data)
	println()

	//Step 3.4
	hash := hex.EncodeToString(data)
	fmt.Println("Encode Bytes data into string", hash)
	println()

	fmt.Println("Final String", hash[:8])

	return hash[:8]

}

// Step 4
func createUrl(OriginalURL string) string {

	//Step 4.1
	shortURL := GenerateShort_URL(OriginalURL)

	// Here we use short url as ID
	id := shortURL

	// Step 4.2
	// insert URL type data into map
	urlDB[id] = URL{
		ID:            id,
		OriginalURL:   OriginalURL,
		Short_URL:     shortURL,
		Creation_date: time.Now(),
	}
	return shortURL
}

//Step 5

func getURL(id string) (URL, error) {
	//step 5.1
	url, ok := urlDB[id]
	fmt.Println("GetURL function:", url)
	if !ok {
		return URL{}, errors.New("URL Not Found")
	}
	fmt.Println(url)
	return url, nil

}

// step 6.2
func RootpageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world") //Fprintf kay krta hai ise jaha bolagya hai waha likhta hai insted or stander output windown
}

func ShortURLhandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid requset body", http.StatusBadRequest)
		return
	}
	shortURL_ := createUrl(data.URL)

	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(response)

}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):] // Extract the ID from the URL path
	println("id redirect ", id)
	url, err := getURL(id)
	fmt.Println("redirect URL Handler", url)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {

	//OriginalURL := "www.linkedin.com/in/pritam-bhure-2aa821314"
	//GenerateShort_URL("https://www.youtube.com/watch?v=NTpbbQUBbuo")

	//step 6.1
	//Handle Function
	//Register the handle Function to Handle all request to the root URL("/")
	http.HandleFunc("/", RootpageURL)
	http.HandleFunc("/shorten", ShortURLhandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	//Step 6
	// Start HTTP server on port 8000
	fmt.Println("Starting Server on post 8000....")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Error on Starting Server:", err)
	}

}
