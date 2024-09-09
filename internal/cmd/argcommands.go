package cmd

type ArgCommand struct {
	Name      string
	ShortDesc string
	LongDesc  string
}

func NewCommand(name string, shortDesc string, longDesc string) *ArgCommand {
	c := ArgCommand{Name: name, ShortDesc: shortDesc, LongDesc: longDesc}
	return &c
}

func HandleCommands() map[ArgCommand]func([]string) bool {
	newcmd := *NewCommand("new", "creates a new project directory", "creates a new project directory, Usage:\n\tcreate new <flag> where flag is the project name")
	buildcmd := *NewCommand("build", "builds the project", "builds the project and outputs to the bin folder")
	helpcmd := *NewCommand("help", "Displaying commands", "Displaying all commands")
	runcmd := *NewCommand("run", "Runs the project", "runs the project executable, requires build command to be used first")
	get_packagescmd := *NewCommand("get_packages", "Downloads the packages listed in the config.toml", "Downloads and installs the packages listed in the config.toml into special directories in the users home/.config/cppcreate/pkg directory")

	return map[ArgCommand]func([]string) bool{
		newcmd:          New,
		buildcmd:        Build,
		runcmd:          Run,
		helpcmd:         Help,
		get_packagescmd: Get_Packages,
	}
}
