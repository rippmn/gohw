package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

	var logger  = log.New(os.Stdout, log.Prefix(), log.Flags())

func handler(w http.ResponseWriter, r *http.Request){
	

	logger.Printf("Request received at %s", r.URL.Path[1:])
	index := os.Getenv("CF_INSTANCE_INDEX")
	if index != ""{
		fmt.Fprintf(w, "App Instance Index:%s\n", index)
	}else
	{
		fmt.Fprintf(w, "No App Instance available\n")
	}		
	
}

func exitNow(w http.ResponseWriter, r *http.Request){
	logger.Println("exit called")
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
	logger.Printf("Listening on port %s\n", port)
	logger.Println("****Cheat commandos, ROCK ROCK ON!!!!") 
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
}
}
