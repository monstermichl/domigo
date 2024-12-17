/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOROBJECT_CLASS.html */
package notescolorobject_test

import (
	"testing"

	"github.com/monstermichl/domigo/domino/notescolorobject"
	"github.com/monstermichl/domigo/domino/notessession"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var colorobject notescolorobject.NotesColorObject

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	db, _ := testhelpers.CreateTestDatabase(session)
	colorobject, _ = session.CreateColorObject()

	defer db.Release()
	defer db.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BLUE_PROPERTY_COLOR.html */
func TestBlue(t *testing.T) {
	_, err := colorobject.Blue()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GREEN_PROPERTY_COLOR.html */
func TestGreen(t *testing.T) {
	_, err := colorobject.Green()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HUE_PROPERTY_COLOR.html */
func TestHue(t *testing.T) {
	_, err := colorobject.Hue()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LUMINANCE_PROPERTY_COLOR.html */
func TestLuminance(t *testing.T) {
	_, err := colorobject.Luminance()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY_COLOR.html */
func TestNotesColor(t *testing.T) {
	_, err := colorobject.NotesColor()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY_COLOR.html */
func TestSetNotesColor(t *testing.T) {
	s, err := colorobject.NotesColor()
	require.Nil(t, err)

	err = colorobject.SetNotesColor(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RED_PROPERTY_COLOR.html */
func TestRed(t *testing.T) {
	_, err := colorobject.Red()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SATURATION_PROPERTY_COLOR.html */
func TestSaturation(t *testing.T) {
	_, err := colorobject.Saturation()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHSL_METHOD_COLOR.html */
func TestSetHSL(t *testing.T) {
	_, err := colorobject.SetHSL(0, 0, 0)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETRGB_METHOD_COLOR.html */
func TestSetRGB(t *testing.T) {
	_, err := colorobject.SetRGB(0, 0, 0)
	require.Nil(t, err)
}
