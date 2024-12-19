/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACLENTRY_CLASS.html */
package notesaclentry_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

const TEST_ROLE = "TestRole"
const TEST_ENTRY_NAME = "TestUser"
const TEST_ENTRY_LEVEL = 6

var aclentry domigo.NotesACLEntry

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := domigo.Initialize()
	db, _ := testhelpers.CreateTestDatabase(session)
	acl, _ := db.ACL()

	acl.AddRole(TEST_ROLE)
	aclentry, _ = acl.CreateACLEntry(TEST_ENTRY_NAME, TEST_ENTRY_LEVEL)

	defer db.Release()
	defer db.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORNOCREATE_PROPERTY.html */
func TestCanCreateDocuments(t *testing.T) {
	_, err := aclentry.CanCreateDocuments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORNOCREATE_PROPERTY.html */
func TestSetCanCreateDocuments(t *testing.T) {
	s, err := aclentry.CanCreateDocuments()
	require.Nil(t, err)

	err = aclentry.SetCanCreateDocuments(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATELSORJAVAAGENT_PROPERTY_6760.html */
func TestCanCreateLSOrJavaAgent(t *testing.T) {
	_, err := aclentry.CanCreateLSOrJavaAgent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATELSORJAVAAGENT_PROPERTY_6760.html */
func TestSetCanCreateLSOrJavaAgent(t *testing.T) {
	s, err := aclentry.CanCreateLSOrJavaAgent()
	require.Nil(t, err)

	err = aclentry.SetCanCreateLSOrJavaAgent(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEAGENT_PROPERTY.html */
func TestCanCreatePersonalAgent(t *testing.T) {
	_, err := aclentry.CanCreatePersonalAgent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEAGENT_PROPERTY.html */
func TestSetCanCreatePersonalAgent(t *testing.T) {
	s, err := aclentry.CanCreatePersonalAgent()
	require.Nil(t, err)

	err = aclentry.SetCanCreatePersonalAgent(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEFOLDER_PROPERTY.html */
func TestCanCreatePersonalFolder(t *testing.T) {
	_, err := aclentry.CanCreatePersonalFolder()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEFOLDER_PROPERTY.html */
func TestSetCanCreatePersonalFolder(t *testing.T) {
	s, err := aclentry.CanCreatePersonalFolder()
	require.Nil(t, err)

	err = aclentry.SetCanCreatePersonalFolder(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATESHAREDFOLDER_PROPERTY_7622.html */
func TestCanCreateSharedFolder(t *testing.T) {
	_, err := aclentry.CanCreateSharedFolder()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATESHAREDFOLDER_PROPERTY_7622.html */
func TestSetCanCreateSharedFolder(t *testing.T) {
	s, err := aclentry.CanCreateSharedFolder()
	require.Nil(t, err)

	err = aclentry.SetCanCreateSharedFolder(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNODELETE_PROPERTY.html */
func TestCanDeleteDocuments(t *testing.T) {
	_, err := aclentry.CanDeleteDocuments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNODELETE_PROPERTY.html */
func TestSetCanDeleteDocuments(t *testing.T) {
	s, err := aclentry.CanDeleteDocuments()
	require.Nil(t, err)

	err = aclentry.SetCanDeleteDocuments(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANREPLICATEORCOPYDOCUMENTS_PROPERTY_ACLENTRY.html */
func TestCanReplicateOrCopyDocuments(t *testing.T) {
	_, err := aclentry.CanReplicateOrCopyDocuments()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANREPLICATEORCOPYDOCUMENTS_PROPERTY_ACLENTRY.html */
func TestSetCanReplicateOrCopyDocuments(t *testing.T) {
	s, err := aclentry.CanReplicateOrCopyDocuments()
	require.Nil(t, err)

	err = aclentry.SetCanReplicateOrCopyDocuments(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_88.html */
func TestIsAdminReaderAuthor(t *testing.T) {
	_, err := aclentry.IsAdminReaderAuthor()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_88.html */
func TestSetIsAdminReaderAuthor(t *testing.T) {
	s, err := aclentry.IsAdminReaderAuthor()
	require.Nil(t, err)

	err = aclentry.SetIsAdminReaderAuthor(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINSERVER_PROPERTY_2946.html */
func TestIsAdminServer(t *testing.T) {
	_, err := aclentry.IsAdminServer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINSERVER_PROPERTY_2946.html */
func TestSetIsAdminServer(t *testing.T) {
	s, err := aclentry.IsAdminServer()
	require.Nil(t, err)

	err = aclentry.SetIsAdminServer(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISGROUP_PROPERTY_4731.html */
func TestIsGroup(t *testing.T) {
	_, err := aclentry.IsGroup()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISGROUP_PROPERTY_4731.html */
func TestSetIsGroup(t *testing.T) {
	s, err := aclentry.IsGroup()
	require.Nil(t, err)

	err = aclentry.SetIsGroup(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPERSON_PROPERTY_651.html */
func TestIsPerson(t *testing.T) {
	_, err := aclentry.IsPerson()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPERSON_PROPERTY_651.html */
func TestSetIsPerson(t *testing.T) {
	s, err := aclentry.IsPerson()
	require.Nil(t, err)

	err = aclentry.SetIsPerson(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICREADER_PROPERTY.html */
func TestIsPublicReader(t *testing.T) {
	_, err := aclentry.IsPublicReader()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICREADER_PROPERTY.html */
func TestSetIsPublicReader(t *testing.T) {
	s, err := aclentry.IsPublicReader()
	require.Nil(t, err)

	err = aclentry.SetIsPublicReader(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICWRITER_PROPERTY.html */
func TestIsPublicWriter(t *testing.T) {
	_, err := aclentry.IsPublicWriter()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICWRITER_PROPERTY.html */
func TestSetIsPublicWriter(t *testing.T) {
	s, err := aclentry.IsPublicWriter()
	require.Nil(t, err)

	err = aclentry.SetIsPublicWriter(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSERVER_PROPERTY_1888.html */
func TestIsServer(t *testing.T) {
	_, err := aclentry.IsServer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSERVER_PROPERTY_1888.html */
func TestSetIsServer(t *testing.T) {
	s, err := aclentry.IsServer()
	require.Nil(t, err)

	err = aclentry.SetIsServer(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY.html */
func TestLevel(t *testing.T) {
	_, err := aclentry.Level()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY.html */
func TestSetLevel(t *testing.T) {
	s, err := aclentry.Level()
	require.Nil(t, err)

	err = aclentry.SetLevel(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ACLENTRY.html */
func TestName(t *testing.T) {
	_, err := aclentry.Name()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ACLENTRY.html */
func TestSetName(t *testing.T) {
	s, err := aclentry.Name()
	require.Nil(t, err)

	err = aclentry.SetName(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOBJECT_PROPERTY.html */
func TestNameObject(t *testing.T) {
	_, err := aclentry.NameObject()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROLES_PROPERTY_ACLENTRY.html */
func TestRoles(t *testing.T) {
	_, err := aclentry.Roles()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERTYPE_PROPERTY_ACLENTRY.html */
func TestUserType(t *testing.T) {
	_, err := aclentry.UserType()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERTYPE_PROPERTY_ACLENTRY.html */
func TestSetUserType(t *testing.T) {
	s, err := aclentry.UserType()
	require.Nil(t, err)

	err = aclentry.SetUserType(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLEROLE_METHOD.html */
func TestDisableRole(t *testing.T) {
	err := aclentry.DisableRole(TEST_ROLE)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENABLEROLE_METHOD.html */
func TestEnableRole(t *testing.T) {
	err := aclentry.EnableRole(TEST_ROLE)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROLEENABLED_METHOD.html */
func TestIsRoleEnabled(t *testing.T) {
	_, err := aclentry.IsRoleEnabled(TEST_ROLE)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_ACLENTRY.html */
func TestRemove(t *testing.T) {
	err := aclentry.Remove()
	require.Nil(t, err)
}
