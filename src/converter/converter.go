package converter

import (
	"errors"
)



func TransNumToPhrase(numb uint) (string, string, error){

	if numb > 65535 || numb < 0 {
		return "", "", errors.New("Invalid number")
	}

	adj := numb / 256
	noun := numb % 256

	return ADJ_WORDS[adj], NOUN_WORDS[noun], nil 
		
}


func TransPhraseToNum(adj string, noun string) (uint, error){

	adjnum := -1
	nounnum := -1

	for i, listword := range ADJ_WORDS {
		 if adj == listword {
		 	adjnum = i
		 	break
		 }
	}

	if adjnum == -1 {
		return 0, errors.New("Adjective not found")
	}
	
	for i, listword := range NOUN_WORDS {
		 if noun == listword {
		 	nounnum = i
		 	break
		 }
	}

	if nounnum == -1 {
		return 0, errors.New("Noun not found")
	}

	number := uint(adjnum * 256 + nounnum)

	return number, nil 
		
}