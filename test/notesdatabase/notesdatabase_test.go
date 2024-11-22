package notesform_test

import (
	"domigo/domino/notesdatabase"
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/assert"
)

var session notessession.NotesSession

func mustGetDb(t *testing.T) notesdatabase.NotesDatabase {
	db, err := session.GetDatabase("", "GoInterface.nsf")

	assert.Nil(t, err)
	return db
}

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ = notessession.New()
	m.Run()
	session.Release()
}

func TestNotesDatabaseGetDatabase(t *testing.T) {
	mustGetDb(t)
}

func TestNotesDatabaseACL(t *testing.T) {
	_, err := mustGetDb(t).ACL()
	assert.Nil(t, err)
}

func TestNotesDatabaseForms(t *testing.T) {
	_, err := mustGetDb(t).Forms()
	assert.Nil(t, err)
}

func TestNotesDatabaseIsDesignLockingEnabled(t *testing.T) {
	_, err := mustGetDb(t).IsDesignLockingEnabled()
	assert.Nil(t, err)
}

func TestNotesDatabaseSetIsDesignLockingEnabled(t *testing.T) {
	err := mustGetDb(t).SetIsDesignLockingEnabled(false)
	assert.Nil(t, err)
}

func TestNotesDatabaseTitle(t *testing.T) {
	title, err := mustGetDb(t).Title()

	assert.Nil(t, err)
	assert.Equal(t, "GoInterface", title)
}
