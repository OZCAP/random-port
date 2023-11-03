# random-port 🔌⚡️
A simple CLI tool written in Golang for generating random and safe port numbers to use for local development.

### But why? 🤔
When building software locally, it can sometimes be a pain to choose a port number which is not in use by another project or system process. My current workflow when starting a new project is to generate a random port number for all its services *(web interface, database, etc)* to ensure that there is no port sharing conflict.

### Installation 🎉
#### Mac OSX (Homebrew):
```bash
brew tap ozcap/random-port && brew install random-port
```
#### Manual installation (build from source):
1. Ensure [go is installed globally](https://go.dev/doc/install)
2. Clone this repository and `cd` into it
3. Run the `install.sh` script to build and install the binary


### Usage
To pick a random port:
```bash
random-port generate
```

For help screen:
```bash
random-port --help
```
