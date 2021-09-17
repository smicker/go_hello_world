/**
Programmet startar en web-server med en sida där den skriver ut "Hello World". Starta programmet
och surfa till localhost:8080 för att se resultatet.

För att få en executable i Linux så bygger du detta i Linux med
	"go build main.go"
Du kan också bygga en exe fil till Windows från linux (cross compilation) genom att
istället skriva:
	"GOOS=windows GOARCH=amd64 go build main.go"
Om du bara vill testköra det så skriver du
	"go run main.go".
*/

package main

import (
	"fmt"
	"net/http"

	"example.com/go-hello-world/texter"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, texter.GetGreating())
}

func main() {
	fmt.Println("Starting a web server on localhost:8080")

	// Source: https://sv.wikipedia.org/wiki/Go_(programspr%C3%A5k)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
