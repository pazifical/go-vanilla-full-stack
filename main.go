package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type Person struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	Age          int    `json:"age"`
	Professional bool   `json:"professional"`
	SkillLevel   int    `json:"skill_level"`
	Gender       string `json:"gender"`
	Message      string `json:"message"`
	Color        string `json:"color"`
}

type IndexData struct {
	People []Person
}

var templates *template.Template
var peopleTable = createPeopleMap()
var newID = 0

func main() {
	t, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	templates = t

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("GET /api/person", getPersonTable)
	http.HandleFunc("POST /api/person", addPerson)

	err = http.ListenAndServe(":3337", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	data := IndexData{People: createPersonList()}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func getPersonTable(w http.ResponseWriter, r *http.Request) {
	people := createPersonList()

	err := templates.ExecuteTemplate(w, "person_table.html", people)
	if err != nil {
		http.Error(w, "", http.StatusUnprocessableEntity)
		return
	}
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "", http.StatusUnprocessableEntity)
		return
	}

	peopleTable[newID] = person
	newID += 1

	http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
}

func createPersonList() []Person {
	people := make([]Person, 0)
	for _, person := range peopleTable {
		people = append(people, person)
	}
	return people
}

func createPeopleMap() map[int]Person {
	people := []Person{
		{
			Name:         "Alice",
			Age:          30,
			Gender:       "female",
			Password:     "ABC123",
			Professional: true,
			SkillLevel:   5,
			Message:      "I am a programmer",
			Color:        "blue",
		},
		{
			Name:         "Bob",
			Age:          40,
			Gender:       "male",
			Password:     "admin",
			Professional: false,
			SkillLevel:   1,
			Message:      "I am a janitor",
			Color:        "red",
		},
	}

	peopleMap := make(map[int]Person, 0)
	for _, person := range people {
		peopleMap[newID] = person
		newID += 1
	}
	return peopleMap
}
