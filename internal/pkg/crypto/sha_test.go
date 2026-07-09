package crypto

import (
	"testing"
)

// TestGenerateSHA generates the sha from given data and verifies whether it is correct or not
func TestGenerateSHA(t *testing.T) {
	data := "www.stakater.com"
	sha := "bfe96a7c834576eef1b9a307159b607f0edceef48fc92d3db563ec0a07a90082"
	result := GenerateSHA(data)
	if result != sha {
		t.Errorf("Failed to generate SHA")
	}
}

// TestGenerateSHAEmptyString verifies that empty string generates a valid hash
// This ensures consistent behavior and avoids issues with string matching operations
func TestGenerateSHAEmptyString(t *testing.T) {
	result := GenerateSHA("")
	expected := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	if result != expected {
		t.Errorf("Failed to generate SHA for empty string. Expected: %s, Got: %s", expected, result)
	}
	if len(result) != 64 {
		t.Errorf("SHA hash should be 64 characters long, got %d", len(result))
	}
}
