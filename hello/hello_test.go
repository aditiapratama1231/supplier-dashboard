package hello

import "testing"

func TestCreateProduct(t *testing.T) {
	helloRefactory := Name("Refactory")
	if helloRefactory != "Hello, Refactory!" {
		t.Errorf("Invalid message %s", helloRefactory)
	}
}
