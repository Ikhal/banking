package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	
	"net/http"
	
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}


func greet (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}  

func getAllCustomers(w http.ResponseWriter, r *http.Request)  {
	customers := []Customer {
		{Name: "Ashish", City: "Abuja", Zipcode: "900102"},
		{Name: "Ibrahim", City: "Lagos", Zipcode: "900209"},
	}
	if r.Header.Get("Content-Type")== "application/xml"{
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
	
}