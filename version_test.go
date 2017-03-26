package Skapt

import . "gopkg.in/check.v1"

func (s *SkaptSuite) TestVersionFromFile(c *C) {
	v := Version{}
	v.SetVersionFromFile()
	c.Assert(v.version, Equals, "0")
	c.Assert(v.majorRevision, Equals, "0")
	c.Assert(v.minorRevision, Equals, "1")
	c.Assert(v.fixRevisionDet, Equals, "1rld")
	c.Assert(v.String(), Equals, "0.0.1.1rld")
}
func (s *SkaptSuite) TestVersionFromApp(c *C) {
	app := NewApp()
	app.SetVersion(false, "0.0.1.1rld")
	c.Assert(app.version.version, Equals, "0")
	c.Assert(app.version.majorRevision, Equals, "0")
	c.Assert(app.version.minorRevision, Equals, "1")
	c.Assert(app.version.fixRevisionDet, Equals, "1rld")
	c.Assert(app.version.String(), Equals, "0.0.1.1rld")

}
