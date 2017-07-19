package main

import (
	"fmt"
//	"reflect"

//	"github.com/bitly/go-simplejson"
	"github.com/verdverm/frisby"
)
func main() {
	fmt.Println("Using Frisby! for API Testing \n")

	frisby.Create("Test GET Skycoin homepage").
		Get("http://127.0.0.1:6420/wallet/newSeed").
		Send().
		ExpectStatus(200).
		ExpectContent("seed")
frisby.Global.PrintReport()
}
