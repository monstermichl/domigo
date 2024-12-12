/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACL_CLASS.html */
package notesacl_test

import (
	"domigo/domino/notesacl"
	"domigo/domino/notessession"
	testhelpers "domigo/test/helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

const TEST_ROLE = "TestRole"
const TEST_ENTRY_NAME = "TestUser"
const TEST_ENTRY_LEVEL = 2

var acl notesacl.NotesACL

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	defer session.Release()

	db, _ := testhelpers.CreateTestDatabase(session)
	defer db.Release()
	defer db.Remove()

	acl, _ = db.ACL()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func TestAdministrationServer(t *testing.T) {
	_, err := acl.AdministrationServer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func TestSetAdministrationServer(t *testing.T) {
	s, err := acl.AdministrationServer()
	require.Nil(t, err)

	err = acl.SetAdministrationServer(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func TestInternetLevel(t *testing.T) {
	_, err := acl.InternetLevel()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func TestSetInternetLevel(t *testing.T) {
	s, err := acl.InternetLevel()
	require.Nil(t, err)

	err = acl.SetInternetLevel(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func TestIsAdminNames(t *testing.T) {
	_, err := acl.IsAdminNames()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func TestSetIsAdminNames(t *testing.T) {
	s, err := acl.IsAdminNames()
	require.Nil(t, err)

	err = acl.SetIsAdminNames(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func TestIsAdminReaderAuthor(t *testing.T) {
	_, err := acl.IsAdminReaderAuthor()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func TestSetIsAdminReaderAuthor(t *testing.T) {
	s, err := acl.IsAdminReaderAuthor()
	require.Nil(t, err)

	err = acl.SetIsAdminReaderAuthor(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func TestIsExtendedAccess(t *testing.T) {
	_, err := acl.IsExtendedAccess()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func TestSetIsExtendedAccess(t *testing.T) {
	s, err := acl.IsExtendedAccess()
	require.Nil(t, err)

	err = acl.SetIsExtendedAccess(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROLES_PROPERTY_ACL.html */
func TestRoles(t *testing.T) {
	_, err := acl.Roles()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func TestUniformAccess(t *testing.T) {
	_, err := acl.UniformAccess()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func TestSetUniformAccess(t *testing.T) {
	s, err := acl.UniformAccess()
	require.Nil(t, err)

	err = acl.SetUniformAccess(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDROLE_METHOD.html */
func TestAddRole(t *testing.T) {
	err := acl.AddRole(TEST_ROLE)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEACLENTRY_METHOD.html */
func TestCreateACLEntry(t *testing.T) {
	_, err := acl.CreateACLEntry(TEST_ENTRY_NAME, TEST_ENTRY_LEVEL)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD.html */
func TestGetEntry(t *testing.T) {
	_, err := acl.GetEntry(TEST_ENTRY_NAME)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTENTRY_METHOD.html */
func TestGetFirstEntry(t *testing.T) {
	_, err := acl.GetFirstEntry()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_METHOD.html */
func TestGetNextEntry(t *testing.T) {
	first, _ := acl.GetFirstEntry()
	_, err := acl.GetNextEntry(first)

	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEROLE_METHOD.html */
func TestRenameRole(t *testing.T) {
	err := acl.RenameRole(TEST_ENTRY_NAME, TEST_ENTRY_NAME)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEACLENTRY_METHOD_ACL_COM.html */
// func TestRemoveACLEntry(t *testing.T) {
// 	err := acl.RemoveACLEntry(TEST_ENTRY_NAME)
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEROLE_METHOD.html */
// func TestDeleteRole(t *testing.T) {
// 	err := acl.DeleteRole(TEST_ROLE)
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_ACL.html */
func TestSave(t *testing.T) {
	err := acl.Save()
	require.Nil(t, err)
}
