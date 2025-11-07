package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
var recipesList []Recipe
type Recipe struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Cooker *Cooker `json:"cooker"`
}
type Cooker struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func health(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Ok")
}

func getAllRecipes(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(recipesList); err != nil{
		http.Error(w,"Failed to patch data",http.StatusNotFound)
	}

}
func getRecipeById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range recipesList{

		if item.ID == params["id"] {
		 if err := json.NewEncoder(w).Encode(item); err != nil{
		    http.Error(w,"Failed to patch data",http.StatusNotFound)
	      }
		 return
		}
	}
  http.Error(w,"Invalid Id",http.StatusNoContent)
}
func deleteRecipeById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range recipesList{
		if item.ID == params["id"] {
         recipesList = append(recipesList[:index],recipesList[index+1:]...)
		 if err := json.NewEncoder(w).Encode(item); err != nil{
		    log.Fatal(err)
	      }
		 return
		}
	}
   http.Error(w,"Invalid Id",http.StatusNoContent)
}
func main() {
    recipesList = append(recipesList, Recipe{ID: "1",Name: "Biriyani",Type:"Lunch",Cooker: &Cooker{Name: "Hina",Age: 30}})
	recipesList = append(recipesList, Recipe{ID: "2",Name: "Polaw",Type:"Dinner",Cooker: &Cooker{Name: "Zoe",Age: 20}})
	r := mux.NewRouter()
	r.HandleFunc("/healthCheck",health)
	r.HandleFunc("/recipes",getAllRecipes).Methods("GET")
	r.HandleFunc("/recipe/{id}",getRecipeById).Methods("GET")
	r.HandleFunc("/delete/{id}",deleteRecipeById).Methods("DELETE")
	// r.HandleFunc("/create",createRecipe)
	// r.HandleFunc("/update",updateRecipe)

	fmt.Println("server running on 8080")
    
	if err := http.ListenAndServe(":8080",r); err != nil{
        fmt.Println("Failed to start server",err)        
		log.Fatal(err)
	}

}