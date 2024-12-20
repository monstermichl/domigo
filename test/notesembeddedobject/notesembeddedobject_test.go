/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESEMBEDDEDOBJECT_CLASS.html */
package notesembeddedobject_test

import (
	"path/filepath"
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var embeddedobject domigo.NotesEmbeddedObject

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		doc, err := db.CreateDocument()
		defer doc.Release()

		if err != nil {
			return "Document could not be created", err
		}
	
		rti, err := doc.CreateRichTextItem("RTI")
		defer rti.Release()

		if err != nil {
			return "Richtext item could not be created", err
		}

		path, err := filepath.Abs("test.md")
		embeddedobject, _ = rti.EmbedObject(domigo.NOTESEMBEDDEDOBJECT_EMBED_ATTACHMENT, "", path)
		defer embeddedobject.Release()

		if err != nil {
			return "Embedded object could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLASS_PROPERTY.html */
func TestClass(t *testing.T) {
	_, err := embeddedobject.Class()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILECREATED_PROPERTY.html */
func TestFileCreated(t *testing.T) {
	_, err := embeddedobject.FileCreated()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEENCODING_PROPERTY.html */
func TestFileEncoding(t *testing.T) {
	_, err := embeddedobject.FileEncoding()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEMODIFIED_PROPERTY.html */
func TestFileModified(t *testing.T) {
	_, err := embeddedobject.FileModified()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILESIZE_PROPERTY.html */
func TestFileSize(t *testing.T) {
	_, err := embeddedobject.FileSize()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func TestFitBelowFields(t *testing.T) {
	_, err := embeddedobject.FitBelowFields()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func TestSetFitBelowFields(t *testing.T) {
	s, err := embeddedobject.FitBelowFields()
	require.Nil(t, err)

	err = embeddedobject.SetFitBelowFields(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func TestFitToWindow(t *testing.T) {
	_, err := embeddedobject.FitToWindow()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func TestSetFitToWindow(t *testing.T) {
	s, err := embeddedobject.FitToWindow()
	require.Nil(t, err)

	err = embeddedobject.SetFitToWindow(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OBJECT.html */
func TestName(t *testing.T) {
	_, err := embeddedobject.Name()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OBJECT_PROPERTY.html */
func TestObject(t *testing.T) {
	_, err := embeddedobject.Object()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func TestRunReadOnly(t *testing.T) {
	_, err := embeddedobject.RunReadOnly()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func TestSetRunReadOnly(t *testing.T) {
	s, err := embeddedobject.RunReadOnly()
	require.Nil(t, err)

	err = embeddedobject.SetRunReadOnly(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SOURCE_PROPERTY.html */
func TestSource(t *testing.T) {
	_, err := embeddedobject.Source()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OBJECT.html */
func TestType(t *testing.T) {
	_, err := embeddedobject.Type()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERBS_PROPERTY.html */
func TestVerbs(t *testing.T) {
	_, err := embeddedobject.Verbs()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACTIVATE_METHOD.html */
func TestActivate(t *testing.T) {
	_, err := embeddedobject.Activate(false)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOVERB_METHOD.html */
func TestDoVerb(t *testing.T) {
	_, err := embeddedobject.DoVerb("test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXTRACTFILE_METHOD.html */
func TestExtractFile(t *testing.T) {
	err := embeddedobject.ExtractFile("loaded-test.md")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_OBJECT.html */
func TestRemove(t *testing.T) {
	err := embeddedobject.Remove()
	require.Nil(t, err)
}
