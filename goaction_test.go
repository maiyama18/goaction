package goaction

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	got := Hello()
	if want != got {
		t.Fatalf("test failed")
	}
}

