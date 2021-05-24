# UserManagement
## server
    API server written in Golang for user management

## Installation
------------
A little intro about the installation. 

$ git clone https://github.com/onkarsutar/UserManagement.git
$ cd ../UserManagement/server/server.go
$ go run server.go

## REQUIREMENTS
------------

This module requires the following modules:

MongoDB
Golang

## CONFIGURATION
------------
Add server port, MongoDB DSN in config.json file located at ./configs/config.json

## Using Docker
Download code from Github

Run below command which starts both mongo and go application in seperate containers.

docker-compose up --build

## API's
This application contains below API's.

1. http://localhost:8000 
    Entry point of application
    Method Type :  GET
    Request Data : None
2. http://localhost:8000/checkstatus
    API to check status of application
    Method Type :  GET
    Request Data : None
3. http://localhost:8000/o/serverstats
    API to check API statistics of application
    Method Type :  GET
    Request Data : None
4. http://localhost:8000/o/user/add
    API to add user in Database
    Method Type :  POST
    Request Data : 
    {
        "loginID": "user1",
        "password": "oldP@ss0rd",
        "userName": "User One",
        "emailID": "user1@gmail.com"
    }
5. http://localhost:8000/o/user/changepassword
    API to update password
    Method Type :  POST
    Request Data : 
    {
        "loginID": "user1",
        "password": "oldP@ss0rd",
        "newPassword": "newP@ss0rd"
    }