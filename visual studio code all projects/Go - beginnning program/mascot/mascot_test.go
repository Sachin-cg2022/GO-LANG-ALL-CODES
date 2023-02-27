package mascot_test

import (
	"myapp/mascot"
	"testing"
)

func TestMascot(t *testing.T) {

	if mascot.BestMascot() != "Go Gofher" {
		t.Fatal("wrong mascot :(")
	}

}
