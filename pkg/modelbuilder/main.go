package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Parse command-line flags
	u := flag.String("u", "user", "user")
	p := flag.String("p", "", "password")
	d := flag.String("d", "csw", "database")
	t := flag.String("t", "yourtables", "tables comma separated")
	db := flag.String("db", "localhost", "database ip")

	flag.Parse()

	// Execute the gentool command to generate models
	cmd := exec.Command("gentool", "-dsn", *u+`:`+*p+`@tcp(`+*db+`:5432)/`+*d+`?charset=utf8mb4`, "-outPath=library/struct/model/", "-modelPkgName=model", "-onlyModel", "-fieldNullable", "-tables="+*t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// Get all filenames in a folder
	dir := "./library/struct/model/"
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	list, err := file.Readdir(-1)
	if err != nil {
		fmt.Println(err)

	}

	// Clean up generated files
	for _, f := range list {
		file, err := os.ReadFile(dir + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		fileString := string(file)
		var clean string
		splits := strings.Split(fileString, "package model")
		if len(splits) > 1 {
			clean = "package model" + splits[1]
		}

		clean = strings.ReplaceAll(clean, "int32", "int")

		err = os.WriteFile(dir+f.Name(), []byte(clean), 0)
		if err != nil {
			fmt.Println(err)
		}

		// Rename files if necessary
		name := strings.TrimSuffix(f.Name(), "gen.go")
		if name == f.Name() { // no suffix so no need to rename
			continue
		}
		err = os.Rename(dir+f.Name(), dir+name+"go")
		if err != nil {
			fmt.Println(err)
		}
	}

}
