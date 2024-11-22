package notessession_test

import (
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/assert"
)

var session notessession.NotesSession

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ = notessession.New()
	m.Run()
	session.Release()
}

func TestNotesSessionNew(t *testing.T) {
	sTmp, err := notessession.New()
	assert.Nil(t, err)

	session = sTmp
}

func TestNotesSessionGetDatabase(t *testing.T) {
	_, err := session.GetDatabase("", "GoInterface.nsf")
	assert.Nil(t, err)
}

func TestNotesSessionGetRelease(t *testing.T) {
	session.Release()
}
