package mcp

import (
	"testing"
)

func TestAnonymizeValue(t *testing.T) {
	// Test email anonymization
	emailResult := AnonymizeValue("email", "john.doe@example.com")
	if emailResult != "jo***@example.com" {
		t.Errorf("Expected 'jo***@example.com', got '%v'", emailResult)
	}

	// Test password anonymization
	pwdResult := AnonymizeValue("password_hash", "secret123")
	if pwdResult != "********" {
		t.Errorf("Expected '********', got '%v'", pwdResult)
	}

	// Test phone anonymization
	phoneResult := AnonymizeValue("phone", "5551234567")
	if phoneResult != "******4567" {
		t.Errorf("Expected '******4567', got '%v'", phoneResult)
	}
}
