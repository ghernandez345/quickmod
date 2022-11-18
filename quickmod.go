package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

type Module struct {
	Name     string
	Location string
}

func convertToSnakecase(str string) string {
	return str
}

func createFile(extension string, tmplName string, data Module) {
	tmpl, err := template.ParseFiles("templates/" + tmplName)
	if err != nil {
		log.Print(err)
		return
	}

	file, err := os.Create(data.Location + ".ts")
	if err != nil {
		log.Println("create file:", err)
		return
	}

	tmpl.Execute(file, data)
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

	module := Module{*nameFlag, *locationFlag}

	createFile(".ts", "index_export.tmpl", module)

	// convertToSnakecase(*nameFlag)
}
