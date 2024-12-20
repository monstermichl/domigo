package notessession_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"
	"github.com/stretchr/testify/require"
)

var session domigo.NotesSession

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(sessionTemp domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		session = sessionTemp
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Methods ------------------------------------ */
func TestNotesSessionNew(t *testing.T) {
	sTmp, err := domigo.Initialize()
	require.Nil(t, err)

	session = sTmp
}

func TestNotesSessionGetDatabase(t *testing.T) {
	_, err := session.GetDatabase("", "GoInterface.nsf")
	require.Nil(t, err)
}

func TestNotesSessionGetRelease(t *testing.T) {
	session.Release()
}
