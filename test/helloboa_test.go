package main

import (
	"testing"

	"github.com/motomux/hellobop"
)

func TestHello(t *testing.T) {
	if hellobop.Hello("Go") != "Hello Go from binary-only package!" {
		t.Errorf("Wrong output")
	}
}
