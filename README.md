# Agilka - console TaskManager

Agilka must use with GIT repository and without Database and Internet connections. 

Agilka is _Taskmanager like a code_.

Powered with Golang for Faster.

## Install

1. Install Golang
2. `$ go get github.com/realb0t/agilka`
3. `$ cd $GOPATH/src/github.com/realb0t/agilka`
4. `$ go install`
5. `$ agilka help`

## Usage

Agilka help:
```
$ agilka help
```

Initialize project:
```
$ agilka init
```

Work with tasks and tickets:
```
$ agilka task help
```

Create task:
```
$ agilka task create code=gooVal title='Hello world' --json='{"code": "gooBar"}'
```

Edit task:
```
$ agilka task edit <taskCode> fieldName=value
```

Task to TODO:
```
$ agilka task plan <taskCode>
```

Display task list
```
$ agilka task list [state1] ... [stateN]
```

## Project structure

```
/tasks/*.json
/attaches/**.*
/Agilkafile
/.git
/.gitignore
```

## Current Task format

```json
{
  "code": "ticketCodeName",
  "title": "Заголовок таска",
  "desc": "# Описание таска\nС поддержкой **mdoc**",
  "author": "gitUserName",
  "state": "backlog,todo,process,...,completed,...",
}
```

## Planned

* GIT integration
* More task functionality for Agile
* Attaches support
* Web Interface