package server

import (
	"errors"
	"net/http"
	"converter"
)


func getPhraseFromNum(w http.ResponseWriter, r *http.Request) {
	
	numb, err := stringtoUint(r.FormValue("Number"))
	if err != nil {
		WriteError(w, errors.New("Error parsing Number form value"), 400)
		return
	}

	adj, noun, err := converter.TransNumToPhrase(numb)
	if err != nil {
		WriteError(w, err, 400)
		return
	}

	retMap := make(map[string]string)
	retMap["Adjective"] = adj
	retMap["Noun"] = noun

	WriteJson(w, retMap)

}


func getNumFromPhrase(w http.ResponseWriter, r *http.Request) {
	
	adj := r.FormValue("Adjective")
	if adj == "" {
		WriteError(w, errors.New("Error parsing Adjective form value"), 400)
		return
	}

	noun := r.FormValue("Noun")
	if noun == "" {
		WriteError(w, errors.New("Error parsing Noun form value"), 400)
		return
	}

	numb, err := converter.TransPhraseToNum(adj, noun)
	if err != nil {
		WriteError(w, err, 400)
		return
	}

	retMap := make(map[string]uint)
	retMap["Number"] = numb

	WriteJson(w, retMap)

}
