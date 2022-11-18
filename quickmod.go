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

func createFileName(name string, location string, extension string) string {
	if extension == ".scss" {
		return path.Join(location, "_styles.scss")
	} else {
		return path.Join(location, name+extension)
	}
}

func createFile(extension string, tmplName string, data Module) {
	tmpl, err := template.ParseFiles("templates/" + tmplName)
	if err != nil {
		log.Print(err)
		return
	}

	fileName := createFileName(data.Name, data.Location, extension)
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

	module := Module{*nameFlag, *locationFlag, strcase.ToKebab}

	createFile(".ts", "index_export.tmpl", module)
	createFile(".tsx", "module_source.tmpl", module)
	createFile(".scss", "styles.tmpl", module)
}
