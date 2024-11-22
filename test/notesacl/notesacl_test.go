package notesacl_test

import (
	"domigo/domino/notesacl"
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/assert"
)

var acl notesacl.NotesACL

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.New()
	db, _ := session.GetDatabase("", "GoInterface.nsf")
	aclTmp, _ := db.ACL()
	acl = aclTmp

	m.Run()
	session.Release()
}

func TestNotesACLAdministrationServer(t *testing.T) {
	_, err := acl.AdministrationServer()
	assert.Nil(t, err)
}

func TestNotesACLInternetLevel(t *testing.T) {
	_, err := acl.InternetLevel()
	assert.Nil(t, err)
}

func TestNotesACLIsAdminNames(t *testing.T) {
	_, err := acl.IsAdminNames()
	assert.Nil(t, err)
}

func TestNotesACLIsAdminReaderAuthor(t *testing.T) {
	_, err := acl.IsAdminReaderAuthor()
	assert.Nil(t, err)
}

func TestNotesACLIsExtendedAccess(t *testing.T) {
	_, err := acl.IsExtendedAccess()
	assert.Nil(t, err)
}

func TestNotesACLParent(t *testing.T) {
	p, err := acl.Parent()

	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestNotesACLRoles(t *testing.T) {
	_, err := acl.Roles()
	assert.Nil(t, err)
}

func TestNotesACLSetAdministrationServer(t *testing.T) {
	s, err := acl.AdministrationServer()
	assert.Nil(t, err)

	err = acl.SetAdministrationServer(s)
	assert.Nil(t, err)
}

func TestNotesACLSetInternetLevel(t *testing.T) {
	l, err := acl.InternetLevel()
	assert.Nil(t, err)

	err = acl.SetInternetLevel(l)
	assert.Nil(t, err)
}

func TestNotesACLSetIsAdminNames(t *testing.T) {
	i, err := acl.IsAdminNames()
	assert.Nil(t, err)

	err = acl.SetIsAdminNames(i)
	assert.Nil(t, err)
}

func TestNotesACLSetIsAdminReaderAuthor(t *testing.T) {
	i, err := acl.IsAdminReaderAuthor()
	assert.Nil(t, err)

	err = acl.SetIsAdminReaderAuthor(i)
	assert.Nil(t, err)
}

func TestNotesACLSetIsExtendedAccess(t *testing.T) {
	i, err := acl.IsExtendedAccess()
	assert.Nil(t, err)

	err = acl.SetIsExtendedAccess(i)
	assert.Nil(t, err)
}

func TestNotesACLSetUniformAccess(t *testing.T) {
	u, err := acl.UniformAccess()
	assert.Nil(t, err)

	err = acl.SetUniformAccess(u)
	assert.Nil(t, err)
}

func TestNotesACLAddRole(t *testing.T) {
	err := acl.AddRole("TestRole")
	assert.Nil(t, err)
}

func TestNotesACLRenameRole(t *testing.T) {
	err := acl.RenameRole("TestRole", "TestRole1")
	assert.Nil(t, err)
}

func TestNotesACLDeleteRole(t *testing.T) {
	err := acl.DeleteRole("TestRole1")
	assert.Nil(t, err)
}

func TestNotesACLCreateACLEntry(t *testing.T) {
	err := acl.CreateACLEntry("TestGroup", 0)
	assert.Nil(t, err)
}

func TestNotesACLGetEntry(t *testing.T) {
	_, err := acl.GetEntry("TestGroup")
	assert.Nil(t, err)
}

func TestNotesACLGetFirstEntry(t *testing.T) {
	_, err := acl.GetFirstEntry()
	assert.Nil(t, err)
}

/* TODO: Find out how to implement GetNextEntry. */
// func TestNotesACLGetNextEntry(t *testing.T) {
// 	_, err := acl.GetNextEntry()
// 	assert.Nil(t, err)
// }

func TestNotesACLRemoveACLEntry(t *testing.T) {
	err := acl.RemoveACLEntry("TestGroup")
	assert.Nil(t, err)
}

func TestNotesACLSave(t *testing.T) {
	err := acl.Save()
	assert.Nil(t, err)
}
