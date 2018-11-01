Golang REST framework
=====================

## Directory structure

      cmd/                contains cli commands
      configs/            contains configurations
      controllers/        contains Web controllers
      database/           contains database package and migrations
      errors/             contains errors and HTTP errors handler
      generator/          contains boilerplate code generator
      ignore/             contains files generated during usage application
      middleware/         contains middlewares
      models/             contains models 
      repositories/       contains repositories 
      services/           contains services with main app logic
      router/             contains router and routes
      vendor/             contains dependent 3rd-party packages


## Install
Copy `.env.dist` to `.env`

Configure `.env`

Install dependencies using dep: `dep ensure`

Start app: `go run main.go`

You can then access the application through the following URL: http://127.0.0.1:8080

## Install with Docker
Copy `.env.dist` to `.env`

Start the container: `docker-compose up -d`

You can then access the application through the following URL: http://127.0.0.1:8080

## Cli

For **Windows** use `cmd.bat`

For **Linux** use `cmd.sh`

## Cli Boilerplate code generator

Templates file contains in `templates` folders.

___

Model generation: `cmd.bat generator model Order`

Config generation: `cmd.bat generator config Cart`

Repository generation: `cmd.bat generator repository Product`

Service generation: `cmd.bat generator service Product`

Middleware generation: `cmd.bat generator middleware Product`

## Cli Auth

Create user: `cmd.bat auth create test@mail.com StrongPassword`

Delete user: `cmd.bat auth delete test@mail.com`

Delete user refresh tokens: `cmd.bat auth delete-tokens test@mail.com`