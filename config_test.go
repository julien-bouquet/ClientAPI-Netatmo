package main

import (
	"os"
	"strings"
	"testing"
)

func TestGetFileName_succeedNoEnvSet(t *testing.T) {
	// Given
	_ = os.Unsetenv("ENV")
	// When
	result := getFileName()
	// Then
	if result != "properties.development" {
		t.Errorf("Error on Environment variable, %v", result)
	}
}

func TestGetFileName_succeedWithCustomEnv(t *testing.T) {
	// Given
	_ = os.Setenv("ENV", "azert")
	// When
	result := getFileName()
	// Then
	if result != "properties.azert" {
		t.Errorf("Error on Environment variable")
	}
}

func TestGetPathFile_succeed(t *testing.T) {
	result := getPathFolderOfFile()
	if !strings.Contains(result, "config") {
		t.Errorf("Error on file")
	}
}
