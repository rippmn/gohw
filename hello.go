package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
	index := os.Getenv("INDEX")
	if index != ""{
		fmt.Fprintf(w, "App Instance Index:%s\n", index)
	}else
	{
		fmt.Fprintf(w, "No App Instance available\n")
	}		
	
	fmt.Fprintf(w, "Request received at %s\n", r.URL.Path[1:])
}

func exitNow(w http.ResponseWriter, r *http.Request){
	log.Println("exit called")
	os.Exit(14)
}

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/exit", exitNow)
	port_num := os.Getenv("PORT")
	if port_num == ""{
		port_num = "8080"
	}
	port := ":" + port_num
	log.Printf("Listening on port %s\n", port)
	log.Println("****Cheat commandos, ROCK ROCK ON!!!!") 
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
}
}
