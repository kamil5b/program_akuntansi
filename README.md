# Accountancy Service

A REST API Accountancy Service in pure go

## System Spesification
- Programming Language Version : Go 1.21.0
- Essential packages:
  * [Go Fiber v2](https://gofiber.io) : a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go
  * [GORM](https://gorm.io/docs/) : an ORM for Go for database management
- Local package management:
  * utilities : tools to make things easier
  * models : program models, in this case only user model
  * database : connecting the program with database server
  * repositories : communicating with database, responsible for database CURD
  * controllers : responsible to processing data. Bridge between controller and repositories
  * requests : the API Gateway, the bridge between input, controller, and output
  * routes : API routes to the controllers

## How to use
1. Download and install go 1.21.1
2. Clone this repository
3. Tidy up importing, download, and installing using: 
  ```console
  go mod tidy
  ```
4. the .envexample turns to .env and fill in / change necessary keys
5. To test it properly you can use:
for testing all packages
  ```console
  go test -p 1 --v ./...
  ```
  ```console
  go test -p 1 --v ./package_name
  ```

6. To run the program : 
  ```console
  go run driver.go
  ```
