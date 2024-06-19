package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type R struct {
	Downloads []Data `json:"downloads"`
}

type Data struct {
	Downloads int `json:"downloads"`
}

func main() {
	var r R
	b, err := os.ReadFile("./metadata.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, d := range r.Downloads {
		sum += d.Downloads
	}
	fmt.Println(sum)
}
