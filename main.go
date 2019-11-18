package main

import (
	"github.com/taisa831/sandbox-go-impl/json"
	"log"
)

func main() {
	newB, _ := json.CreateJson()
	log.Printf(string(newB))

	json.PrintJson(newB)
	newBB, _ := json.UpdateJson(json.B)
	json.PrintJson(newBB)
}

