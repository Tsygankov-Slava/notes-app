# Notes App
This application implements REST API for working with notes.


## Table of contents
- [Technical requirements](#technical-requirements)
- [Features](#features)
- [Project architecture](#project-architecture)
  - [Project structure](#project-structure)
  - [Tables for links](#tables-for-links)
- [Run and Testing my REST API](#run-and-testing-my-rest-api)
- [License](#license)


## Technical requirements
#### Programming Language: `Go 1.22.2`
#### Database: `PostgreSQL`
#### Main libs:
  - `github.com/swaggo/gin-swagger` - This library integrates Swagger with the Gin framework, which allows you to automatically generate API documentation based on annotations in the code.
  - `github.com/dgrijalva/jwt-go` - This library provides tools for creating and validating JWTs.
  - `github.com/gin-gonic/gin` - It is a web framework for the Go language that facilitates the creation of web applications. It provides many useful tools such as routing, request processing and more.
  - `github.com/jmoiron/sqlx` - This is an extension of the standard database package in Go for working with SQL. It provides additional functions and a user-friendly interface for executing SQL queries and working with the resulting data.
  - `github.com/lib/pq` - This library is the PostgreSQL driver for Go. It provides an interface for connecting to a PostgreSQL database and making queries to it.
  - `github.com/spf13/viper"` - This is a library for working with configuration files in Go. It allows you to conveniently download and use configuration parameters from various sources, such as files, environment variables, command line flags, and others.
  - `github.com/sirupsen/logrus` - This is a logger for Go that provides rich logging functionality such as various logging levels, log formatting and more.
  - `crypto/sha1` - This package is included in the standard Go library and provides an implementation of the SHA-1 hashing algorithm.
  - `net/http` - This package is included in the standard Go library and provides functionality for working with HTTP requests and HTTP servers.


## Features
#### Implemented functionality:
- User authorization
- User registration
  > - Authorization and registration is performed by sending a username and password to the application.
  > - In response, the user receives a JWT-token.
- Posting note
  > - The note is posting by sending data in JSON format: title, text.
  > - Notes are posted only by authorized users.
  > - If the response is successful, the data of the created note is returned.
- Editing note
  > - The functionality is the same as when posting a note.
  > - Only the creator of the note can edit the note if the time limit for posting the note is less than one day.
- Displaying a list of notes
  > - Notes are displayed by the date they were added.
  > - For each note, its data is returned.

[🔝Table of contents](#table-of-contents)


## Project architecture
### Project structure
![](docs_app/images/project-structure.png)
### Tables for links
![](docs_app/images/tables-for-links.png)

[🔝Table of contents](#table-of-contents)


##  Run and Testing my REST API
Setting up the database:
```
/* optional */
$ docker pull postgres

$ migrate -path ./schema -database 'postgres://postgres:root@localhost:5436/notes-db?sslmode=disable' up
$ docker run --name=notes-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root -e POSTGRES_DB=notes-db -p 5436:5432 -d --rm postgres
```

Setting up the swagger documentation:
```
$ swag init -g cmd/main.go
```

Run app:
```
$ go run cmd/main.go
```

Open swagger documentation in browser: `http://localhost:8001/swagger/index.html`

[🔝Table of contents](#table-of-contents)


## License
`notes-app` is distributed under the [MIT License](https://github.com/Tsygankov-Slava/notes-app/blob/main/LICENSE), on behalf of Tsygankov Viacheslav.

[🔝Table of contents](#table-of-contents)