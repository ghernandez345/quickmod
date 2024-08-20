package main

import (
	"flag"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/iancoleman/strcase"
)

type Module struct {
	Name     string
	Location string
	ToKebab  func(string) string
}

func createDirectory(name string, location string) {
	err := os.MkdirAll(path.Join(location, name), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func createFilePath(name string, location string, extension string) string {
	if extension == ".scss" {
		return path.Join(location, name, "_styles.scss")
	} else {
		return path.Join(location, name, name+extension)
	}
}

func createFile(extension string, tmplName string, data Module) {
	tmpl, err := template.ParseFiles("templates/" + tmplName)
	if err != nil {
		log.Print(err)
		return
	}

	fileName := createFilePath(data.Name, data.Location, extension)
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("create file:", err)
		return
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

	file.Close()
}

func main() {

	nameFlag := flag.String("n", "", "name of the new module")
	locationFlag := flag.String("l", ".", "location of the new module")
	flag.Parse()

	if *nameFlag == "" {
		log.Fatal("name is required")
	}

	if *locationFlag == "" {
		log.Fatal("location is required")
	}

	createDirectory(*nameFlag, *locationFlag)

	module := Module{*nameFlag, *locationFlag, strcase.ToKebab}

	createFile(".ts", "index_export.tmpl", module)
	createFile(".tsx", "module_source.tmpl", module)
	createFile(".scss", "styles.tmpl", module)
}
