package Skapt

import (
	"fmt"
	"os"
	"strings"
)

/**
 * Private vars
 */
var (
	filePath    = getPathVersion()
	contentFile string
)

// Version that stores all the
// basic information
type Version struct {
	// tool version
	version string
	// tool major revision
	majorRevision string
	// app minor revision
	minorRevision string
	// app fix revision detail
	fixRevisionDet string
}

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
		name}

	path := strings.Join(holder, "/")

	return path
}

func openFileVersion(path string) *os.File {

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		fmt.Println("Can't open file VERSION")
	}

	return file

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

	_, err := file.Read(cnt)

	if err != nil {
		fmt.Println("Can't read VERSION file")
	}

	return string(cnt)
}

func (v *Version) loadVersion() {
	contentFile = getContentVersion()
	s := strings.Split(contentFile, ".")

	v.version = s[0]
	v.majorRevision = s[1]
	v.minorRevision = s[2]
	v.fixRevisionDet = s[3]

}

// SetVersionFromFile public metod to export the version
// number of the tool
func (v *Version) SetVersionFromFile() string {
	// loading the new version
	v.loadVersion()
	// join all the fields into one bulk of data
	s := []string{
		v.version,
		v.majorRevision,
		v.minorRevision,
		v.fixRevisionDet}
	n := strings.Join(s, ".")

	return n
}

// GetVersion does the same thing above but return just the
// version loaded from keyboard not from file.
func (v Version) GetVersion() string {
	s := []string{
		v.version,
		v.majorRevision,
		v.minorRevision,
		v.fixRevisionDet}
	n := strings.Join(s, ".")

	return n
}
