package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	existenceJson()

	if len(os.Args) > 1 && os.Args[1] == "help" {
		printHelp()
		return
	}

	fmt.Println("Invalid option. Use \"./make.exe help\" for usage instructions.")
}

func printHelp() {
	fmt.Println("Usage about \"make\":")
	fmt.Println("this program runs command : \"make [option1] [option2]\".")
	fmt.Println("---------------------------------------------------")
	fmt.Println("make help")
	fmt.Println("   - show this command help")
	fmt.Println()
}

func existenceJson() {
	type Setting struct {
		FolderPath string `json:"folder_path" default:"./"`
		PyVenv     string `json:"py_venv" default:"myenv"`
	}

	if _, err := os.Stat("setting.json"); os.IsNotExist(err) {
		file, err := os.OpenFile("setting.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}

		setting := Setting{
			FolderPath: "./",
			PyVenv:     "myenv",
		}
		if err := writeJSONToFile("setting.json", setting); err != nil {
			log.Fatalf("Error writing JSON to file: %v", err)
		}

		defer file.Close()
	} else if err != nil {
		log.Fatal(err)
	}
}

func writeJSONToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

/*
追加したいコマンド
pyinit
pyadd
pyremove
pyoutput

gitinit
gitupdate
gitupdate-ac
*/
