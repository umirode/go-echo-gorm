package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"os"
	"text/template"
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
	case "Model":
		generateModel(nameForGenerate)
		break
	case "Repository":
		generateRepository(nameForGenerate)
		break
	default:
		fmt.Printf("You can not generate: %s", nameForGenerate)
	}
}

func generateModel(name string) {
	file := createFile("models", fmt.Sprintf("%s%s", name, "Model"))
	defer file.Close()

	modelTemplate := getTemplate("models")
	err := modelTemplate.Execute(file, struct {
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

func generateConfig(name string) {
	file := createFile("configs", fmt.Sprintf("%s%s", name, "Config"))
	defer file.Close()

	configTemplate := getTemplate("configs")
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

func generateRepository(name string) {
	file := createFile("repositories", fmt.Sprintf("%s%s", name, "Repository"))
	defer file.Close()

	configTemplate := getTemplate("repositories")
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

func generateService() {

}

func successMessage(name string) {
	fmt.Printf("Fill the file: %s", name)
}

func getTemplate(dir string) *template.Template {
	templateBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/template.txt", dir))
	if err != nil {
		fmt.Printf("Template read error: %s", err)
	}

	return template.Must(template.New("").Funcs(template.FuncMap{
		"ToCamel":      strcase.ToCamel,
		"ToLowerCamel": strcase.ToLowerCamel,
	}).Parse(string(templateBytes)))
}

func createFile(dir string, name string) *os.File {
	file, err := os.Create(fmt.Sprintf("%s/%s.go", dir, name))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return file
}
