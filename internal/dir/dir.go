package dir

import (
	"fmt"
	"log"
	"os"
	// "path/filepath"
	"strings"
)

const TTemplateDir string = "/templates"

func GetTemplateDir() string {
	return GetHomeDir() + "/.config/cppcreate/templates"
}

func GetHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname
}

func GetHomeConfigDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return (dirname + "/.config/cppcreate")
}

func Check(err error, detail string) bool {
	if err != nil {
		log.Println(detail, err)
		return false
	}
	return true
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateFile(name string, input string) bool {
	file, err := os.Create(name)
	defer file.Close()
	Check(err, "Cannot create file: ")
	_, err = file.WriteString(input)
	file.Sync()
	Check(err, "Cannot write to file: ")
	return true
}

func CreateDir(path string) bool {
	exist := Exists(path)
	if exist {
		fmt.Println("Cannot create directory: it already exists")
		return false
	}
	if !exist {
		err := os.Mkdir(path, os.ModePerm)
		return Check(err, "Cannot create directory")
	}
	return false
	// err := os.Mkdir(path, os.ModePerm)
	// return Check(err, "Cannot create directory")
}

func ReadTemplateFile(filepath string, projname string) string {
	var datab []byte
	var datas string
	var err error

	if strings.Contains(filepath, "config.toml") {
		// handle config file differently
		// datas = config.GetDefaultConfig(filepath, projname)
	} else {
		datab, err = os.ReadFile(filepath)
		datas = string(datab)
		Check(err, "Cannot Read file: ")
	}
	return datas
}

func GetCurrentDir() string {
  cwd, err := os.Getwd()
  if err != nil{
    log.Fatalf("Error: %#v\n", err)
  }
  return cwd 
}


func GetObjFile(file string) string {
	// get the object file name and extension
	split := strings.Split(file, ".")
	split[1] = "o"
	fileobj := strings.Join(split, ".")
	fileobj_split := strings.Split(fileobj, "/")
	for _, val := range fileobj_split {
		if strings.Contains(val, ".o") {
			fileobj = val
		}
	}
	return fileobj
}
