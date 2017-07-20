package main

import (
	"fmt"
	//	"github.com/bitly/go-simplejson"
	"reflect"

	"github.com/verdverm/frisby"
)

func main() {
	fmt.Println("Using Frisby! for API Testing \n")
	//Test 1
	frisby.Create("Generate Wallet new seed").
		Get("http://127.0.0.1:6420//wallet/newSeed").
		Send().
		ExpectStatus(200).
		ExpectContent("seed")
		//ExpectJsonLength("seed", 0) type assertion to []interface{} failed

	//Test 2
	frisby.Create("Create Wallet").
		Post("http://127.0.0.1:6420/wallet/create").
		Send().
		ExpectStatus(200).
		ExpectContent("meta").
		ExpectJsonType("meta.coin", reflect.String).
		ExpectJsonType("meta.filename", reflect.String).
		ExpectJsonType("meta.label", reflect.String).
		ExpectJsonType("meta.lastSeed", reflect.String).
		ExpectJsonType("meta.seed", reflect.String).
		ExpectJsonType("meta.tm", reflect.String).
		ExpectJsonType("meta.type", reflect.String).
		ExpectJsonType("meta.version", reflect.String).
		// Expect(func(F *frisby.Frisby) (bool, string) {
		// 	frisby.Frisby.
		// 	return true, ""
		// }).
		// AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
		// 	val, _ := json.Get("meta").String()
		// 	fmt.Println("url =", val)
		// }).
		ExpectJsonLength("entries", 1)

	frisby.Global.PrintReport()
}
