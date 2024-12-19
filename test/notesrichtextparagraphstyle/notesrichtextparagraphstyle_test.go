/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTPARAGRAPHSTYLE_CLASS_3756.html */
package notesrichtextparagraphstyle_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	"github.com/stretchr/testify/require"
)

var richtextparagraphstyle domigo.NotesRichTextParagraphStyle

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := domigo.Initialize()
	richtextparagraphstyle, _ = session.CreateRichTextParagraphStyle()

	defer richtextparagraphstyle.Release()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_6587.html */
func TestAlignment(t *testing.T) {
	_, err := richtextparagraphstyle.Alignment()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_6587.html */
func TestSetAlignment(t *testing.T) {
	s, err := richtextparagraphstyle.Alignment()
	require.Nil(t, err)

	err = richtextparagraphstyle.SetAlignment(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIRSTLINELEFTMARGIN_PROPERTY_4412.html */
func TestFirstLineLeftMargin(t *testing.T) {
	_, err := richtextparagraphstyle.FirstLineLeftMargin()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIRSTLINELEFTMARGIN_PROPERTY_4412.html */
func TestSetFirstLineLeftMargin(t *testing.T) {
	s, err := richtextparagraphstyle.FirstLineLeftMargin()
	require.Nil(t, err)

	err = richtextparagraphstyle.SetFirstLineLeftMargin(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERLINESPACING_PROPERTY_9542.html */
func TestInterLineSpacing(t *testing.T) {
	_, err := richtextparagraphstyle.InterLineSpacing()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERLINESPACING_PROPERTY_9542.html */
func TestSetInterLineSpacing(t *testing.T) {
	s, err := richtextparagraphstyle.InterLineSpacing()
	require.Nil(t, err)

	err = richtextparagraphstyle.SetInterLineSpacing(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEFTMARGIN_PROPERTY_4793.html */
func TestLeftMargin(t *testing.T) {
	_, err := richtextparagraphstyle.LeftMargin()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEFTMARGIN_PROPERTY_4793.html */
func TestSetLeftMargin(t *testing.T) {
	s, err := richtextparagraphstyle.LeftMargin()
	require.Nil(t, err)

	err = richtextparagraphstyle.SetLeftMargin(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PAGINATION_PROPERTY_RTPARAGRAPHSTYLE.html */
func TestPagination(t *testing.T) {
	_, err := richtextparagraphstyle.Pagination()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PAGINATION_PROPERTY_RTPARAGRAPHSTYLE.html */
func TestSetPagination(t *testing.T) {
	s, err := richtextparagraphstyle.Pagination()
	require.Nil(t, err)

	err = richtextparagraphstyle.SetPagination(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RIGHTMARGIN_PROPERTY_480.html */
func TestRightMargin(t *testing.T) {
	_, err := richtextparagraphstyle.RightMargin()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RIGHTMARGIN_PROPERTY_480.html */
func TestSetRightMargin(t *testing.T) {
	s, err := richtextparagraphstyle.RightMargin()
	require.Nil(t, err)

	err = richtextparagraphstyle.SetRightMargin(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGABOVE_PROPERTY_7663.html */
func TestSpacingAbove(t *testing.T) {
	_, err := richtextparagraphstyle.SpacingAbove()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGABOVE_PROPERTY_7663.html */
func TestSetSpacingAbove(t *testing.T) {
	s, err := richtextparagraphstyle.SpacingAbove()
	require.Nil(t, err)

	err = richtextparagraphstyle.SetSpacingAbove(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGBELOW_PROPERTY_8858.html */
func TestSpacingBelow(t *testing.T) {
	_, err := richtextparagraphstyle.SpacingBelow()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGBELOW_PROPERTY_8858.html */
func TestSetSpacingBelow(t *testing.T) {
	s, err := richtextparagraphstyle.SpacingBelow()
	require.Nil(t, err)

	err = richtextparagraphstyle.SetSpacingBelow(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TABS_PROPERTY_5218.html */
func TestTabs(t *testing.T) {
	_, err := richtextparagraphstyle.Tabs()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARALLTABS_METHOD_3625.html */
func TestClearAllTabs(t *testing.T) {
	err := richtextparagraphstyle.ClearAllTabs()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETTAB_METHOD_2382.html */
func TestSetTab(t *testing.T) {
	err := richtextparagraphstyle.SetTab(domigo.NOTESRICHTEXTPARAGRAPHSTYLE_RULER_ONE_INCH)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETTABS_METHOD_181.html */
func TestSetTabs(t *testing.T) {
	err := richtextparagraphstyle.SetTabs(1, domigo.NOTESRICHTEXTPARAGRAPHSTYLE_RULER_ONE_INCH)
	require.Nil(t, err)
}
