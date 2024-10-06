package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		Help()
		return
	}

	command := os.Args[1]

	switch command {
	case "common":
		Common()
	case "help":
		Help()
	case "go-mvc":
		if len(os.Args) < 3 {
			fmt.Println("Error! Set app name\nmyapp go-mvc <имя_приложения>")
			return
		}
		Go_mvc()
	default:
		Help()
	}
}

// Cоздание конфиг. файлов, общих для всех проектов и языков
func Common() {
	CreateFile("./config/backend/config.json")
	CreateFile("makefile")
	CreateFile(".env")
	CreateFile("TODO")
	CreateFile("readme.md")
}

// Функция Help для показа команд помощи
func Help() {
	fmt.Println("Use: myapp <команда>")
	fmt.Println("Create common files: myapp common")
	fmt.Println("Create GoLang MVC app: myapp go-mvc")
}

// Функция для создания приложения MVC на GoLang
func Go_mvc() {
	app_name := "./" + os.Args[2] + "/"
	CreateFile(app_name + "view.go")
	WriteToFile(app_name+"view.go", fmt.Sprintf("package %s", os.Args[2]))
	CreateFile(app_name + "model.go")
	WriteToFile(app_name+"model.go", fmt.Sprintf("package %s", os.Args[2]))
	CreateFile(app_name + "service.go")
	WriteToFile(app_name+"service.go", fmt.Sprintf("package %s", os.Args[2]))
}

// функция записи в файл чего-либо, например "package main"
func WriteToFile(path string, content string) {
	file, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
	}
}

// функция создания файлов для конфигурирования проекта
//
// path(string): ./config/backend/config.json, или config.json, если без вложенных папок
func CreateFile(path string) {
	if path[0] == '.' {
		str := path[2:]
		parsedStr := strings.Split(str, "/")
		dirs := parsedStr[:len(parsedStr)-1]
		var direction string
		for _, dir := range dirs {
			direction += "./" + string(dir)
			err := os.Mkdir(direction, 0755)
			if err != nil {
				fmt.Println(err)
			}
		}
		file, err := os.OpenFile(path, os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	} else {
		file, err := os.OpenFile(path, os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	}
}
