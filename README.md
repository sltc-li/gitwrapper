# gitwrapper

### Prerequisite
* Go 1.16

### Usage
`go install github.com/li-go/gitwrapper/cmd/gw@latest`

```
$ gw -h
NAME:
   gw - a simple wrapper command for git

USAGE:
   gw [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   update, u   update current branch
   create, c   create a new refresh branch
   commit, ci  commit changes
   push, p     push to remote
   github, g   open github
   release, r  add release tag and push tags
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --fish         generate fish completion
   --help, -h     show help
   --version, -v  print the version
```
