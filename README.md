# Basic API in GoLang

## Table of contents

- [ Introduction ](#intro)
- [ Frameworks ](#frameworks)
- [ How it works ](#about)
- [ Updates ](#updates)
- [ TODO ](#todo)

<a name="intro"></a>
## Introduction

A basic Tasks Api for adding the tasks, with name, description and a boolean flag, whether the task has been completed. 

<a name="frameworks"></a>
## Frameworks

- GoLang
- Gorilla/Mux

<a name="about"></a>
## How it Works

A general premise of this application is creation of a "Task" struct, which will be saved into the database ([ TODO ](#todo)) and available to be accessed, updated and viewed at any moment. Option to create additional tasks will be available as well. 

By creating and initialising a new router using Mux framework 
```go
router := mux.NewRouter()
``` 
routes can be established, which will be executing given functions that hold the functionality of the API. 

To run the application simply run:

```
go run TasksApi.go
```

The program will be listen at port 8000. 

Available http methods: 

```
GET - All tasks or a single task
POST - Create a new task
PUT - Update currently existing entry
DELETE - Delete currently existing entry
```

<a name="updates"></a>
## Updates

* 30/07/2020
    * Initial commit

<a name="todo"><a>
## TODO

[] Add a database