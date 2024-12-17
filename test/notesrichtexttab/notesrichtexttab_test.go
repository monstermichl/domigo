/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RICHTEXTTAB_CLASS.html */
package notesrichtexttab_test

import (
	"testing"

	"github.com/monstermichl/domigo/domino/notesrichtexttab"
	"github.com/monstermichl/domigo/domino/notessession"

	"github.com/stretchr/testify/require"
)

var richtexttab notesrichtexttab.NotesRichTextTab

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	style, _ := session.CreateRichTextParagraphStyle()

	style.SetTab(0)

	tabs, _ := style.Tabs()
	richtexttab = tabs[0]

	defer richtexttab.Release()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_4695.html */
func TestPosition(t *testing.T) {
	_, err := richtexttab.Position()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_7073.html */
func TestType(t *testing.T) {
	_, err := richtexttab.Type()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEAR_METHOD_1029.html */
func TestClear(t *testing.T) {
	err := richtexttab.Clear()
	require.Nil(t, err)
}
