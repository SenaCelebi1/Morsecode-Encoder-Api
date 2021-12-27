package service

import (
	"encoding/json"
	"fmt"
	"morsecode/model"
	"net/http"

	"github.com/gorilla/mux"
)

func MorseCodeFunc(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var morsecode model.MorseCode
	_ = json.NewDecoder(r.Body).Decode(&morsecode)

	params := mux.Vars(r)
	text := params["text"]
	fmt.Println("text:")
	fmt.Println(text)

	morseMap := make(map[string]string)
	morseMap = map[string]string{
		"a":  ".-",
		"b":  "-...",
		"c":  "-.-.",
		"d":  "-..",
		"e":  ".",
		"f":  "..-.",
		"g":  "--.",
		"h":  "....",
		"i":  "..",
		"j":  ".---",
		"k":  "-.-",
		"l":  ".-..",
		"m":  "--",
		"n":  "-.",
		"o":  "---",
		"p":  ".--.",
		"q":  "--.-",
		"r":  ".-.",
		"s":  "...",
		"t":  "-",
		"u":  "..-",
		"v":  "...-",
		"w":  ".--",
		"x":  "-..-",
		"y":  "-.--",
		"z":  "--..",
		"ä":  ".-.-",
		"ö":  "---.",
		"ü":  "..--",
		"Ch": "----",
		"0":  "-----",
		"1":  ".----",
		"2":  "..---",
		"3":  "...--",
		"4":  "....-",
		"5":  ".....",
		"6":  "-....",
		"7":  "--...",
		"8":  "---..",
		"9":  "----.",
		".":  ".-.-.-",
		",":  "--..--",
		"?":  "..--..",
		"!":  "..--.",
		":":  "---...",
		"\"": ".-..-.",
		"'":  ".----.",
		"=":  "-...-",
	}

	var result = ""

	for i := 0; i < len(text); i++ {

		if v, found := morseMap[string(text[i])]; found {
			result += string(v)
			fmt.Println(string(text[i]))
			fmt.Println(v)
		}

	}

	morsecode.Code = result

	_ = json.NewDecoder(r.Body).Decode(&morsecode)
	json.NewEncoder(w).Encode(morsecode)

}
