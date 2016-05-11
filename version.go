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

//SetVersionFromFile  number of the tool
func (v *Version) SetVersionFromFile() string {
	// loading the new version
	v.loadVersion()
	// join all the fields into one bulk of data
	s := []string{
		v.version,
		v.majorRevision,
		v.minorRevision,
		v.fixRevisionDet}

	return strings.Join(s, ".")
}

// Full loaded from struct not from file.
func (v Version) Full() string {
	s := []string{
		v.version,
		v.majorRevision,
		v.minorRevision}
	if len(v.fixRevisionDet) > 1 {
		s = append(s, v.fixRevisionDet)
	}
	n := strings.Join(s, ".")

	return n
}
