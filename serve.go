package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// APIURL Entry Point for Mux Gorilla Client
const APIURL = "https://reqres.in/api/users"

// Page is Root element
type Page struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []User
}

// PageDTO is root element of DTO
type PageDTO struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"perPage"`
	Total      int       `json:"total"`
	TotalPages int       `json:"totalPages"`
	Users      []UserDTO `json:"users"`
}

// User is a list Data
type User struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string
}

// UserDTO main model
type UserDTO struct {
	ID     int
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// ToDTO Convert Page to PageDTO
func ToDTO(page Page) PageDTO {
	usersDto := make([]UserDTO, 0)
	for _, user := range page.Data {
		usersDto = append(usersDto, UserDTO{ID: user.ID, Name: user.FirstName + " " + user.LastName, Avatar: user.Avatar})
	}
	return PageDTO{Page: page.Page, PerPage: page.PerPage, Total: page.Total, TotalPages: page.TotalPages, Users: usersDto}
}

// GetUsers from Chino Urls
func GetUsers(w http.ResponseWriter, r *http.Request) {

	requestDump, _ := httputil.DumpRequest(r, true)
	log.Println(string(requestDump))

	ch := make(chan Page)
	page := r.URL.Query().Get("page")
	go makeRequest(APIURL+"?page="+page, ch)
	json.NewEncoder(w).Encode(ToDTO(<-ch))
}

// Handlers define Entry Points API.  Declare Mux Router
func Handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", GetUsers).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	return router
}

// Make Async Request to url parameter, and set response in a Channel of Page
func makeRequest(url string, ch chan<- Page) {
	response, _ := http.Get(url)
	responseData, parseErr := ioutil.ReadAll(response.Body)
	if parseErr != nil {
		log.Fatal("Something is wrong: " + parseErr.Error())
	}
	var page Page
	json.Unmarshal([]byte(responseData), &page)
	ch <- page
}

// Main Function => Run HTTP Server
func main() {
	fmt.Println("Server Starting ...")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":4000", handlers.CORS(corsObj)(Handlers())))
}
