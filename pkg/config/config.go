package config

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	configDir  = os.Getenv("HOME") + "/.config/gitall"
	configFile = filepath.Join(configDir, "config.yml")
)

type configData struct {
	Directories []string
}

func List() ([]string, error) {
	res, err := readConfig()
	if err != nil {
		return nil, err
	}
	return res.Directories, nil
}

func Configure(path string, folders []string) error {
	directories := GetDirectories(path, folders)
	return persistConfig(directories)
}

func CleanConfiguration() error {
	return os.Remove(configFile)
}

func GetDirectories(path string, folders []string) []string {
	includeAll := len(folders) == 0
	directories := make([]string, 0)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() && (includeAll || contains(folders, file.Name())) {
			if _, err := os.Stat(filepath.Join(path, file.Name(), ".git")); err == nil {
				directories = append(directories, filepath.Join(path, file.Name()))
			}
		}
	}

	return directories
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func persistConfig(directories []string) error {
	err := createConfigDir()
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(&configData{
		Directories: directories,
	})
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(configFile, data, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func readConfig() (*configData, error) {
	yfile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return &configData{Directories: make([]string, 0)}, nil
	}
	var data *configData
	err = yaml.Unmarshal(yfile, &data)

	return data, err
}

func createConfigDir() error {
	err := os.MkdirAll(configDir, os.ModePerm)
	return err
}
