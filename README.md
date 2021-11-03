# GoCrud
###### Crud operataions with go.

## Features

- go 1.16
- gin-gonic for routing
- Postgre sql
- gorm for PostgreSQL database access
- crypto/bcrypt for password crypt
- testify, mockery for unit test
- swagger for api documentation
- go modules for package management
- dockerized environment
- structure design inspired by https://github.com/golang-standards/project-layout

## Usage

####  - API

##### #Register (Insert) New User
- PUT http://localhost:8080/v1/api/users
- body:  { "Name": "Arden", "Email": "arden@mail.com", "Password": "strongpassword123" }
- returns auth token of inserted user
```sh
curl --location --request PUT 'http://localhost:8080/v1/api/users' \
--header 'Content-Type: application/json' \
--data-raw '{
	"Name":"Arden",
    "Email":"arden@mail.com",
    "Password":"strongpassword123"
}'
```

##### #Login (Authorize) With Given User Credentials
- POST http://localhost:8080/v1/api/auth/login
- body:  { "Email": "arden@mail.com", "Password": "strongpassword123" }
- returns auth token of loggedin user
```sh
curl --location --request POST 'http://localhost:8080/v1/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
	"Email":"arden@mail.com",
    "Password":"strongpassword123"
}'
```

##### #GetUserById - (Authorization Required)
- GET http://localhost:8080/v1/api/users/1 
- headers : { "Authorization": "AUTHTOKEN"}
```sh
curl --location --request GET 'http://localhost:8080/v1/api/users/1' \
--header 'Authorization: AUTHTOKEN'
```

##### # GetAllUsers - (Authorization Required)
- GET http://localhost:8080/v1/api/users
- headers : { "Authorization": "AUTHTOKEN"}
```sh
curl --location --request GET 'http://localhost:8080/v1/api/users' \
--header 'Authorization: AUTHTOKEN'
```

##### # UpdateUser Credentials With Given Id Of User - (Authorization Required)
- PATCH http://localhost:8080/v1/api/users/1
- headers : { "Authorization": "AUTHTOKEN"}
- body:  { "Name": "ArdenNew", "Email": "arden_new@mail.com", "Password": "strongpassword1234" }
```sh
curl --location --request PATCH 'http://localhost:8080/v1/api/users/1' \
--header 'Authorization: AUTHTOKEN' \
--header 'Content-Type: application/json' \
--data-raw '{
	"Name":"ArdenNew",
	"Email":"arden_new@mail.com",
    "Password":"strongpassword1234"
}'
```

##### # DeleteUser (Set DeletedDate) With Id Of User - (Authorization Required)
- DELETE http://localhost:8080/v1/api/users/1
- headers : { "Authorization": "AUTHTOKEN"}
```sh
curl --location --request DELETE 'http://localhost:8080/v1/api/users/1' \
--header 'Authorization: AUTHTOKEN'
```

>  Note: Response model showen on below
```sh
type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}
```
> Or just run the app, and browse to http://localhost:8080/swagger/index.html, but not forget to authorize on swagger.

## Installation & Run

##### # docker-compose
- Setup docker environemnt use "docker-compose.yml" to docker-compose up.
- Go to project folder directory and execute
```sh
# start docker environment
$ docker-compose up -d (--build)

# list running services
$ docker-compose ps

# stop all containers
$ docker-compose stop

# remove all containers
$ docker-compose rm
```

##### # Run Locally

- Download posgresql from https://www.postgresql.org/download/ and Setup
- Setup postgresql connection variables with local.env file or just setup yours and change parameters inside local.env
```sh
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_DB=DBNAME
POSTGRES_PASSWORD=STRONGPASSWORD
```

- After installation create DB on postgre with your given "DBNAME" then just go to project directory and execute
```sh
$ go run main.go
```

> ✨-postgre sql runs automatically on 5432 port
> ✨ server run on 8080 port