package Skapt

import "strings"

// Full loaded from keyboard not from file.
func (v Version) Full() string {
	s := []string{
		v.version,
		v.majorRevision,
		v.minorRevision,
		v.fixRevisionDet}
	n := strings.Join(s, ".")

	return n
}
