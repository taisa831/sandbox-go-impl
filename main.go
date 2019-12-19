package main

import (
	"fmt"
	"github.com/taisa831/sandbox-go-impl/car"
)

func main() {
	//newB, _ := json.CreateJson()
	//log.Printf(string(newB))
	//
	//json.PrintJson(newB)
	//newBB, _ := json.UpdateJson(json.B)
	//json.PrintJson(newBB)

	subaru := &car.Subaru{}
	//subaru := car.NewSubaru()
	ret := subaru.Run()
	fmt.Print(ret)
}
