/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"jchambrin.fr/gitall/pkg/config"
	"jchambrin.fr/gitall/pkg/exec"
)

func Execute() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("no argument provided")
	}

	path := getPath()
	switch args[0] {
	case "configure":
		executeConfigure(path, args[1:])
	default:
		exec.Exec(path, os.Args[1:])
	}
}

func executeConfigure(path string, args []string) {
	if len(args) == 0 {
		log.Fatal("no argument provided")
	}
	switch args[0] {
	case "clean":
		err := config.CleanConfiguration()
		if err != nil {
			log.Fatal(err)
		}
	case "list":
		list, err := config.List()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(list)
	case "create":
		err := config.Configure(path, args[1:])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path
}
