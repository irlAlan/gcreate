# cppcreate

The idea of this project is to make a project manager and creator for C++ in golang.

The main functions of cppcreate will be to create new projects and build/run them; it should be able to read (mostly, hopefully) with cmake and conan2.

# Install

To build the program call: `make build` you can then use it in the bin folder as an executable

To install the program call: `make install` and it will be added to path with the templates folder in your config/cppcreate folder

# Cmdline Usage

For help, `cppcreate help` or `cppcreate`

To create a new project run: `cppcreate new <project_name>`

To build a project go to the root directory with config.toml file and run: `cppcreate build`

To run a project go to the root directory with config.toml file and run: `cppcreate run`

To download the packages go to the root directory with config.toml file and run `cppcreate get_packages`

# Config.toml

# Tests
All the tests are located in the `tests/` directory.

To run all tests navigate to the test/ directory and use the command: `go test`

To run individual tests use the command `go test -run <test func name>`

# TODO:
- [X] Create Marshler so I can correctly output default config file
- [ ] fix current tests
- [ ] Refactor code & add more tests
- [X] Check if file has been changed and if it has re-build only that one
- [ ] Create custom logs i.e. output to file
- [X] add support for specific flags i.e. compile_commands.json
- [ ] serialise into make files
- [ ] handle calls to github for package download
  - [X] clone git repos in packages
  - [ ] clone using tags/branches
