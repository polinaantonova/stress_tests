package inMemory

import (
	"encoding/json"
	"net/http"
	"postgres/internal/person"
	"strconv"
)

type InMemory struct {
	dict map[int]*person.Person
}

func NewInMemory(dict map[int]*person.Person) *InMemory {
	return &InMemory{
		dict: dict,
	}
}

func (iM *InMemory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	person, exists := iM.dict[id]

	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
