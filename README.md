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

## Cli

For **Windows** use `cmd.bat`

For **Linux** use `cmd.sh`