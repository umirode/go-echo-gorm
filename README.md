Golang REST framework
=====================

## Directory structure

      config/             contains application configurations
      controllers/        contains Web controllers
      database/           contains database package and migrations
      errors/             contains errors and HTTP errors handler
      generator/          contains boilerplate code generator
      ignore/             contains files generated during usage application
      middleware/         contains middlewares
      models/             contains models 
      repositories/       contains repositories 
      response/           contains Web controllers response template 
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

## Generator

For **Windows** use `gen.bat`

For **Linux** use `gen.sh`

---

Model generation: `gen.sh -t model -name Order`

Config generation: `gen.sh -t config -name Cart`

Repository generation: `gen.sh -t repository -name Product`

Service generation: `gen.sh -t service -name Product`

Middleware generation: `gen.sh -t middleware -name Product`