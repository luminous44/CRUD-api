package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func health(w http.ResponseWriter, r *http.Request){
	fmt.Println("Ok")
}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/healthCheck",health)
	// r.HandleFunc("/recipes",getAllRecipes)
	// r.HandleFunc("/recipe",getRecipe)
	// r.HandleFunc("/delete",deleteRecipe)
	// r.HandleFunc("/create",createRecipe)
	// r.HandleFunc("/update",updateRecipe)

	fmt.Println("server running on 8080")
    
	if err := http.ListenAndServe(":8080",r); err != nil{
        fmt.Println("Failed to start server",err)        
		log.Fatal(err)
	}

}