package loginlog

import (
	"log"
	"testing"
)

func TestList(t *testing.T) {
	loginLog, count, err := List(0, 20)
	if err != nil {
		t.Fatal(err)
	}

	log.Println("count:", count)
	log.Print("log:", loginLog)
}
