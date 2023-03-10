package main

import (
	"os"
	"testing"
)

var done = make(chan struct{})
var logged = make(chan struct{})

func TestFunc1(t *testing.T) {
	t.Parallel()
	go func() {
		defer close(logged)
		<-done
		t.Log("log in go routine!")
	}()
	t.Log("log1?")
}

func TestMain(m *testing.M) {
	code := m.Run()
	close(done) // screw up TestFunc1
	<-logged
	os.Exit(code)
}
