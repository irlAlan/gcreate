package cmd

import (
	"github.com/irlalan/gcreate/internal/config"
	"github.com/irlalan/gcreate/internal/dir"
	"log"
	"os/exec"
)

/*
g++ -Wall -Werror -cmain.cpp -obin/main.o
compiler, []flags src_file obj_file
*/

// Builds source files into obj
type SrcBuild struct {
	Compiler string
	Flags    []string
	ObjCmd   string
	SrcFile  string
	ObjFile  string
}

type SrcBuilder interface {
	SetCompiler(string) SrcBuilder
	SetFlags([]string) SrcBuilder
	SetSrcFile(string) SrcBuilder
	SetObjCmd(string) SrcBuilder
	SetObjFile(string, string) SrcBuilder
	Build() *exec.Cmd
}

type ConcereteSrcBuilder struct {
	FinishedSrcBuild *SrcBuild
}

func NewConcreteSrcBuilder() *ConcereteSrcBuilder {
	return &ConcereteSrcBuilder{FinishedSrcBuild: &SrcBuild{}}
}

func (csb *ConcereteSrcBuilder) SetCompiler(compiler string) SrcBuilder {
	comp, err := exec.LookPath(compiler)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	csb.FinishedSrcBuild.Compiler = comp
	return csb
}

func (csb *ConcereteSrcBuilder) SetFlags(flags []string) SrcBuilder {
	csb.FinishedSrcBuild.Flags = flags
	return csb
}
func (csb *ConcereteSrcBuilder) SetSrcFile(src_file string) SrcBuilder {
	csb.FinishedSrcBuild.SrcFile = src_file
	return csb
}
func (csb *ConcereteSrcBuilder) SetObjCmd(cmd string) SrcBuilder {
	csb.FinishedSrcBuild.ObjCmd = cmd
	return csb
}

func (csb *ConcereteSrcBuilder) SetObjFile(bin_dir string, obj_file string) SrcBuilder {
	csb.FinishedSrcBuild.ObjFile = "-o" + bin_dir + (dir.GetObjFile(obj_file))
	return csb
}

func (csb *ConcereteSrcBuilder) Build() *exec.Cmd {
	var cmd exec.Cmd
	// Possibly nil dereferencing
	cmd.Path = csb.FinishedSrcBuild.Compiler
	// dont know why but appending an empty string before everything stops g++ error
	// thinking the next option passed is a file or something
	// other option is having the user put an empty string in the flags array
	cmd.Args = append(cmd.Args, "")
	cmd.Args = append(cmd.Args, csb.FinishedSrcBuild.Flags...)
	cmd.Args = append(cmd.Args, csb.FinishedSrcBuild.ObjCmd)
	cmd.Args = append(cmd.Args, csb.FinishedSrcBuild.SrcFile)
	cmd.Args = append(cmd.Args, csb.FinishedSrcBuild.ObjFile)
	return &cmd
}

type Director struct {
	builder SrcBuilder
}

func NewDirector(builder SrcBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct(conf config.TConfig, objcmd string, srcfile string) *exec.Cmd {
	d.builder.SetCompiler(conf.Conf.Compiler).SetFlags(conf.Conf.Flags).SetObjCmd(objcmd).SetSrcFile(srcfile).SetObjFile(conf.Conf.Bin_dir, srcfile)
	return d.builder.Build()
}

// will probs do something more with this if linking gets more complex
// func (d *Director) ConstructLink(conf config.TConfig, obj_files []string) *exec.Cmd {
// 	return d.builder.Build()
// }
