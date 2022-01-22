package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Max"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Max") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
