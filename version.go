package Skapt

import "strings"

///////////////////////////////////////////////////////
//
//				GLOBAL PACKAGE
//					VARS
///////////////////////////////////////////////////////

var (
	filePath    = getPathVersion()
	contentFile string
)

///////////////////////////////////////////////////////
//
//				GLOBAL TYPE
//
///////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////
//
//				METHODS
//				GET
///////////////////////////////////////////////////////

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

func (v *Version) loadVersion() {
	contentFile = getContentVersion()
	s := strings.Split(contentFile, ".")

	v.version = s[0]
	v.majorRevision = s[1]
	v.minorRevision = s[2]
	v.fixRevisionDet = s[3]

}

///////////////////////////////////////////////////////
//
//				METHODS
//				SET
///////////////////////////////////////////////////////

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
