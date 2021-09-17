/**
Testa dina funktioner med koden i den här filen
*/

package texter_test

import (
	"testing"

	"example.com/go-hello-world/texter"
)

func TestTexter(t *testing.T) {
	if texter.GetGreating() != "Hello World!!" {
		t.Fatal("Wrong return text!!")
	}
}
