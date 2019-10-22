package main

import (
	"os"
	"testing"
)

func TestIntegration_GetConfig_succeed(t *testing.T) {
	// Given
	_ = os.Setenv("ENV", "test")
	// When
	result := GetConfig()
	// Then
	if result.Auth.ClientID != "test_clientID" || result.Auth.ClientSecret != "test_clientSecret" ||
		result.Auth.Username != "test_login" || result.Auth.Password != "test_password" {
		t.Errorf("Error on Unmarshal")
	}

	if result.Api.Core != "http://localhost/" && result.Api.Auth != "token" {
		t.Errorf("Error on Unmarshal")
	}

}

func TestIntegration_GetConfig_failMissingFileConfig(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// Given
	_ = os.Setenv("ENV", "azerty")
	// When
	GetConfig()

}
