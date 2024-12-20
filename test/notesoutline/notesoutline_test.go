/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESOUTLINE_CLASS.html */
package notesoutline_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

const TEST_OUTLINE_NAME = "test-outline"

var database domigo.NotesDatabase
var outline domigo.NotesOutline

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error
		database = db

		outline, err = database.CreateOutline(TEST_OUTLINE_NAME)
		defer outline.Release()

		if err != nil {
			return "Outline could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINE.html */
func TestAlias(t *testing.T) {
	_, err := outline.Alias()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINE.html */
func TestSetAlias(t *testing.T) {
	s, err := outline.Alias()
	require.Nil(t, err)

	err = outline.SetAlias(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY_OUTLINE.html */
func TestComment(t *testing.T) {
	_, err := outline.Comment()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY_OUTLINE.html */
func TestSetComment(t *testing.T) {
	s, err := outline.Comment()
	require.Nil(t, err)

	err = outline.SetComment(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OUTLINE.html */
func TestName(t *testing.T) {
	_, err := outline.Name()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OUTLINE.html */
func TestSetName(t *testing.T) {
	s, err := outline.Name()
	require.Nil(t, err)

	err = outline.SetName(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEENTRY_METHOD_OUTLINE.html */
func TestCreateEntry(t *testing.T) {
	_, err := outline.CreateEntry("outline-entry-name-1")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEENTRYFROM_METHOD_MEMDEF_NOTESOUTLINE.html */
func TestCreateEntryFrom(t *testing.T) {
	entry, err := outline.CreateEntry("outline-entry-name-2")
	require.Nil(t, err, "Creation failed")

	_, err = outline.CreateEntryFrom(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_NOTESOUTLINE.html */
func TestGetChild(t *testing.T) {
	entry, err := outline.CreateEntry("outline-entry-name-3")
	require.Nil(t, err, "Creation failed")

	_, err = outline.GetChild(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_OUTLINE.html */
func TestGetFirst(t *testing.T) {
	_, err := outline.GetFirst()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_OUTLINE.html */
func TestGetLast(t *testing.T) {
	_, err := outline.GetLast()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_OUTLINE.html */
func TestGetNext(t *testing.T) {
	entry, err := outline.CreateEntry("outline-entry-name-4")
	require.Nil(t, err, "Creation failed")

	_, err = outline.GetNext(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_OUTLINE.html */
func TestGetNextSibling(t *testing.T) {
	entry, err := outline.CreateEntry("outline-entry-name-5")
	require.Nil(t, err, "Creation failed")

	_, err = outline.GetNextSibling(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_OUTLINE.html */
func TestGetParent(t *testing.T) {
	entry, err := outline.CreateEntry("outline-entry-name-6")
	require.Nil(t, err, "Creation failed")

	_, err = outline.GetParent(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_OUTLINE.html */
func TestGetPrev(t *testing.T) {
	entry, err := outline.CreateEntry("outline-entry-name-7")
	require.Nil(t, err, "Creation failed")

	_, err = outline.GetPrev(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_OUTLINE.html */
func TestGetPrevSibling(t *testing.T) {
	entry, err := outline.CreateEntry("outline-entry-name-8")
	require.Nil(t, err, "Creation failed")

	_, err = outline.GetPrevSibling(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEENTRY_METHOD_OUTLINE.html */
func TestMoveEntry(t *testing.T) {
	firstEntry, err := outline.CreateEntry("outline-entry-name-9")
	require.Nil(t, err, "Creation failed")

	secondEntry, err := outline.CreateEntry("outline-entry-name-10")
	require.Nil(t, err, "Creation failed")

	err = outline.MoveEntry(firstEntry, secondEntry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEENTRY_METHOD_OUTLINE.html */
func TestRemoveEntry(t *testing.T) {
	entry, err := outline.CreateEntry("outline-entry-name-11")
	require.Nil(t, err, "Creation failed")

	err = outline.RemoveEntry(entry)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_OUTLINE.html */
func TestSave(t *testing.T) {
	err := outline.Save()
	require.Nil(t, err)
}
