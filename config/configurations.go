package config

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

// VERSION - The Current Semantic Version of API
const VERSION = "0.1.1"

var (
	User        = "elastic"   //default
	Pass        = "elastic"   //default
	Host        = "127.0.0.1" //default
	Port        = "7775"      //default
	Manager     = "elastic"   //default
	ManagerPort = "9200"      //default
	Routes      = flag.Bool("routes", false, "Generate router documentation")
)

func getLinesInFile(fileName *os.File) []string {

	scanner := bufio.NewScanner(fileName)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	fileName.Close()
	return result
}

func Settings() {
	fileName := "config/apinosql.conf"
	fileOpen, err := os.Open(fileName)

	if err != nil {
		log.Printf(err.Error() + " Using default configuration")
	}

	for index, line := range getLinesInFile(fileOpen) {
		_ = index
		vConfig := strings.Split(line, "=")
		vKey, vValue := vConfig[0], vConfig[1]
		if vKey == "user" {
			User = vValue
		}
		if vKey == "password" {
			Pass = vValue
		}
		if vKey == "host" {
			Host = vValue
		}
		if vKey == "port" {
			Port = vValue
		}
		if vKey == "manager" {
			Manager = vValue
		}

		if vKey == "manager_port" {
			ManagerPort = vValue
		}

	}
}
