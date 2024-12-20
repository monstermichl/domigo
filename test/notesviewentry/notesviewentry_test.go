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
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		document, err := db.CreateDocument()
		defer document.Release()

		if err != nil {
			return "Document could not becreated", err
		}
		_, err = document.ReplaceItemValue("testField", fmt.Sprint(time.Now().Unix()))

		if err != nil {
			return "Test value could not be set", err
		}
		_, err = document.Save(true, false)

		if err != nil {
			return "Document could not be saved", err
		}
		view, err := db.GetView("TestView")
		defer view.Release()

		if err != nil {
			return "View could not beretrieved", err
		}
		viewentries, err := view.AllEntries()
		defer viewentries.Release()

		if err != nil {
			return "View entries could not be retrieved", err
		}
		count, err := viewentries.Count()

		if err != nil {
			return "View entries count could not be retrieved", err
		} else if count <= 0 {
			return "No view entries found", err
		}
		viewentry, err = viewentries.GetFirstEntry()

		if err != nil {
			return "View entry could not be retrieved", err
		}
		m.Run()
		return "", nil
	})
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
