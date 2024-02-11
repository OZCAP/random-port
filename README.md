# random-port 🔌⚡️
A simple CLI tool written in Golang for generating random and safe port numbers to use for local development.

### Why? 🤔
While developing software, it can sometimes be a pain to choose a local port number which is not in use by another project or system process. My current workflow when starting a new project is to generate a random port number for all its services *(web interface, database, etc)* to ensure that there is no port sharing conflict.

### Installation 🎉
#### Install or update latest version (Mac/Linux):
```
curl -fsSL https://raw.githubusercontent.com/OZCAP/random-port/main/install.sh | bash
```
#### Uninstall:
```bash
sudo rm -f /usr/local/bin/random-port
```

### Usage
To pick a random port:
```bash
random-port generate
```

For help:
```bash
random-port --help
```
