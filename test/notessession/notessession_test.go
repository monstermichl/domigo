package notessession_test

import (
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/require"
)

var session notessession.NotesSession

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ = notessession.Initialize()
	m.Run()
	session.Release()
}

/* --------------------------------- Methods ------------------------------------ */
func TestNotesSessionNew(t *testing.T) {
	sTmp, err := notessession.Initialize()
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
