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

	// path := "/Users/jchambrin/Work/dev/"
	path := getPath()
	switch args[0] {
	case "clean-config":
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
	case "configure":
		err := config.Configure(path, args[1:])
		if err != nil {
			log.Fatal(err)
		}
	default:
		exec.Exec(path, os.Args[1:])
	}

}

func getPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path
}
