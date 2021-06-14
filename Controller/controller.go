package controller

import (
	"MartinCandidate/translator"
	"encoding/json"
	"fmt"
	"net/http"
)

// Default handler for all other routes, returning 404 response.
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404")
}

// Handles the word endpoint, translates the word and returns the response.
func WordHandler(w http.ResponseWriter, req *http.Request) {
	// We only support POST methods on this endpoint.
	if "POST" != req.Method {
		http.Error(w, "POST endpoint", http.StatusBadRequest)
		return
	}

	var word translator.Translator

	// Decode the word into the translator struct.
	err := json.NewDecoder(req.Body).Decode(&word)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Translate the word and return errors if any.
	translatedWord, err := word.TranslateWord()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Map the response so that we return a correct JSON.
	response := map[string]string{"gopher-word": translatedWord}
	result, _ := json.Marshal(response)

	// Return the result.
	fmt.Fprintf(w, "%s", result)
	return
}

func SentenceHandler(w http.ResponseWriter, req *http.Request) {
	// We only support POST methods on this endpoint.
	if "POST" != req.Method {
		http.Error(w, "POST endpoint", http.StatusBadRequest)
		return
	}

	var sentence translator.Translator

	// Decode the sentence into the translator struct.
	err := json.NewDecoder(req.Body).Decode(&sentence)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Translate the sentence and return errors if any.
	translatedSentence, err := sentence.TranslateSentence()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Map the response so that we return a correct JSON.
	response := map[string]string{"gopher-sentence": translatedSentence}
	result, _ := json.Marshal(response)

	// Return the result.
	fmt.Fprintf(w, "%s", result)
	return
}

func History(w http.ResponseWriter, req *http.Request) {

	return
}

func ControllerInit() {
	http.HandleFunc("/word", WordHandler)
	http.HandleFunc("/sentence", SentenceHandler)
	http.HandleFunc("/history", History)

	http.HandleFunc("/", DefaultHandler)
}
