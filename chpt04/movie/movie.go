package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"` // `json:"released"` - дескриптор полей
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	// Преобразование структуры данных Go наподобие movies в JSON называется маршачингом (marshaling).
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Сбой маршалинга JSON: %s", err)
	}

	fmt.Printf("Marshal:\n%s\n", data)

	data, err = json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("Сбой маршалинга JSON: %s", err)
	}
	fmt.Printf("Marshal intend:\n%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Сбой маршалинга JSON: %s", err)
	}
	fmt.Printf("Unmarshal:\n%s\n", titles)
}
