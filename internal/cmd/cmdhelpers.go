package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/irlalan/gcreate/internal/config"
	"github.com/irlalan/gcreate/internal/dir"
)

func build_file(conf config.TConfig, file string) bool{

	var cmd *exec.Cmd
	var err error
	var stderr bytes.Buffer
	var out bytes.Buffer

	fileobj := dir.GetObjFile(file)
	fmt.Printf("building file: %s to -> %s\n", file, conf.Conf.Bin_dir+fileobj)
	// builds to a *.o file
  var args []string
  var tmp_args = []string{"-c", file, "-o", conf.Conf.Bin_dir+fileobj}
  // tmp_args = append(tmp_args,"-o", conf.Conf.Bin_dir+fileobj)
  args = append(args, conf.Conf.Flags...)
  args = append(args, tmp_args...)

	cmd = exec.Command(conf.Conf.Compiler, args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	} else {
		fmt.Println("Result: " + out.String())
		return true
	}
}


func gobuild_file(conf config.TConfig, file string, ret chan bool) {

	var cmd *exec.Cmd
	var err error
	var stderr bytes.Buffer
	var out bytes.Buffer

	fileobj := dir.GetObjFile(file)
	fmt.Printf("building file: %s to -> %s\n", file, conf.Conf.Bin_dir+fileobj)
  var args []string
  var obj_args = []string{"-c", file, "-o", conf.Conf.Bin_dir+fileobj}
  args = append(args, conf.Conf.Flags...) // check if value has been changed
  args = append(args, obj_args...)
	// builds to a *.o file
	cmd = exec.Command(conf.Conf.Compiler,args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		ret <- false
	} else {
		fmt.Println("Result: " + out.String())
		ret <- true
	}
}

func build_src(conf config.TConfig, files []string) bool {
	fmt.Println("Building source")
	// return source file that coudnt be built
	ret := make(chan bool)
	for _, file := range files {
		go gobuild_file(conf,file, ret)
	}
	time.Sleep(time.Second * 1)
	val := <-ret
	return val
}

func link_obj(conf config.TConfig, obj_files []string) bool {
	var cmd exec.Cmd
	var err error
	var stderr bytes.Buffer
	var out bytes.Buffer

	fmt.Printf("linking obj files: %s to -> %s\n", obj_files, conf.Conf.Bin_dir+conf.Conf.Name)
	// builds to a *.o file
  // cmd = exec.Command(compiler, "-o", obj_files)

  var tmp_linkflag []string

  if len(conf.Conf.Lib_dir) != 0{
    tmp_linkflag = append(tmp_linkflag, "-L"+conf.Conf.Lib_dir)
  }

  if len(conf.Conf.Linker_flags) != 0{
    fmt.Println("linker flag defined")
    for _, link_flag := range conf.Conf.Linker_flags{
      tmp_linkflag = append(tmp_linkflag, "-l"+link_flag)
    }
    cmd.Args = append(cmd.Args, tmp_linkflag...)
    fmt.Printf("linker flags: %#v\n", tmp_linkflag)
  }

  cmd.Args = append(cmd.Args, "-o", conf.Conf.Bin_dir+conf.Conf.Name)
  cmd.Args = append(cmd.Args, obj_files...)

  fmt.Printf("cmd args: %#v\n", cmd.Args)

  cmd = *exec.Command(conf.Conf.Compiler,cmd.Args...)

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	} else {
		fmt.Println("Result: " + out.String())
		return true
	}
}
