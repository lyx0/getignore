package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var pLang string
	// Name and url
	language := map[string]string{
		"c":    "https://raw.githubusercontent.com/github/gitignore/main/C.gitignore",
		"go":   "https://raw.githubusercontent.com/github/gitignore/main/Go.gitignore",
		"rust": "https://raw.githubusercontent.com/github/gitignore/main/Rust.gitignore",
	}

	flag.StringVar(&pLang, "lang", "go", "Programming language")
	flag.Parse()
	// fmt.Println(pLang)

	fmt.Println("INPUT:", pLang)
	for name, url := range language {
		if string(name) == pLang {
			if err := getFile(url); err != nil {
				fmt.Println(err)
				fmt.Println("FOUND", name, url)
			}

		}
	}

	// fmt.Println(pLang)
}

func getFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		// log.Error(err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Error(err)
		return err
	}

	fmt.Println("---FILE CONTENTS----")
	fmt.Println(string(body))

	writeToFile(body)
	return nil
}

func writeToFile(contents []byte) error {
	err := os.WriteFile(".gitignore", contents, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
