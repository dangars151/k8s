package main

import (
	"fmt"
	"log"
	"os"
)

var (
	Info  = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn  = log.New(os.Stdout, "[WARN]: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
)

const path = "var/log/log.txt"

func InitLogger() {
	createFile()

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}

	Info.SetOutput(file)
	Warn.SetOutput(file)
	Error.SetOutput(file)
}

func createFile() {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}

	fmt.Println("Create file success", path)
}
