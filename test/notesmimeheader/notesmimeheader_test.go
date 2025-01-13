/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESMIMEHEADER_CLASS.html */
package notesmimeheader_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var mimeheader domigo.NotesMIMEHeader

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		doc, err := db.CreateDocument()
		defer doc.Release()

		if err != nil {
			return "Document could not be created", err
		}
		mimeentity, err := doc.CreateMIMEEntity()
		defer mimeentity.Release()

		if err != nil {
			return "MIME entity could not be created", err
		}
		mimeheader, err = mimeentity.CreateHeader("testHeader")
		defer mimeheader.Release()

		if err != nil {
			return "NotesMIMEHeader could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERNAME_PROPERTY_MIMEHEADER.html */
func TestHeaderName(t *testing.T) {
	_, err := mimeheader.HeaderName()
	require.Nil(t, err)
}

/* TODO: Access type for Parent could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_MIMEHEADER.html */
func TestParent(t *testing.T) {
	_, err := mimeheader.Parent()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDVALTEXT_METHOD_MIMEHEADER.html */
func TestAddValText(t *testing.T) {
	_, err := mimeheader.AddValText("testValueText")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETHEADERVAL_METHOD_MIMEHEADER.html */
func TestGetHeaderVal(t *testing.T) {
	_, err := mimeheader.GetHeaderVal( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETHEADERVALANDPARAMS_METHOD_MIMEHEADER.html */
func TestGetHeaderValAndParams(t *testing.T) {
	_, err := mimeheader.GetHeaderValAndParams( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARAMVAL_METHOD_MIMEHEADER.html */
func TestGetParamVal(t *testing.T) {
	_, err := mimeheader.GetParamVal("testParam")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHEADERVAL_METHOD_MIMEHEADER.html */
func TestSetHeaderVal(t *testing.T) {
	_, err := mimeheader.SetHeaderVal("testHeaderValue")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHEADERVALANDPARAMS_METHOD_MIMEHEADER.html */
func TestSetHeaderValAndParams(t *testing.T) {
	_, err := mimeheader.SetHeaderValAndParams("testHeaderValAndParams")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPARAMVAL_METHOD_MIMEHEADER.html */
func TestSetParamVal(t *testing.T) {
	_, err := mimeheader.SetParamVal("testParamName", "testParamVal")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_MIMEHEADER.html */
func TestRemove(t *testing.T) {
	err := mimeheader.Remove()
	require.Nil(t, err)
}
