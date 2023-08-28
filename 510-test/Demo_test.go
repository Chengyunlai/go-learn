package demo

import (
	"log"
	"testing"
)

func TestInitJson(t *testing.T) {
	err := InitJson("./")
	if err != nil {
		log.Fatal(err)
	}
}
