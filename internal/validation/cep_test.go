package validation

import "testing"

func TestIsValidCEP(t *testing.T) {
	if !IsValidCEP("01001000") {
		t.Fatal("expected valid CEP")
	}
	if IsValidCEP("123") {
		t.Fatal("expected invalid CEP")
	}
}
