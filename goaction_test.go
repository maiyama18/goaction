package goaction

import "testing"

func TestHello(t *testing.T) {
	want := "Don't communicate by sharing memory, share memory by communicating."
	got := Go()
	if want != got {
		t.Fatalf("test failed. want=%q, got=%q", want, got)
	}
}
