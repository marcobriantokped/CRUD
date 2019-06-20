package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      int    `json:"ID"`
	Name    string `json:Name`
	Content string `json:Content`
}

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ENDPOINT HIT")
	fmt.Fprintf(os.Stdout, "ENDPOINT HIT")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := json.Marshal(Articles)
	w.Write(jsonData)
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/delete/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/update", updateArticle).Methods("PUT")
	fmt.Println("")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if idint, _ := strconv.Atoi(id); article.ID == idint {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
			message := fmt.Sprintf("Deleted %+v", article)
			w.Write([]byte(message))
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	/*//get body of post request
	var reqBody []byte
	r.Body.Read(reqBody)
	fmt.Fprintf(w, "%+v", string(reqBody))*/

}

func updateArticle(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	id := article.ID
	for index, value := range Articles {
		// if our id path parameter matches one of our
		// articles
		if value.ID == id {
			// updates our Articles array to new update
			Articles[index] = article
			message := fmt.Sprintf("Updated %+v into %+v", value, article)
			w.Write([]byte(message))
			return
		}
	}
	fmt.Println("Fail to update")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {

	mapping := mux.Vars(r)
	value := mapping["id"]
	for _, article := range Articles {
		if valueint, _ := strconv.Atoi(value); valueint == article.ID {
			message, _ := json.Marshal(article)
			w.Write(message)
			return
		}
	}
	w.Write([]byte("No entry found"))
}
func main() {
	article1 := Article{ID: 1, Name: "Marco", Content: "Intern"}
	article2 := Article{ID: 2, Name: "Brian", Content: "Intern"}
	Articles = append(Articles, article1)
	Articles = append(Articles, article2)
	handleRequest()
}
