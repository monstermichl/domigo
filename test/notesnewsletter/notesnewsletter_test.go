/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNEWSLETTER_CLASS.html */
package notesnewsletter_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var db domigo.NotesDatabase
var newsletter domigo.NotesNewsletter

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, dbTemp domigo.NotesDatabase) (string, error) {
		db = dbTemp

		collection, err := db.CreateDocumentCollection()
		defer collection.Release()

		if err != nil {
			return "Document collection could not be created", err
		}
		newsletter, err = session.CreateNewsletter(collection)
		defer newsletter.Release()

		if err != nil {
			return "Newsletter could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSCORE_PROPERTY.html */
func TestDoScore(t *testing.T) {
	_, err := newsletter.DoScore()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSCORE_PROPERTY.html */
func TestSetDoScore(t *testing.T) {
	s, err := newsletter.DoScore()
	require.Nil(t, err)

	err = newsletter.SetDoScore(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSUBJECT_PROPERTY.html */
func TestDoSubject(t *testing.T) {
	_, err := newsletter.DoSubject()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSUBJECT_PROPERTY.html */
func TestSetDoSubject(t *testing.T) {
	s, err := newsletter.DoSubject()
	require.Nil(t, err)

	err = newsletter.SetDoSubject(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBJECTITEMNAME_PROPERTY.html */
func TestSubjectItemName(t *testing.T) {
	_, err := newsletter.SubjectItemName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBJECTITEMNAME_PROPERTY.html */
func TestSetSubjectItemName(t *testing.T) {
	s, err := newsletter.SubjectItemName()
	require.Nil(t, err)

	err = newsletter.SetSubjectItemName(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMATDOCUMENT_METHOD.html */
func TestFormatDocument(t *testing.T) {
	_, err := newsletter.FormatDocument(db, 1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMATMSGWITHDOCLINKS_METHOD.html */
func TestFormatMsgWithDoclinks(t *testing.T) {
	_, err := newsletter.FormatMsgWithDoclinks(db)
	require.Nil(t, err)
}
