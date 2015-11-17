package Skapt

// GetVersion does the same thing above but return just the
import "strings"

// version loaded from keyboard not from file.
func (v Version) Version() string {
	s := []string{
		v.version,
		v.majorRevision,
		v.minorRevision,
		v.fixRevisionDet}
	n := strings.Join(s, ".")

	return n
}
