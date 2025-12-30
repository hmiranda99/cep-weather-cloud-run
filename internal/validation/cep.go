package validation

import "unicode"

func IsValidCEP(cep string) bool {
	if len(cep) != 8 {
		return false
	}
	for _, r := range cep {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
