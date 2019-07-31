package ch4

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title:  "test",
		Year:   1987,
		Color:  true,
		Actors: []string{"hou"},
	},
}

func Movie1() {
	data, err := json.Marshal(movies)

	if err != nil {
		fmt.Println("err")
	}

	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "  ")
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("%s\n", data)
}
