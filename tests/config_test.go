package tests

import (
	"fmt"
	"testing"

	"github.com/irlalan/gcreate/internal/config"
)

func TestHandleConfig(t *testing.T) {
	result := config.GetDefaultConfig("test.txt", "test")
	expected := "done"
	if result != expected {
		t.Fatalf("not working")
	}
	fmt.Println("is working")
}

// TODO:Fix this test
func TestFormatMapString(t *testing.T) {
	test := map[string]any{
		"name": "usr1",
	}
  conf := config.TConfig{}
	result := config.MarshalConfig(conf)
	var expected string
	for cmd, info := range test {
		expected += (cmd + "=" + info.(string)) + "\n"
	}

	if assertneq(result, expected) {
		t.Fatalf("Format string not working. \nexpected: %s\nbut got: %s\n", expected, result)
	}
	fmt.Println("TestFormatMapString succeded")
}
