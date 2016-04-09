package Skapt

import (
	"fmt"
	"os"
	"strings"
)

func getPathVersion() string {
	// our file name
	name := "VERSION"

	rootPath, err := os.Getwd()

	if err != nil {
		fmt.Println("Can't optain the root path of the project")
	}

	//host our root/base path and our fileName(VERSION)
	holder := []string{
		rootPath,
		name,
	}

	path := strings.Join(holder, "/")

	return path
}

func getContentVersion() string {
	// local var that stores the content
	cnt := make([]byte, 10)
	//var path string

	file := openFileVersion(filePath)

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	n, err := file.Read(cnt)

	if err != nil && n == 0 {
		fmt.Println("Can't read VERSION file")
	}

	return string(cnt)
}

func openFileVersion(path string) *os.File {

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		fmt.Println("Can't open file VERSION")
	}

	return file

}
