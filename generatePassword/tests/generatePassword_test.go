package generatepassword_test

import (
	"testing"
)

// test downcase function
func TestDowncase(t *testing.T) {
	// test data
	text := "ASFA"
	// expected result
	expected := "asfa"
	// call function
	result := downcase(text)
	// check result
	if result != expected {
		t.Errorf("downcase(%s) = %s; want %s", text, result, expected)
	}
}
