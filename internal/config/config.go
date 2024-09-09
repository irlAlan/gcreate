package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	// "strings"

	"github.com/BurntSushi/toml"
	"github.com/irlalan/gcreate/internal/dir"
)

// TODO: be able to convert to json compile commands

type General struct {
	Name         string   `toml:"project_name"`
	Lang_version string   `toml:"lang_version"`
	Bin_dir      string   `toml:"bin_dir"`
	Src_files    []string `toml:"src_files"`
  Header_files []string `toml:"header_files"`
	Compiler     string   `toml:"compiler"`
	Build_type   string   `toml:"build_type"`
	Flags        []string `toml:"flags"`
  Linker_flags []string `toml:"linker_flags"`
  Lib_dir      string   `toml:"lib_dir"`
}

// 18/08/24 focus on getting the url and downloading the packages
type Pkg struct {
	Name         string   `toml:"name"`
	GitUrl       string   `toml:"git_url"`
	Tag          string   `toml:"tag"`
	Branch       string   `toml:"branch"`
	Dependancies []string `toml:"dependencies"`
}

type TConfig struct {
	Conf General `toml:"conf"`
	Pkgs []Pkg   `toml:"pkg"`
}

func check(err error, detail string) bool {
	if err != nil {
		log.Println(detail, err)
		return false
	}
	return true
}

func MarshalConfig(data TConfig) string {
	marsh_text, err := toml.Marshal(data)
	if err != nil {
		fmt.Println("Failed to marshal struct: ", err)
	}
	return string(marsh_text)
}

func GetDefaultConfig(filepath string, projname string) string {
	var conf TConfig
	conf.Conf.Name = projname
	conf.Conf.Src_files = []string{"src/main.cpp"}
	conf.Conf.Lang_version = "17"
	conf.Conf.Bin_dir = "bin/"
	conf.Conf.Compiler = "g++"
	conf.Conf.Build_type = "Debug"

	return MarshalConfig(conf)
}

func ReadConfig() TConfig {
	filepath := "./config.toml"
	data, err := os.ReadFile(filepath)
	check(err, "Cannot Read file: ")

	var conf TConfig
	md, err := toml.Decode(string(data), &conf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Undecoded keys: %q\n", md.Undecoded())

	if !md.IsDefined("conf", "src_files") {
		fmt.Printf("conf.src_files is not defined, using default: src/main.cpp\n")
		conf.Conf.Src_files = []string{"src/main.cpp"}
	}
	if !md.IsDefined("conf", "bin_dir") {
		fmt.Printf("conf.bin_dir is not defined, using defualt: bin/default\n")
		conf.Conf.Bin_dir = "bin/default"
	}
	if !md.IsDefined("conf", "lang_version") {
		fmt.Printf("conf.lang_version is not defined, using defualt: 17\n")
		conf.Conf.Lang_version = "17"
	}
	if !md.IsDefined("conf", "project_name") {
		fmt.Printf("conf.project_name is not defined, using defualt: default\n")
		conf.Conf.Bin_dir = "default"
	}
	if !md.IsDefined("conf", "compiler") {
		fmt.Printf("compiler is not defined, using defualt: g++\n")
		conf.Conf.Bin_dir = "g++"
	}
	if !md.IsDefined("conf", "build_type") {
		fmt.Printf("build_type is not defined, using defualt: Debug\n")
		conf.Conf.Build_type = "Debug"
	}

	// fmt.Println(conf)
	// fmt.Println("src_files: ", md.IsDefined("conf.src_files"))
	// fmt.Printf("Name: %s\nLang: %s\nBinary: %s\nSrc: %s\n", conf.Conf.Name,conf.Conf.Lang_version, conf.Conf.Bin_dir, conf.Conf.Src_files);
	return conf
}

type CompileCommandsJson struct{
  Directory string `json:"directory,omitempty"`
  File string `json:"file,omitempty"`
  Output string `json:"output,omitempty"`
}

func (cc CompileCommandsJson) Compile_commands(conf TConfig) string{

  var cc_file_arr []CompileCommandsJson

  for _, src_file := range conf.Conf.Src_files{
    cc.Directory = dir.GetCurrentDir()
    cc.File = src_file 
    cc.Output = conf.Conf.Bin_dir+dir.GetObjFile(src_file)
    cc_file_arr = append(cc_file_arr, cc)
  }

  var final_str []string
  var tmp_str []byte 
  var err error
  for _,val := range cc_file_arr{
    tmp_str, err = json.Marshal(&val)
    if err != nil{
      log.Fatal("Error: ", err)
    }
    final_str = append(final_str, string(tmp_str), "\n")
  }

  return strings.Join(final_str, ",")
}
