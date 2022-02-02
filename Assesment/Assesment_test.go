package main

import (
	controller "NisijAssesment/controller"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Pets struct {
	Id    int     `json:"id"`
	Type  string  `json:"type"`
	Price float32 `json:"price"`
}

func TestGetPets(t *testing.T) {
	router := controller.GetPets()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pets", nil)
	router.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
}
func TestPostPets(t *testing.T) {
	router := controller.PostPets()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/pets", nil)
	router.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
}
func TestGetPets(t *testing.T) {
	router := controller.GetPetByID()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pets/1", nil)
	router.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
}
