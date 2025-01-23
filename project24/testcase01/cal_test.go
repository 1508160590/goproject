package main

import (
	"testing"
)

func TestAddupper(t *testing.T) {
	res := addupper(10)
	if res != 55 {
		t.Fatalf("addupper(10) should be 55, but got %d", res)
	}
	t.Logf("addupper(10) is correct")
}
