package notesviewentry_test

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWENTRY_CLASS_2925.html */

import (
	"fmt"
	"testing"
	"time"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var viewentry domigo.NotesViewEntry

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	var info string

	session, err := domigo.Initialize()
	defer session.Release()

	defer func() {
		fmt.Println(err)
		fmt.Println(info)
	}()

	if err != nil {
		info = "Session could not beinitialized"
		return
	}

	db, err := testhelpers.CreateTestDatabase(session)
	defer db.Release()
	defer db.Remove()

	if err != nil {
		info = "Database could not becreated"
		return
	}

	document, err := db.CreateDocument()
	defer document.Release()

	if err != nil {
		info = "Document could not becreated"
		return
	}

	_, err = document.ReplaceItemValue("testField", fmt.Sprint(time.Now().Unix()))

	if err != nil {
		info = "Test value could not be set"
		return
	}

	_, err = document.Save(true, false)

	if err != nil {
		info = "Document could not be saved"
		return
	}

	view, err := db.GetView("TestView")
	defer view.Release()

	if err != nil {
		info = "View could not beretrieved"
		return
	}

	viewentries, err := view.AllEntries()
	defer viewentries.Release()

	if err != nil {
		info = "View entries could not be retrieved"
		return
	}

	count, err := viewentries.Count()

	if err != nil {
		info = "View entries count could not be retrieved"
		return
	} else if count <= 0 {
		info = "No view entries found"
		return
	}

	viewentry, err = viewentries.GetFirstEntry()

	if err != nil {
		info = "View entry could not be retrieved"
		return
	}

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHILDCOUNT_PROPERTY_8963.html */
func TestChildCount(t *testing.T) {
	_, err := viewentry.ChildCount()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNINDENTLEVEL_PROPERTY_VIEWENTRY.html */
func TestColumnIndentLevel(t *testing.T) {
	_, err := viewentry.ColumnIndentLevel()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNVALUES_PROPERTY_5887.html */
func TestColumnValues(t *testing.T) {
	_, err := viewentry.ColumnValues()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESCENDANTCOUNT_PROPERTY_2777.html */
func TestDescendantCount(t *testing.T) {
	_, err := viewentry.DescendantCount()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENT_PROPERTY_7936.html */
func TestDocument(t *testing.T) {
	_, err := viewentry.Document()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHSCORE_PROPERTY_2767.html */
func TestFTSearchScore(t *testing.T) {
	_, err := viewentry.FTSearchScore()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INDENTLEVEL_PROPERTY_8244.html */
func TestIndentLevel(t *testing.T) {
	_, err := viewentry.IndentLevel()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORY_PROPERTY_9511.html */
func TestIsCategory(t *testing.T) {
	_, err := viewentry.IsCategory()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFLICT_PROPERTY_1204.html */
func TestIsConflict(t *testing.T) {
	_, err := viewentry.IsConflict()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENT_PROPERTY_6058.html */
func TestIsDocument(t *testing.T) {
	_, err := viewentry.IsDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISTOTAL_PROPERTY_9664.html */
func TestIsTotal(t *testing.T) {
	_, err := viewentry.IsTotal()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALID_PROPERTY_2176.html */
func TestIsValid(t *testing.T) {
	_, err := viewentry.IsValid()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY_9904.html */
func TestNoteID(t *testing.T) {
	_, err := viewentry.NoteID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIBLINGCOUNT_PROPERTY_5642.html */
func TestSiblingCount(t *testing.T) {
	_, err := viewentry.SiblingCount()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_8372.html */
func TestUniversalID(t *testing.T) {
	_, err := viewentry.UniversalID()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOSITION_METHOD_8121.html */
func TestGetPosition(t *testing.T) {
	_, err := viewentry.GetPosition(";")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETREAD_METHOD_ENTRY.html */
func TestGetRead(t *testing.T) {
	_, err := viewentry.GetRead()
	require.Nil(t, err)
}
