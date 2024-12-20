/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREGISTRATION_CLASS.html */
package notesregistration_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"
	"github.com/stretchr/testify/require"
)

var registration domigo.NotesRegistration

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error
		registration, err = session.CreateRegistration()
		defer registration.Release()

		if err != nil {
			return "Registration could not be created", err
		}
		/* Don't run tests for now to not mess with some Notes stuff as NotesRegistration
		   is not bound to a specific database. */
		// m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNIT_PROPERTY_NOTESREGISTRATION.html */
func TestAltOrgUnit(t *testing.T) {
	_, err := registration.AltOrgUnit()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNIT_PROPERTY_NOTESREGISTRATION.html */
func TestSetAltOrgUnit(t *testing.T) {
	s, err := registration.AltOrgUnit()
	require.Nil(t, err)

	err = registration.SetAltOrgUnit(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNITLANG_PROPERTY_NOTESREGISTRATION.html */
func TestAltOrgUnitLang(t *testing.T) {
	_, err := registration.AltOrgUnitLang()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNITLANG_PROPERTY_NOTESREGISTRATION.html */
func TestSetAltOrgUnitLang(t *testing.T) {
	s, err := registration.AltOrgUnitLang()
	require.Nil(t, err)

	err = registration.SetAltOrgUnitLang(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERIDFILE_PROPERTY.html */
func TestCertifierIDFile(t *testing.T) {
	_, err := registration.CertifierIDFile()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERIDFILE_PROPERTY.html */
func TestSetCertifierIDFile(t *testing.T) {
	s, err := registration.CertifierIDFile()
	require.Nil(t, err)

	err = registration.SetCertifierIDFile(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERNAME_PROPERTY_REG.html */
func TestCertifierName(t *testing.T) {
	_, err := registration.CertifierName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERNAME_PROPERTY_REG.html */
func TestSetCertifierName(t *testing.T) {
	s, err := registration.CertifierName()
	require.Nil(t, err)

	err = registration.SetCertifierName(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMAILDB_PROPERTY.html */
func TestCreateMailDb(t *testing.T) {
	_, err := registration.CreateMailDb()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMAILDB_PROPERTY.html */
func TestSetCreateMailDb(t *testing.T) {
	s, err := registration.CreateMailDb()
	require.Nil(t, err)

	err = registration.SetCreateMailDb(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENFORCEUNIQUESHORTNAME_PROPERTY_REG.html */
func TestEnforceUniqueShortName(t *testing.T) {
	_, err := registration.EnforceUniqueShortName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENFORCEUNIQUESHORTNAME_PROPERTY_REG.html */
func TestSetEnforceUniqueShortName(t *testing.T) {
	s, err := registration.EnforceUniqueShortName()
	require.Nil(t, err)

	err = registration.SetEnforceUniqueShortName(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPIRATION_PROPERTY.html */
func TestExpiration(t *testing.T) {
	_, err := registration.Expiration()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPIRATION_PROPERTY.html */
func TestSetExpiration(t *testing.T) {
	s, err := registration.Expiration()
	require.Nil(t, err)

	err = registration.SetExpiration(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPLIST_PROPERTY_REG.html */
func TestGroupList(t *testing.T) {
	_, err := registration.GroupList()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPLIST_PROPERTY_REG.html */
func TestSetGroupList(t *testing.T) {
	s, err := registration.GroupList()
	require.Nil(t, err)

	err = registration.SetGroupList(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IDTYPE_PROPERTY.html */
func TestIDType(t *testing.T) {
	_, err := registration.IDType()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IDTYPE_PROPERTY.html */
func TestSetIDType(t *testing.T) {
	s, err := registration.IDType()
	require.Nil(t, err)

	err = registration.SetIDType(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNORTHAMERICAN_PROPERTY.html */
func TestIsNorthAmerican(t *testing.T) {
	_, err := registration.IsNorthAmerican()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNORTHAMERICAN_PROPERTY.html */
func TestSetIsNorthAmerican(t *testing.T) {
	s, err := registration.IsNorthAmerican()
	require.Nil(t, err)

	err = registration.SetIsNorthAmerican(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROAMINGUSER_PROPERTY_REG.html */
func TestIsRoamingUser(t *testing.T) {
	_, err := registration.IsRoamingUser()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROAMINGUSER_PROPERTY_REG.html */
func TestSetIsRoamingUser(t *testing.T) {
	s, err := registration.IsRoamingUser()
	require.Nil(t, err)

	err = registration.SetIsRoamingUser(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILACLMANAGER_PROPERTY_REG.html */
func TestMailACLManager(t *testing.T) {
	_, err := registration.MailACLManager()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILACLMANAGER_PROPERTY_REG.html */
func TestSetMailACLManager(t *testing.T) {
	s, err := registration.MailACLManager()
	require.Nil(t, err)

	err = registration.SetMailACLManager(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILCREATEFTINDEX_PROPERTY_REG.html */
func TestMailCreateFTIndex(t *testing.T) {
	_, err := registration.MailCreateFTIndex()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILCREATEFTINDEX_PROPERTY_REG.html */
func TestSetMailCreateFTIndex(t *testing.T) {
	s, err := registration.MailCreateFTIndex()
	require.Nil(t, err)

	err = registration.SetMailCreateFTIndex(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILINTERNETADDRESS_PROPERTY_REG.html */
func TestMailInternetAddress(t *testing.T) {
	_, err := registration.MailInternetAddress()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILINTERNETADDRESS_PROPERTY_REG.html */
func TestSetMailInternetAddress(t *testing.T) {
	s, err := registration.MailInternetAddress()
	require.Nil(t, err)

	err = registration.SetMailInternetAddress(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILOWNERACCESS_PROPERTY_REG.html */
func TestMailOwnerAccess(t *testing.T) {
	_, err := registration.MailOwnerAccess()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILOWNERACCESS_PROPERTY_REG.html */
func TestSetMailOwnerAccess(t *testing.T) {
	s, err := registration.MailOwnerAccess()
	require.Nil(t, err)

	err = registration.SetMailOwnerAccess(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTASIZELIMIT_PROPERTY_REG.html */
func TestMailQuotaSizeLimit(t *testing.T) {
	_, err := registration.MailQuotaSizeLimit()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTASIZELIMIT_PROPERTY_REG.html */
func TestSetMailQuotaSizeLimit(t *testing.T) {
	s, err := registration.MailQuotaSizeLimit()
	require.Nil(t, err)

	err = registration.SetMailQuotaSizeLimit(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTAWARNINGTHRESHOLD_PROPERTY_REG.html */
func TestMailQuotaWarningThreshold(t *testing.T) {
	_, err := registration.MailQuotaWarningThreshold()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTAWARNINGTHRESHOLD_PROPERTY_REG.html */
func TestSetMailQuotaWarningThreshold(t *testing.T) {
	s, err := registration.MailQuotaWarningThreshold()
	require.Nil(t, err)

	err = registration.SetMailQuotaWarningThreshold(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILREPLICASERVERS_PROPERTY_REG.html */
func TestMailReplicaServers(t *testing.T) {
	_, err := registration.MailReplicaServers()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILREPLICASERVERS_PROPERTY_REG.html */
func TestSetMailReplicaServers(t *testing.T) {
	s, err := registration.MailReplicaServers()
	require.Nil(t, err)

	err = registration.SetMailReplicaServers(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILSYSTEM_PROPERTY_REG.html */
func TestMailSystem(t *testing.T) {
	_, err := registration.MailSystem()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILSYSTEM_PROPERTY_REG.html */
func TestSetMailSystem(t *testing.T) {
	s, err := registration.MailSystem()
	require.Nil(t, err)

	err = registration.SetMailSystem(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILTEMPLATENAME_PROPERTY_REG.html */
func TestMailTemplateName(t *testing.T) {
	_, err := registration.MailTemplateName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILTEMPLATENAME_PROPERTY_REG.html */
func TestSetMailTemplateName(t *testing.T) {
	s, err := registration.MailTemplateName()
	require.Nil(t, err)

	err = registration.SetMailTemplateName(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MINPASSWORDLENGTH_PROPERTY.html */
func TestMinPasswordLength(t *testing.T) {
	_, err := registration.MinPasswordLength()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MINPASSWORDLENGTH_PROPERTY.html */
func TestSetMinPasswordLength(t *testing.T) {
	s, err := registration.MinPasswordLength()
	require.Nil(t, err)

	err = registration.SetMinPasswordLength(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOIDFILE_PROPERTY_REG.html */
func TestNoIDFile(t *testing.T) {
	_, err := registration.NoIDFile()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOIDFILE_PROPERTY_REG.html */
func TestSetNoIDFile(t *testing.T) {
	s, err := registration.NoIDFile()
	require.Nil(t, err)

	err = registration.SetNoIDFile(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT_PROPERTY.html */
func TestOrgUnit(t *testing.T) {
	_, err := registration.OrgUnit()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT_PROPERTY.html */
func TestSetOrgUnit(t *testing.T) {
	s, err := registration.OrgUnit()
	require.Nil(t, err)

	err = registration.SetOrgUnit(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POLICYNAME_PROPERTY_REG.html */
func TestPolicyName(t *testing.T) {
	_, err := registration.PolicyName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POLICYNAME_PROPERTY_REG.html */
func TestSetPolicyName(t *testing.T) {
	s, err := registration.PolicyName()
	require.Nil(t, err)

	err = registration.SetPolicyName(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONLOG_PROPERTY.html */
func TestRegistrationLog(t *testing.T) {
	_, err := registration.RegistrationLog()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONLOG_PROPERTY.html */
func TestSetRegistrationLog(t *testing.T) {
	s, err := registration.RegistrationLog()
	require.Nil(t, err)

	err = registration.SetRegistrationLog(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONSERVER_PROPERTY.html */
func TestRegistrationServer(t *testing.T) {
	_, err := registration.RegistrationServer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONSERVER_PROPERTY.html */
func TestSetRegistrationServer(t *testing.T) {
	s, err := registration.RegistrationServer()
	require.Nil(t, err)

	err = registration.SetRegistrationServer(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPPERIOD_PROPERTY_REG.html */
func TestRoamingCleanupPeriod(t *testing.T) {
	_, err := registration.RoamingCleanupPeriod()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPPERIOD_PROPERTY_REG.html */
func TestSetRoamingCleanupPeriod(t *testing.T) {
	s, err := registration.RoamingCleanupPeriod()
	require.Nil(t, err)

	err = registration.SetRoamingCleanupPeriod(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPSETTING_PROPERTY_REG.html */
func TestRoamingCleanupSetting(t *testing.T) {
	_, err := registration.RoamingCleanupSetting()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPSETTING_PROPERTY_REG.html */
func TestSetRoamingCleanupSetting(t *testing.T) {
	s, err := registration.RoamingCleanupSetting()
	require.Nil(t, err)

	err = registration.SetRoamingCleanupSetting(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSERVER_PROPERTY_REG.html */
func TestRoamingServer(t *testing.T) {
	_, err := registration.RoamingServer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSERVER_PROPERTY_REG.html */
func TestSetRoamingServer(t *testing.T) {
	s, err := registration.RoamingServer()
	require.Nil(t, err)

	err = registration.SetRoamingServer(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSUBDIR_PROPERTY_REG.html */
func TestRoamingSubdir(t *testing.T) {
	_, err := registration.RoamingSubdir()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSUBDIR_PROPERTY_REG.html */
func TestSetRoamingSubdir(t *testing.T) {
	s, err := registration.RoamingSubdir()
	require.Nil(t, err)

	err = registration.SetRoamingSubdir(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SHORTNAME_PROPERTY_REG.html */
func TestShortName(t *testing.T) {
	_, err := registration.ShortName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SHORTNAME_PROPERTY_REG.html */
func TestSetShortName(t *testing.T) {
	s, err := registration.ShortName()
	require.Nil(t, err)

	err = registration.SetShortName(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDTOADDRESSBOOK_PROPERTY.html */
func TestStoreIDInAddressBook(t *testing.T) {
	_, err := registration.StoreIDInAddressBook()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDTOADDRESSBOOK_PROPERTY.html */
func TestSetStoreIDInAddressBook(t *testing.T) {
	s, err := registration.StoreIDInAddressBook()
	require.Nil(t, err)

	err = registration.SetStoreIDInAddressBook(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDINMAILFILE_PROPERTY_REG.html */
func TestStoreIDInMailfile(t *testing.T) {
	_, err := registration.StoreIDInMailfile()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDINMAILFILE_PROPERTY_REG.html */
func TestSetStoreIDInMailfile(t *testing.T) {
	s, err := registration.StoreIDInMailfile()
	require.Nil(t, err)

	err = registration.SetStoreIDInMailfile(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SYNCHINTERNETPASSWORD_PROPERTY_REG.html */
func TestSynchInternetPassword(t *testing.T) {
	_, err := registration.SynchInternetPassword()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SYNCHINTERNETPASSWORD_PROPERTY_REG.html */
func TestSetSynchInternetPassword(t *testing.T) {
	s, err := registration.SynchInternetPassword()
	require.Nil(t, err)

	err = registration.SetSynchInternetPassword(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEADDRESSBOOK_PROPERTY.html */
func TestUpdateAddressBook(t *testing.T) {
	_, err := registration.UpdateAddressBook()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEADDRESSBOOK_PROPERTY.html */
func TestSetUpdateAddressBook(t *testing.T) {
	s, err := registration.UpdateAddressBook()
	require.Nil(t, err)

	err = registration.SetUpdateAddressBook(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_REG.html */
func TestUseCertificateAuthority(t *testing.T) {
	_, err := registration.UseCertificateAuthority()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_REG.html */
func TestSetUseCertificateAuthority(t *testing.T) {
	s, err := registration.UseCertificateAuthority()
	require.Nil(t, err)

	err = registration.SetUseCertificateAuthority(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCERTIFIERTOADDRESSBOOK_METHOD.html */
// func TestAddCertifierToAddressBook(t *testing.T) {
// 	err := registration.AddCertifierToAddressBook( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDSERVERTOADDRESSBOOK_METHOD.html */
// func TestAddServerToAddressBook(t *testing.T) {
// 	err := registration.AddServerToAddressBook( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDUSERPROFILE_METHOD.html */
// func TestAddUserProfile(t *testing.T) {
// 	err := registration.AddUserProfile( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDUSERTOADDRESSBOOK_METHOD.html */
// func TestAddUserToAddressBook(t *testing.T) {
// 	err := registration.AddUserToAddressBook( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CROSSCERTIFY_METHOD.html */
// func TestCrossCertify(t *testing.T) {
// 	err := registration.CrossCertify( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEIDONSERVER_METHOD.html */
// func TestDeleteIDOnServer(t *testing.T) {
// 	err := registration.DeleteIDOnServer( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETIDFROMSERVER_METHOD.html */
// func TestGetIDFromServer(t *testing.T) {
// 	err := registration.GetIDFromServer( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERINFO_METHOD.html */
// func TestGetUserInfo(t *testing.T) {
// 	err := registration.GetUserInfo( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFY_METHOD.html */
// func TestRecertify(t *testing.T) {
// 	err := registration.Recertify( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWCERTIFIER_METHOD.html */
// func TestRegisterNewCertifier(t *testing.T) {
// 	err := registration.RegisterNewCertifier( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWSERVER_METHOD.html */
// func TestRegisterNewServer(t *testing.T) {
// 	err := registration.RegisterNewServer( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWUSER_METHOD.html */
// func TestRegisterNewUser(t *testing.T) {
// 	err := registration.RegisterNewUser( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SWITCHTOID_METHOD.html */
// func TestSwitchToID(t *testing.T) {
// 	err := registration.SwitchToID( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }
