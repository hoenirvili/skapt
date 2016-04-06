package Skapt

import "strings"

// Full loaded from keyboard not from file.
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
