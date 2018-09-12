package main

import (
    "text/template"
    "os"
    "fmt"
    "github.com/iancoleman/strcase"
)

func main() {
    args := os.Args
    if len(args) < 3 {
        fmt.Println("Error! Example usage: generate Config database")
        return
    }

    typeForGenerate := args[1]
    nameForGenerate := args[2]

    switch typeForGenerate {
    case "Config":
        generateConfig(nameForGenerate)
        break
    default:
        fmt.Printf("You can not generate: %s", nameForGenerate)
    }
}

func generateModel() {

}

var configTemplate = getTemplate(`
package configs

import (
    "sync"
)

type {{ .Name | ToCamel }}Config struct {
    Field   string
}

var {{ .Name | ToLowerCamel }}ConfigOnce sync.Once
var {{ .Name | ToLowerCamel }}Config *{{ .Name | ToCamel }}Config

func Get{{ .Name | ToCamel }}Config() *{{ .Name | ToCamel }}Config {
    {{ .Name | ToLowerCamel }}ConfigOnce.Do(func() {
        {{ .Name | ToLowerCamel }}Config = &{{ .Name | ToCamel }}Config{
            Field:   "",
        }
    })

    return {{ .Name | ToLowerCamel }}Config
}
`)

func generateConfig(name string) {
    file := createFile("configs", fmt.Sprintf("%s%s", name, "Config"))
    defer file.Close()

    err := configTemplate.Execute(file, struct {
        Name string
    }{
        Name: name,
    })

    if err != nil {
        fmt.Println(err)
    } else {
        successMessage(file.Name())
    }
}

func generateController() {

}

func generateMiddleware() {

}

func generateRepository() {

}

func generateService() {

}

func successMessage(name string) {
    fmt.Printf("Fill the file: %s", name)
}

func getTemplate(s string) *template.Template {
    return template.Must(template.New("").Funcs(template.FuncMap{
        "ToCamel":      strcase.ToCamel,
        "ToLowerCamel": strcase.ToLowerCamel,
    }).Parse(s))
}

func createFile(dir string, name string) *os.File {
    file, err := os.Create(fmt.Sprintf("%s/%s.go", dir, name))
    if err != nil {
        fmt.Println(err)
        return nil
    }

    return file
}
