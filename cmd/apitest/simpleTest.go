package main

import (
	"fmt"
	//	"github.com/bitly/go-simplejson"
	"reflect"

	simplejson "github.com/bitly/go-simplejson"
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
		AfterJson(func(F *frisby.Frisby, json *simplejson.Json, err error) {
			entries := json.Get("entries").GetIndex(0)
			addressType := entries.Get("address").Interface()
			publicKeyType := entries.Get("public_key").Interface()
			secretKeyType := entries.Get("secret_key").Interface()
			if reflect.ValueOf(addressType).Kind() != reflect.String {
				errStr := fmt.Sprintf("Expect Json %q type to be %q, but got %T", "entries.address", reflect.String, addressType)
				F.AddError(errStr)
			}
			if reflect.ValueOf(publicKeyType).Kind() != reflect.String {
				errStr := fmt.Sprintf("Expect Json %q type to be %q, but got %T", "entries.public_key", reflect.String, publicKeyType)
				F.AddError(errStr)
			}
			if reflect.ValueOf(secretKeyType).Kind() != reflect.String {
				errStr := fmt.Sprintf("Expect Json %q type to be %q, but got %T", "entries.secret_key", reflect.String, secretKeyType)
				F.AddError(errStr)
			}
		}).
		ExpectJsonLength("entries", 1)

	frisby.Global.PrintReport()
}
