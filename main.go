package main


import (
  "github.com/irlalan/gcreate/internal/cmd"
  "os"
)


func main(){
  args := os.Args[1:];
  cmd.HandleArgs(args);
}
