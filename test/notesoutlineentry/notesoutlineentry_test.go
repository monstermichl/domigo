/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESOUTLINEENTRY_CLASS.html */
package notesoutlineentry_test

import (
	"domigo/domino/notesoutlineentry"
	"domigo/domino/notessession"
	testhelpers "domigo/test/helpers"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var outlineentry notesoutlineentry.NotesOutlineEntry

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	var info string

	session, err := notessession.Initialize()
	defer session.Release()

	defer func() {
		fmt.Println(err)
		fmt.Println(info)
	}()

	if err != nil {
		info = "Session could not be initialized"
		return
	}

	db, err := testhelpers.CreateTestDatabase(session)
	defer db.Release()
	defer db.Remove()

	if err != nil {
		info = "Database could not be created"
		return
	}

	outline, err := db.CreateOutline("test-outline")
	defer outline.Release()

	if err != nil {
		info = "Outline could not be created"
		return
	}

	outlineentry, err = outline.CreateEntry("testentry")
	defer outlineentry.Release()

	if err != nil {
		info = "Outline entry could not be created"
		return
	}
	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINEENTRY.html */
func TestAlias(t *testing.T) {
	_, err := outlineentry.Alias()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINEENTRY.html */
func TestSetAlias(t *testing.T) {
	s, err := outlineentry.Alias()
	require.Nil(t, err)

	err = outlineentry.SetAlias(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENT_PROPERTY_OUTLINEENTRY.html */
func TestDocument(t *testing.T) {
	_, err := outlineentry.Document()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRYCLASS_PROPERTY_OUTLINEENTRY.html */
func TestEntryClass(t *testing.T) {
	_, err := outlineentry.EntryClass()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_OUTLINEENTRY.html */
func TestFormula(t *testing.T) {
	_, err := outlineentry.Formula()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FRAMETEXT_PROPERTY_OUTLINEENTRY.html */
func TestFrameText(t *testing.T) {
	_, err := outlineentry.FrameText()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FRAMETEXT_PROPERTY_OUTLINEENTRY.html */
func TestSetFrameText(t *testing.T) {
	s, err := outlineentry.FrameText()
	require.Nil(t, err)

	err = outlineentry.SetFrameText(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASCHILDREN_PROPERTY_OUTLINEENTRY.html */
func TestHasChildren(t *testing.T) {
	_, err := outlineentry.HasChildren()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func TestHideFormula(t *testing.T) {
	_, err := outlineentry.HideFormula()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func TestSetHideFormula(t *testing.T) {
	s, err := outlineentry.HideFormula()
	require.Nil(t, err)

	err = outlineentry.SetHideFormula(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMAGESTEXT_PROPERTY_OUTLINEENTRY.html */
func TestImagesText(t *testing.T) {
	_, err := outlineentry.ImagesText()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMAGESTEXT_PROPERTY_OUTLINEENTRY.html */
func TestSetImagesText(t *testing.T) {
	s, err := outlineentry.ImagesText()
	require.Nil(t, err)

	err = outlineentry.SetImagesText(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY_OUTLINEENTRY.html */
func TestIsHidden(t *testing.T) {
	_, err := outlineentry.IsHidden()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY_OUTLINEENTRY.html */
func TestSetIsHidden(t *testing.T) {
	s, err := outlineentry.IsHidden()
	require.Nil(t, err)

	err = outlineentry.SetIsHidden(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMNOTES_PROPERTY_OUTLINEENTRY.html */
func TestIsHiddenFromNotes(t *testing.T) {
	_, err := outlineentry.IsHiddenFromNotes()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMNOTES_PROPERTY_OUTLINEENTRY.html */
func TestSetIsHiddenFromNotes(t *testing.T) {
	s, err := outlineentry.IsHiddenFromNotes()
	require.Nil(t, err)

	err = outlineentry.SetIsHiddenFromNotes(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMWEB_PROPERTY_OUTLINEENTRY.html */
func TestIsHiddenFromWeb(t *testing.T) {
	_, err := outlineentry.IsHiddenFromWeb()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMWEB_PROPERTY_OUTLINEENTRY.html */
func TestSetIsHiddenFromWeb(t *testing.T) {
	s, err := outlineentry.IsHiddenFromWeb()
	require.Nil(t, err)

	err = outlineentry.SetIsHiddenFromWeb(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINTHISDB_PROPERTY_OUTLINEENTRY.html */
func TestIsInThisDB(t *testing.T) {
	_, err := outlineentry.IsInThisDB()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATE_PROPERTY_OUTLINEENTRY.html */
func TestIsPrivate(t *testing.T) {
	_, err := outlineentry.IsPrivate()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEEPSELECTIONFOCUS_PROPERTY_OUTLINEENTRY.html */
func TestKeepSelectionFocus(t *testing.T) {
	_, err := outlineentry.KeepSelectionFocus()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEEPSELECTIONFOCUS_PROPERTY_OUTLINEENTRY.html */
func TestSetKeepSelectionFocus(t *testing.T) {
	s, err := outlineentry.KeepSelectionFocus()
	require.Nil(t, err)

	err = outlineentry.SetKeepSelectionFocus(s)
	require.Nil(t, err)
}

/* TODO: Access type for Label could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LABEL_PROPERTY_OUTLINEENTRY.html */
func TestLabel(t *testing.T) {
	_, err := outlineentry.Label()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LABEL_PROPERTY_OUTLINEENTRY.html */
func TestSetLabel(t *testing.T) {
	s, err := outlineentry.Label()
	require.Nil(t, err)

	err = outlineentry.SetLabel(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY_OUTLINEENTRY.html */
func TestLevel(t *testing.T) {
	_, err := outlineentry.Level()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEDELEMENT_PROPERTY_OUTLINEENTRY.html */
func TestNamedElement(t *testing.T) {
	_, err := outlineentry.NamedElement()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OUTLINEENTRY.html */
func TestType(t *testing.T) {
	_, err := outlineentry.Type()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_URL_PROPERTY_OUTLINEENTRY.html */
func TestURL(t *testing.T) {
	_, err := outlineentry.URL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEHIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func TestUseHideFormula(t *testing.T) {
	_, err := outlineentry.UseHideFormula()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEHIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func TestSetUseHideFormula(t *testing.T) {
	s, err := outlineentry.UseHideFormula()
	require.Nil(t, err)

	err = outlineentry.SetUseHideFormula(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEW_PROPERTY_OUTLINEENTRY.html */
func TestView(t *testing.T) {
	_, err := outlineentry.View()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETACTION_METHOD_OUTLINEENTRY.html */
func TestSetAction(t *testing.T) {
	_, err := outlineentry.SetAction("@True")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDELEMENT_METHOD_OUTLINEENTRY.html */
// func TestSetNamedElement(t *testing.T) {
// 	_, err := outlineentry.SetNamedElement( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNOTELINK_METHOD_OUTLINEENTRY.html */
// func TestSetNoteLink(t *testing.T) {
// 	_, err := outlineentry.SetNoteLink( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETURL_METHOD_OUTLINEENTRY.html */
func TestSetURL(t *testing.T) {
	_, err := outlineentry.SetURL("http://notesoutlineentry_test.com")
	require.Nil(t, err)
}
