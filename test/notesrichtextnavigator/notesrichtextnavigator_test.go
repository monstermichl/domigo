/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTNAVIGATOR_CLASS.html */
package notesrichtextnavigator_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var richtextnavigator domigo.NotesRichTextNavigator

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		doc, err := db.CreateDocument()
		defer doc.Release()

		if err != nil {
			return "Document could not be created", err
		}
		richtextitem, err := doc.CreateRichTextItem("TestItem")
		defer richtextitem.Release()

		if err != nil {
			return "RichText item could not be created", err
		}
		richtextnavigator, err = richtextitem.CreateNavigator()
		defer richtextnavigator.Release()

		if err != nil {
			return "RichText navigator could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_RTNAV.html */
func TestClone(t *testing.T) {
	_, err := richtextnavigator.Clone()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTELEMENT_METHOD_RTNAV.html */
func TestFindFirstElement(t *testing.T) {
	_, err := richtextnavigator.FindFirstElement(domigo.NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_SECTION)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTSTRING_METHOD_RTNAV.html */
func TestFindFirstString(t *testing.T) {
	_, err := richtextnavigator.FindFirstString("Test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDLASTELEMENT_METHOD_RTNAV.html */
func TestFindLastElement(t *testing.T) {
	_, err := richtextnavigator.FindLastElement(domigo.NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_SECTION)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTELEMENT_METHOD_RTNAV.html */
func TestFindNextElement(t *testing.T) {
	_, err := richtextnavigator.FindNextElement()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTSTRING_METHOD_RTNAV.html */
func TestFindNextString(t *testing.T) {
	_, err := richtextnavigator.FindNextString("Test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHELEMENT_METHOD_RTNAV.html */
func TestFindNthElement(t *testing.T) {
	_, err := richtextnavigator.FindNthElement(domigo.NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_SECTION)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETELEMENT_METHOD_RTNAV.html */
// func TestGetElement(t *testing.T) {
// 	_, err := richtextnavigator.GetElement()
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTELEMENT_METHOD_RTNAV.html */
// func TestGetFirstElement(t *testing.T) {
// 	_, err := richtextnavigator.GetFirstElement( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTELEMENT_METHOD_RTNAV.html */
// func TestGetLastElement(t *testing.T) {
// 	_, err := richtextnavigator.GetLastElement( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTELEMENT_METHOD_RTNAV.html */
// func TestGetNextElement(t *testing.T) {
// 	_, err := richtextnavigator.GetNextElement( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHELEMENT_METHOD_RTNAV.html */
// func TestGetNthElement(t *testing.T) {
// 	_, err := richtextnavigator.GetNthElement( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCHAROFFSET_METHOD_RTNAV.html */
func TestSetCharOffset(t *testing.T) {
	err := richtextnavigator.SetCharOffset(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPOSITION_METHOD_RTNAV.html */
// func TestSetPosition(t *testing.T) {
// 	err := richtextnavigator.SetPosition( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPOSITIONATEND_METHOD_RTNAV.html */
// func TestSetPositionAtEnd(t *testing.T) {
// 	err := richtextnavigator.SetPositionAtEnd( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }
