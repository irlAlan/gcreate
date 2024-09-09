package tests

import (
	"fmt"
	"testing"

	"github.com/irlalan/gcreate/internal/cmd"
)

func TestHandleArgs(t *testing.T){ 
  args := []string{"hey", "ther"};
  result := cmd.HandleArgs(args)
  expected := true;
  if result != expected{
    t.Errorf("Result was incorrect, got %t, want %t", result, expected);
  } else {
    fmt.Println("TestHandleArgs was successfull");
  }
}

func TestNewCommand(t *testing.T){
  result := cmd.NewCommand("test","for testing purposes", "long desc for testing");
  expected := cmd.ArgCommand{Name: "test",ShortDesc: "for testing purposes", LongDesc:"long desc for testing"};

  if (result.Name != expected.Name) && (result.ShortDesc != expected.ShortDesc) && (result.LongDesc != expected.LongDesc){
    t.Errorf("Result was incorrect, got %+v, want %+v", *result, expected);
  }
  fmt.Println("TestNewCommand was successfull");
}

func TestHandleCommands(t *testing.T){
  expected := *cmd.NewCommand("new","creates a new project directory", "creates a new project directory, Usage:\n\tcreate new <flag> where flag is the project name");
  result := cmd.HandleCommands();
  
  keys := make([]cmd.ArgCommand, 0, len(result))
  for c := range result {
    keys = append(keys, c)
  }
  if keys[0].Name != expected.Name {
    t.Errorf("Result was incorrect, got %+v, want %+v", result, expected);
  }
  fmt.Println("TestNewCommand was successfull");
}

func TestNewFunc(t *testing.T){ 
  args := []string{"hey", "ther"};
  result := cmd.New(args);
  expected := true;
  if !result{
    t.Errorf("Directory creation failed, got %t, want %t", result, expected);
  } 
  fmt.Println("TestNewFunc was successfull");
}
