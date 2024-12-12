/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTSTYLE_CLASS.html */
package notesrichtextstyle_test

import (
	"domigo/domino/notesrichtextstyle"
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/require"
)

var richtextstyle notesrichtextstyle.NotesRichTextStyle

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	richtextstyle, _ = session.CreateRichTextStyle()

	defer richtextstyle.Release()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOLD_PROPERTY.html */
func TestBold(t *testing.T) {
	_, err := richtextstyle.Bold()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOLD_PROPERTY.html */
func TestSetBold(t *testing.T) {
	s, err := richtextstyle.Bold()
	require.Nil(t, err)

	err = richtextstyle.SetBold(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTS_PROPERTY.html */
func TestEffects(t *testing.T) {
	_, err := richtextstyle.Effects()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTS_PROPERTY.html */
func TestSetEffects(t *testing.T) {
	s, err := richtextstyle.Effects()
	require.Nil(t, err)

	err = richtextstyle.SetEffects(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSIZE_PROPERTY.html */
func TestFontSize(t *testing.T) {
	_, err := richtextstyle.FontSize()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSIZE_PROPERTY.html */
func TestSetFontSize(t *testing.T) {
	s, err := richtextstyle.FontSize()
	require.Nil(t, err)

	err = richtextstyle.SetFontSize(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULT_PROPERTY_RTSTYLE.html */
func TestIsDefault(t *testing.T) {
	_, err := richtextstyle.IsDefault()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITALIC_PROPERTY.html */
func TestItalic(t *testing.T) {
	_, err := richtextstyle.Italic()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITALIC_PROPERTY.html */
func TestSetItalic(t *testing.T) {
	s, err := richtextstyle.Italic()
	require.Nil(t, err)

	err = richtextstyle.SetItalic(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY.html */
func TestNotesColor(t *testing.T) {
	_, err := richtextstyle.NotesColor()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY.html */
func TestSetNotesColor(t *testing.T) {
	s, err := richtextstyle.NotesColor()
	require.Nil(t, err)

	err = richtextstyle.SetNotesColor(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFONT_PROPERTY.html */
func TestNotesFont(t *testing.T) {
	_, err := richtextstyle.NotesFont()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFONT_PROPERTY.html */
func TestSetNotesFont(t *testing.T) {
	s, err := richtextstyle.NotesFont()
	require.Nil(t, err)

	err = richtextstyle.SetNotesFont(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PASSTHRUHTML_PROPERTY_514_ABOUT.html */
func TestPassThruHTML(t *testing.T) {
	_, err := richtextstyle.PassThruHTML()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PASSTHRUHTML_PROPERTY_514_ABOUT.html */
func TestSetPassThruHTML(t *testing.T) {
	s, err := richtextstyle.PassThruHTML()
	require.Nil(t, err)

	err = richtextstyle.SetPassThruHTML(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STRIKETHROUGH_PROPERTY.html */
func TestStrikeThrough(t *testing.T) {
	_, err := richtextstyle.StrikeThrough()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STRIKETHROUGH_PROPERTY.html */
func TestSetStrikeThrough(t *testing.T) {
	s, err := richtextstyle.StrikeThrough()
	require.Nil(t, err)

	err = richtextstyle.SetStrikeThrough(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDERLINE_PROPERTY.html */
func TestUnderline(t *testing.T) {
	_, err := richtextstyle.Underline()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDERLINE_PROPERTY.html */
func TestSetUnderline(t *testing.T) {
	s, err := richtextstyle.Underline()
	require.Nil(t, err)

	err = richtextstyle.SetUnderline(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
