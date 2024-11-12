package tests

import (
	"fmt"
	"testing"
  "github.com/irlalan/gcreate/internal/dir"
)

func assertneq(lexp any, rexp any) bool {
	return lexp != rexp
}

func TestCheck(t *testing.T) {
	result := dir.Check(fmt.Errorf("Testing"), "for testing")
	expected := false
	if result != expected {
		t.Fatalf("Failed to correctly check error, result: %t, expected: %t", result, expected)
	}
	fmt.Println("TestCheck Succeded")
}

func TestCreateFile(t *testing.T) {
	filename := "test.txt"
	fileWrite := "hey there"
	result := dir.CreateFile(filename, fileWrite)
	expected := true
	if result != expected {
		t.Fatalf("Failed to create file, result: %t, expected: %t", result, expected)
	}
	fmt.Println("TestCreateFile Succeded")
}

func TestCreateDir(t *testing.T) {
	filename := "test.txt"
	result := dir.CreateDir(filename)
	expected := true
	if result != expected {
		t.Fatalf("Failed to create directory %s, result: %t, expected: %t", filename, result, expected)
	}
	fmt.Println("TestCreateDir Succeded")
}

func TestReadTemplateFile(t *testing.T) {
	filename := "./test.txt"
	result := dir.ReadTemplateFile(filename, "test")
	expected := "hey there"
	if assertneq(result, expected) {
		t.Fatalf("Failed to create directory %s, result: %s, expected: %s", filename, result, expected)
	}
	fmt.Println("TestReadTemplateFile Succeded")
}

func TestGetHomeDir(t *testing.T) {
	expected := "/home/irlalan"
	result := dir.GetHomeDir()
	if assertneq(result, expected) {
		t.Fatalf("Failed to query home directory, result: %s, expected: %s", result, expected)
	}
	fmt.Println("TestGetHomeDir succeded; ", result)
}
