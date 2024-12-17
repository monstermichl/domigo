/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTRANGE_CLASS.html */
package notesrichtextrange_test

import (
	"testing"

	"github.com/monstermichl/domigo/domino/notesrichtextitem"
	"github.com/monstermichl/domigo/domino/notesrichtextrange"
	"github.com/monstermichl/domigo/domino/notessession"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var richtextitem notesrichtextitem.NotesRichTextItem
var richtextrange notesrichtextrange.NotesRichTextRange

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	db, _ := testhelpers.CreateTestDatabase(session)
	doc, _ := db.CreateDocument()
	richtextitem, _ = doc.CreateRichTextItem("TestItem")
	richtextrange, _ = richtextitem.CreateRange()

	defer richtextrange.Release()
	defer db.Release()
	defer db.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAVIGATOR_PROPERTY_RTRANGE.html */
func TestNavigator(t *testing.T) {
	_, err := richtextrange.Navigator()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STYLE_PROPERTY_RTRANGE.html */
func TestStyle(t *testing.T) {
	_, err := richtextrange.Style()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXTPARAGRAPH_PROPERTY_RTRANGE.html */
func TestTextParagraph(t *testing.T) {
	_, err := richtextrange.TextParagraph()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXTRUN_PROPERTY_RTRANGE.html */
func TestTextRun(t *testing.T) {
	_, err := richtextrange.TextRun()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_RTRANGE.html */
func TestType(t *testing.T) {
	_, err := richtextrange.Type()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_RTRANGE.html */
func TestClone(t *testing.T) {
	_, err := richtextrange.Clone()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDANDREPLACE_METHOD_RTRANGE.html */
func TestFindAndReplace(t *testing.T) {
	_, err := richtextrange.FindAndReplace("Test", "Test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_RTRANGE.html */
func TestRemove(t *testing.T) {
	err := richtextrange.Remove()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESET_METHOD_RTRANGE.html */
func TestReset(t *testing.T) {
	err := richtextrange.Reset()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETBEGIN_METHOD_RTRANGE.html */
// func TestSetBegin(t *testing.T) {
// 	err := richtextrange.SetBegin( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETEND_METHOD_RTRANGE.html */
// func TestSetEnd(t *testing.T) {
// 	err := richtextrange.SetEnd( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETSTYLE_METHOD_RTRANGE.html */
// func TestSetStyle(t *testing.T) {
// 	err := richtextrange.SetStyle( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }
