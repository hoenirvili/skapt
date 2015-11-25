package Skapt

import "strings"

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

func (v *Version) loadVersion() {
	contentFile = getContentVersion()
	s := strings.Split(contentFile, ".")
	v.version = s[0]
	v.majorRevision = s[1]
	v.minorRevision = s[2]
	v.fixRevisionDet = s[3]

}
