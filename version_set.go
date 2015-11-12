package Skapt

// SetVersionFromFile public metod to export the version
import "strings"

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