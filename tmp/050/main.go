package main

import (
	"fmt"
	"github.com/integrii/flaggy"
	"github.com/tauinger-de/training.go-basics.lib/compare"
)

func main() {
	var algFlag = "sha256"
	var inputArg = ""
	flaggy.SetName("gohash")
	flaggy.SetVersion("0.1")
	flaggy.String(&algFlag, "a", "alg", "The hashing algorithm to use")
	flaggy.AddPositionalValue(&inputArg, "input", 1, true, "the input to hash (file or URL)")
	flaggy.Parse()

	fmt.Println(inputArg)

	fmt.Println(compare.EqualInt(1, 2))
}