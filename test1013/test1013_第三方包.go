package main

//go get github.com/tidwall/gjson
//或者go mod tity
import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())
	fmt.Println(value.String())
}
