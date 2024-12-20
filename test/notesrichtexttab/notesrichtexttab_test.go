/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RICHTEXTTAB_CLASS.html */
package notesrichtexttab_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"
	"github.com/stretchr/testify/require"
)

var richtexttab domigo.NotesRichTextTab

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		style, err := session.CreateRichTextParagraphStyle()
		defer style.Release()

		if err != nil {
			return "RichText style could not be created", err
		}
		err = style.SetTab(0)

		if err != nil {
			return "Tab could not be set", err
		}
		tabs, err := style.Tabs()

		if err != nil {
			return "Tabs could not be retrieved", err
		} else {
			for _, t := range tabs {
				defer t.Release()
			}
		}
		richtexttab = tabs[0]

		m.Run()
		return "", nil
	})
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
