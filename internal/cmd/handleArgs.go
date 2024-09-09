package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	// "strings"

	// 	"strings"
	"errors"
	// "github.com/fsnotify/fsnotify"
	"github.com/irlalan/gcreate/internal/config"
	"github.com/irlalan/gcreate/internal/dir"
	"github.com/irlalan/gcreate/internal/net"
)

func HandleArgs(args []string) bool {

	cmds := HandleCommands()

	for cmd, cmdfunc := range cmds {
		if len(args) < 1 {
			Help(args)
		} else if args[0] == cmd.Name {
			cmdfunc(args[1:])
		}
	}

	fmt.Println(args)
	return true
}

func New(args []string) bool {
	projname := args[0]
	maindir := filepath.Join("./", projname)
	srcdir := filepath.Join(maindir, "/src")
	bindir := filepath.Join(maindir, "/bin")
	if !dir.CreateDir(maindir) {
		return false
	}
	if !dir.CreateDir(srcdir) {
		return false
	}
	if !dir.CreateDir(bindir) {
		return false
	}
	if !dir.CreateFile(srcdir+"/main.cpp", dir.ReadTemplateFile(dir.GetTemplateDir()+"/main.cpp", projname)) {
		return false
	}

	if !dir.CreateFile(maindir+"/config.toml", config.GetDefaultConfig(dir.GetTemplateDir()+"/config.toml", projname)) {
		return false
	}
	return true
}

func Build(args []string) bool {
	// TODO: Build source files seperatly
	// Serialise to compile_commands.json
	conf := config.ReadConfig()
	//var src string
	//for _, data := range conf.Conf.Src_files {
	//	src += data + " "
	//}

	// TODO: check if a file has been changed
	for _, file := range conf.Conf.Src_files {
		binfileinfo, binerr := os.Stat(conf.Conf.Bin_dir + dir.GetObjFile(file))
		srcfileinfo, _ := os.Stat(file)
		if errors.Is(binerr, os.ErrNotExist) {
			// Doesnt exist so build
			fmt.Println("binary file does not exist")
			if !build_file(conf, file) {
				return false
			}
		} else {
			// exists so check modification date then build if needs to
			fmt.Println("File exists")
			binh, binm, bins := binfileinfo.ModTime().Clock()
			srch, srcm, srcs := srcfileinfo.ModTime().Clock()
			if binh >= srch && binm >= srcm && bins >= srcs {
				// fmt.Printf("file: %s Hasnt changed\n\t->src: h: %d, m: %d, s: %d; ->bin: h: %d, m: %d, s: %d\n", file, srch, srcm, srcs, binh,binm, bins)
				fmt.Printf("file: %s Hasnt changed\n", file)
			} else {
				// fmt.Printf("file: %s changed\n\t->src: h: %d, m: %d, s: %d; ->bin: h: %d, m: %d, s: %d\n", file, srch, srcm,srcs, binh,binm, bins)
				fmt.Printf("file: %s changed\n", file)
				if !build_file(conf, file) {
					return false
				}
			}
		}
	}
  // TODO: add flags to show/disables created obj files maybe
	// linking obj files
  var obj_files []string
  for _,srcfile := range conf.Conf.Src_files{
    obj_files = append(obj_files, conf.Conf.Bin_dir+ dir.GetObjFile(srcfile))
  }
  // be able to link  other dependancies and stuff
  fmt.Println("Started linking files now")
  if !link_obj(conf,obj_files){
    fmt.Println("Cannot link the files")
    return false
  }

  var compile_cmd config.CompileCommandsJson
  val := compile_cmd.Compile_commands(conf)
  dir.CreateFile("compile_commands.json", val)

	//if build_src(conf.Conf.Compiler, conf.Conf.Src_files, conf.Conf.Bin_dir) {
	//	fmt.Println("Finished building source files")
	//} else {
	//	fmt.Println("unable to build source files")
	//}

	//	var val exec.Cmd
	//
	//	val.Path = conf.Conf.Compiler
	//	fmt.Println("Path: ", val.Path)
	//	for i, val := range conf.Conf.Src_files {
	//		conf.Conf.Src_files[i] = strings.Trim(val, " ")
	//	}
	//	val.Args = append(val.Args, conf.Conf.Flags...)
	//	val.Args = append(val.Args, "-std=c++"+conf.Conf.Lang_version)
	//	val.Args = append(val.Args, conf.Conf.Src_files...)
	//	val.Args = append(val.Args, "-o", conf.Conf.Bin_dir+"/"+conf.Conf.Name)
	//	fmt.Println("Args: ", val.Args)
	//
	//	// val.Args[0] = "-o"
	//	// val.Args = conf.Conf.Src_files
	//	// val.Run()
	//	cmd := exec.Command(val.Path, val.Args...)
	//	// cmd := exec.Command(conf.Conf.Compiler,src,"-o", strings.Trim(conf.Conf.Bin_dir," ")+strings.Trim(conf.Conf.Name, " "));
	//	// cmd := exec.Command("g++", "--version");
	//	var stderr bytes.Buffer
	//	var out bytes.Buffer
	//	val.Stdout = &out
	//	val.Stderr = &stderr
	//	err := cmd.Run() // using cmd cause val.Path() not working
	//	if err != nil {
	//		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	//		return false
	//	}
	//	fmt.Println("Result: " + out.String())
	return true
}

func Run(args []string) bool {
	projname := config.ReadConfig().Conf.Name
	var cmd exec.Cmd
	cmd.Path = "./bin/" + projname
	cmd.Args = append(cmd.Args, projname)
	// cmd := exec.Command("cd", "./bin/", "./"+projname)
	var stderr bytes.Buffer
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	}
	fmt.Println("Result: " + out.String())
	return true
}

func Get_Packages(args []string) bool {
	// checks if the home/.config/cppcreate/pkgs exists
	// creates it if not
	path := dir.GetHomeDir() + "/.config/cppcreate/pkgs/"
	isDirMade := dir.CreateDir(path)
	if isDirMade {
		fmt.Printf("package path: %s created\n", path)
	}
	// reads the config for the [[pkgs]] tag and if none exits and complains
	conf := config.ReadConfig()
	// if there are packages, makes a directory in the home/.config/cppcreate/pkgs
	//  -> of the name+tag/branch/version of the pkg
	for _, pkg := range conf.Pkgs {
		fmt.Printf("name: %s,\n-->  git_url: %s\n  -->  tag: %s\n", pkg.Name, pkg.GitUrl, pkg.Tag)
	}
	net.HandlePkgs(conf.Pkgs)
	// time.Sleep(time.Second * 1)
	// downlods the pkg here then compiles/builds the pkg into sub directories
	//  -> that we will link to through compiler
	return true
}

func Help(args []string) bool {
	fmt.Println("Displaying Help -->")
	cmd_desc := HandleCommands()
	for key, _ := range cmd_desc {
		fmt.Printf("%s, short description: %s\n\t-> long description: %s\n", key.Name, key.ShortDesc, key.LongDesc)
	}
	return true
}
