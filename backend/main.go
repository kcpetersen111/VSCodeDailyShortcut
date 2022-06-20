package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

type data struct{
	Easy []shortCut `json:"easy"`
	Medium []shortCut `json:"medium"`
	Hard []shortCut	`json:"hard"`

}
type shortCut struct{
	Description string `json:"description"`
	Mac bool `json:"mac"`
	Windows bool `json:"windows"`
	TimesShown int `json:"timesShown"`
	Id int `json:"id"`
}
// type tester struct{
// 	Asdf string 
// 	Number int
// }

var	masterList data


func main() {

	//open the data base
	// test := tester{
	// 	Asdf:"asdf",
	// 	Number: 10,
	// }
	test := shortCut{
		Description:"test",
		Mac:false,
		Windows:true,
		TimesShown:0,
		Id:0,
	}
	testEasyArr := make([]shortCut,0)
	testEasyArr= push(testEasyArr, test)

	testData :
	testArr, err := json.Marshal(test)

	if err != nil{
		return;
	}
	fmt.Println(string(testArr))
		

	//start the server 
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/getShortCut", getShortCut).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080",myRouter))
}

func getShortCut(w http.ResponseWriter, r *http.Request){
	
}




// [difficulty:{
	// description:
	// mac:
	// windows:
	// timesShown:
	// id

// }]

//difficulty of command
	//