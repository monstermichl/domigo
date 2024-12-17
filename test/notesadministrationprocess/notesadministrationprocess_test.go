/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESADMINISTRATIONPROCESS_CLASS.html */
package notesadministrationprocess_test

import (
	"testing"

	"github.com/monstermichl/domigo/domino/notesadministrationprocess"
	"github.com/monstermichl/domigo/domino/notessession"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var administrationprocess notesadministrationprocess.NotesAdministrationProcess

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	db, _ := testhelpers.CreateTestDatabase(session)
	administrationprocess, _ = session.CreateAdministrationProcess("")

	defer db.Release()
	defer db.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEAUTHORITYORG_PROPERTY_ADMINP.html */
func TestCertificateAuthorityOrg(t *testing.T) {
	_, err := administrationprocess.CertificateAuthorityOrg()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEAUTHORITYORG_PROPERTY_ADMINP.html */
func TestSetCertificateAuthorityOrg(t *testing.T) {
	s, err := administrationprocess.CertificateAuthorityOrg()
	require.Nil(t, err)

	err = administrationprocess.SetCertificateAuthorityOrg(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEEXPIRATION_PROPERTY_ADMINP.html */
func TestCertificateExpiration(t *testing.T) {
	_, err := administrationprocess.CertificateExpiration()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFICATEEXPIRATION_PROPERTY_ADMINP.html */
func TestSetCertificateExpiration(t *testing.T) {
	s, err := administrationprocess.CertificateExpiration()
	require.Nil(t, err)

	err = administrationprocess.SetCertificateExpiration(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERFILE_PROPERTY_ADMINP.html */
func TestCertifierFile(t *testing.T) {
	_, err := administrationprocess.CertifierFile()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERFILE_PROPERTY_ADMINP.html */
func TestSetCertifierFile(t *testing.T) {
	s, err := administrationprocess.CertifierFile()
	require.Nil(t, err)

	err = administrationprocess.SetCertifierFile(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERPASSWORD_PROPERTY_ADMINP.html */
func TestCertifierPassword(t *testing.T) {
	_, err := administrationprocess.CertifierPassword()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERPASSWORD_PROPERTY_ADMINP.html */
func TestSetCertifierPassword(t *testing.T) {
	s, err := administrationprocess.CertifierPassword()
	require.Nil(t, err)

	err = administrationprocess.SetCertifierPassword(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCERTIFICATEAUTHORITYAVAILABLE_PROPERTY_ADMINP.html */
func TestIsCertificateAuthorityAvailable(t *testing.T) {
	_, err := administrationprocess.IsCertificateAuthorityAvailable()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_ADMINP.html */
func TestUseCertificateAuthority(t *testing.T) {
	_, err := administrationprocess.UseCertificateAuthority()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_ADMINP.html */
func TestSetUseCertificateAuthority(t *testing.T) {
	s, err := administrationprocess.UseCertificateAuthority()
	require.Nil(t, err)

	err = administrationprocess.SetUseCertificateAuthority(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDGROUPMEMBERS_METHOD_ADMINP.html */
func TestAddGroupMembers(t *testing.T) {
	_, err := administrationprocess.AddGroupMembers( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDINTERNETCERTIFICATETOUSER_METHOD_ADMINP.html */
func TestAddInternetCertificateToUser(t *testing.T) {
	_, err := administrationprocess.AddInternetCertificateToUser( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDSERVERTOCLUSTER_METHOD_ADMINP.html */
func TestAddServerToCluster(t *testing.T) {
	_, err := administrationprocess.AddServerToCluster( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDELETEPERSONINDIRECTORY_METHOD_ADMINP.html */
func TestApproveDeletePersonInDirectory(t *testing.T) {
	_, err := administrationprocess.ApproveDeletePersonInDirectory( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDELETEPERSONINDIRECTORY_METHOD_ADMINP.html */
func TestApproveDeleteServerInDirectory(t *testing.T) {
	_, err := administrationprocess.ApproveDeleteServerInDirectory( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEDESIGNELEMENTDELETION_METHOD_ADMINP.html */
func TestApproveDesignElementDeletion(t *testing.T) {
	_, err := administrationprocess.ApproveDesignElementDeletion( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEHOSTEDORGSTORAGEDELETION_METHOD_ADMINP.html */
func TestApproveHostedOrgStorageDeletion(t *testing.T) {
	_, err := administrationprocess.ApproveHostedOrgStorageDeletion( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEMAILFILEDELETION_METHOD_ADMINP.html */
func TestApproveMailFileDeletion(t *testing.T) {
	_, err := administrationprocess.ApproveMailFileDeletion( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEMOVEDREPLICADELETION_METHOD_ADMINP.html */
func TestApproveMovedReplicaDeletion(t *testing.T) {
	_, err := administrationprocess.ApproveMovedReplicaDeletion( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVENAMECHANGERETRACTION_METHOD_ADMINP.html */
func TestApproveNameChangeRetraction(t *testing.T) {
	_, err := administrationprocess.ApproveNameChangeRetraction( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERENAMEPERSONINDIRECTORY_METHOD_ADMINP.html */
func TestApproveRenamePersonInDirectory(t *testing.T) {
	_, err := administrationprocess.ApproveRenamePersonInDirectory( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERENAMEPERSONINDIRECTORY_METHOD_ADMINP.html */
func TestApproveRenameServerInDirectory(t *testing.T) {
	_, err := administrationprocess.ApproveRenamePersonInDirectory( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVEREPLICADELETION_METHOD_ADMINP.html */
func TestApproveReplicaDeletion(t *testing.T) {
	_, err := administrationprocess.ApproveReplicaDeletion( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPROVERESOURCEDELETION_METHOD_ADMINP.html */
func TestApproveResourceDeletion(t *testing.T) {
	_, err := administrationprocess.ApproveResourceDeletion( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHANGEHTTPPASSWORD_METHOD_ADMINP.html */
func TestChangeHTTPPassword(t *testing.T) {
	_, err := administrationprocess.ChangeHTTPPassword( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONFIGUREMAILAGENT_METHOD_ADMINP.html */
func TestConfigureMailAgent(t *testing.T) {
	_, err := administrationprocess.ConfigureMailAgent( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREPLICA_METHOD_ADMINP.html */
func TestCreateReplica(t *testing.T) {
	_, err := administrationprocess.CreateReplica( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELEGATEMAILFILE_METHOD_ADMINP.html */
func TestDelegatemailfile(t *testing.T) {
	_, err := administrationprocess.DelegateMailFile( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEGROUP_METHOD_ADMINP.html */
func TestDeleteGroup(t *testing.T) {
	err := administrationprocess.DeleteGroup( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEREPLICAS_METHOD_ADMINP.html */
func TestDeleteReplicas(t *testing.T) {
	_, err := administrationprocess.DeleteReplicas( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETESERVER_METHOD_ADMINP.html */
func TestDeleteServer(t *testing.T) {
	_, err := administrationprocess.DeleteServer( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEUSER_METHOD_ADMINP.html */
func TestDeleteUser(t *testing.T) {
	_, err := administrationprocess.DeleteUser( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDGROUPINDOMAIN_METHOD_ADMINP.html */
func TestFindGroupInDomain(t *testing.T) {
	_, err := administrationprocess.FindGroupInDomain( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDSERVERINDOMAIN_METHOD_ADMINP.html */
func TestFindServerInDomain(t *testing.T) {
	_, err := administrationprocess.FindServerInDomain( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDUSERINDOMAIN_METHOD_ADMINP.html */
func TestFindUserInDomain(t *testing.T) {
	_, err := administrationprocess.FindUserInDomain( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEMAILUSER_METHOD_ADMINP.html */
func TestMoveMailUser(t *testing.T) {
	_, err := administrationprocess.MoveMailUser( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEREPLICA_METHOD_ADMINP.html */
func TestMoveReplica(t *testing.T) {
	_, err := administrationprocess.MoveReplica( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEROAMINGUSER_METHOD_ADMINP.html */
func TestMoveRoamingUser(t *testing.T) {
	_, err := administrationprocess.MoveRoamingUser( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEUSERINHIERARCHYCOMPLETE_METHOD_ADMINP.html */
func TestMoveUserInHierarchyComplete(t *testing.T) {
	_, err := administrationprocess.MoveUserInHierarchyComplete( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEUSERINHIERARCHYREQUEST_METHOD_ADMINP.html */
func TestMoveUserInHierarchyRequest(t *testing.T) {
	_, err := administrationprocess.MoveUserInHierarchyRequest( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFYSERVER_METHOD_ADMINP.html */
func TestRecertifyServer(t *testing.T) {
	_, err := administrationprocess.RecertifyServer( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFYUSER_METHOD_ADMINP.html */
func TestRecertifyUser(t *testing.T) {
	_, err := administrationprocess.RecertifyUser( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVESERVERFROMCLUSTER_METHOD_ADMINP.html */
func TestRemoveServerFromCluster(t *testing.T) {
	_, err := administrationprocess.RemoveServerFromCluster( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEGROUP_METHOD_ADMINP.html */
func TestRenameGroup(t *testing.T) {
	_, err := administrationprocess.RenameGroup( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMENOTESUSER_METHOD_ADMINP.html */
func TestRenameNotesUser(t *testing.T) {
	_, err := administrationprocess.RenameNotesUser( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEWEBUSER_METHOD_ADMINP.html */
func TestRenameWebUser(t *testing.T) {
	_, err := administrationprocess.RenameWebUser( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETENABLEOUTLOOKSUPPORT_METHOD_ADMINP.html */
func TestSetEnableOutlookSupport(t *testing.T) {
	_, err := administrationprocess.SetEnableOutlookSupport( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETSERVERDIRECTORYASSISTANCESETTINGS_METHOD_ADMINP.html */
func TestSetServerDirectoryAssistanceSettings(t *testing.T) {
	_, err := administrationprocess.SetServerDirectoryAssistanceSettings( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETUSERPASSWORDSETTINGS_METHOD_ADMINP.html */
func TestSetUserPasswordSettings(t *testing.T) {
	_, err := administrationprocess.SetUserPasswordSettings( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIGNDATABASEWITHSERVERID_METHOD_ADMINP.html */
func TestSignDatabaseWithServerID(t *testing.T) {
	_, err := administrationprocess.SignDatabaseWithServerID( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPGRADEUSERTOHIERARCHICAL_METHOD_ADMINP.html */
func TestUpgradeUserToHierarchical(t *testing.T) {
	_, err := administrationprocess.UpgradeUserToHierarchical( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}
