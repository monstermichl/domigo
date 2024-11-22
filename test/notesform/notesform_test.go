package notesform_test

import (
	"domigo/domino/notesform"
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/assert"
)

var form notesform.NotesForm

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.New()
	db, _ := session.GetDatabase("", "GoInterface.nsf")
	forms, _ := db.Forms()
	form = forms[0]

	/* Prepare database for NotesForm tests. */
	dle, edle := db.IsDesignLockingEnabled()

	if edle == nil {
		db.SetIsDesignLockingEnabled(true)
	}
	m.Run()

	/* Restore original values. */
	if edle == nil {
		db.SetIsDesignLockingEnabled(dle)
	}
	session.Release()
}

func TestNotesFormAliases(t *testing.T) {
	_, err := form.Aliases()
	assert.Nil(t, err)
}

func TestNotesFormFields(t *testing.T) {
	_, err := form.Fields()
	assert.Nil(t, err)
}

func TestNotesFormUsers(t *testing.T) {
	_, err := form.FormUsers()
	assert.Nil(t, err)
}

func TestNotesFormGetFieldType(t *testing.T) {
	fields, err := form.Fields()
	assert.Nil(t, err)

	for _, f := range fields {
		_, err = form.GetFieldType(f)
		assert.Nil(t, err)
	}
}

func TestNotesFormHttpURL(t *testing.T) {
	_, err := form.HttpURL()
	assert.Nil(t, err)
}

func TestNotesFormIsSubForm(t *testing.T) {
	_, err := form.IsSubForm()
	assert.Nil(t, err)
}

func TestNotesFormLock(t *testing.T) {
	_, err := form.Lock([]string{}, false)
	assert.Nil(t, err)
}

func TestNotesFormLockHolders(t *testing.T) {
	_, err := form.LockHolders()
	assert.Nil(t, err)
}

func TestNotesFormLockProvisional(t *testing.T) {
	_, err := form.LockProvisional([]string{})
	assert.Nil(t, err)
}

func TestNotesFormName(t *testing.T) {
	name, err := form.Name()

	assert.Nil(t, err)
	assert.NotEmpty(t, name)
}

func TestNotesFormNotesURL(t *testing.T) {
	_, err := form.NotesURL()
	assert.Nil(t, err)
}

func TestNotesFormParent(t *testing.T) {
	_, err := form.Parent()
	assert.Nil(t, err)
}

func TestNotesFormNotesProtectReaders(t *testing.T) {
	_, err := form.ProtectReaders()
	assert.Nil(t, err)
}

func TestNotesFormProtectUsers(t *testing.T) {
	_, err := form.ProtectUsers()
	assert.Nil(t, err)
}

func TestNotesFormNotesReaders(t *testing.T) {
	_, err := form.Readers()
	assert.Nil(t, err)
}

func TestNotesFormNotesSetFormUsers(t *testing.T) {
	err := form.SetFormUsers([]string{})
	assert.Nil(t, err)
}

func TestNotesFormNotesSetProtectReaders(t *testing.T) {
	err := form.SetProtectReaders(false)
	assert.Nil(t, err)
}

func TestNotesFormNotesSetProtectUsers(t *testing.T) {
	err := form.SetProtectUsers(false)
	assert.Nil(t, err)
}

func TestNotesFormNotesSetReaders(t *testing.T) {
	err := form.SetReaders([]string{})
	assert.Nil(t, err)
}

func TestNotesFormNotesUnlock(t *testing.T) {
	err := form.UnLock()
	assert.Nil(t, err)
}
