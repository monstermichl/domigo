/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREGISTRATION_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesRegistration struct {
	NotesStruct
}

func newNotesRegistration(dispatchPtr *ole.IDispatch) NotesRegistration {
	return NotesRegistration{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNIT_PROPERTY_NOTESREGISTRATION.html */
func (r NotesRegistration) AltOrgUnit() ([]String, error) {
	return getComArrayProperty[String](r, "AltOrgUnit")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNIT_PROPERTY_NOTESREGISTRATION.html */
func (r NotesRegistration) SetAltOrgUnit(v []String) error {
	return putComProperty(r, "AltOrgUnit", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNITLANG_PROPERTY_NOTESREGISTRATION.html */
func (r NotesRegistration) AltOrgUnitLang() ([]String, error) {
	return getComArrayProperty[String](r, "AltOrgUnitLang")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNITLANG_PROPERTY_NOTESREGISTRATION.html */
func (r NotesRegistration) SetAltOrgUnitLang(v []String) error {
	return putComProperty(r, "AltOrgUnitLang", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERIDFILE_PROPERTY.html */
func (r NotesRegistration) CertifierIDFile() (String, error) {
	return getComProperty[String](r, "CertifierIDFile")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERIDFILE_PROPERTY.html */
func (r NotesRegistration) SetCertifierIDFile(v String) error {
	return putComProperty(r, "CertifierIDFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERNAME_PROPERTY_REG.html */
func (r NotesRegistration) CertifierName() (String, error) {
	return getComProperty[String](r, "CertifierName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERNAME_PROPERTY_REG.html */
func (r NotesRegistration) SetCertifierName(v String) error {
	return putComProperty(r, "CertifierName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMAILDB_PROPERTY.html */
func (r NotesRegistration) CreateMailDb() (Boolean, error) {
	return getComProperty[Boolean](r, "CreateMailDb")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMAILDB_PROPERTY.html */
func (r NotesRegistration) SetCreateMailDb(v Boolean) error {
	return putComProperty(r, "CreateMailDb", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENFORCEUNIQUESHORTNAME_PROPERTY_REG.html */
func (r NotesRegistration) EnforceUniqueShortName() (Boolean, error) {
	return getComProperty[Boolean](r, "EnforceUniqueShortName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENFORCEUNIQUESHORTNAME_PROPERTY_REG.html */
func (r NotesRegistration) SetEnforceUniqueShortName(v Boolean) error {
	return putComProperty(r, "EnforceUniqueShortName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPIRATION_PROPERTY.html */
func (r NotesRegistration) Expiration() (Time, error) {
	return getComProperty[Time](r, "Expiration")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPIRATION_PROPERTY.html */
func (r NotesRegistration) SetExpiration(v Time) error {
	return putComProperty(r, "Expiration", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPLIST_PROPERTY_REG.html */
func (r NotesRegistration) GroupList() ([]String, error) {
	return getComArrayProperty[String](r, "GroupList")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPLIST_PROPERTY_REG.html */
func (r NotesRegistration) SetGroupList(v []String) error {
	return putComProperty(r, "GroupList", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IDTYPE_PROPERTY.html */
func (r NotesRegistration) IDType() (Long, error) {
	return getComProperty[Long](r, "IDType")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IDTYPE_PROPERTY.html */
func (r NotesRegistration) SetIDType(v Long) error {
	return putComProperty(r, "IDType", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNORTHAMERICAN_PROPERTY.html */
func (r NotesRegistration) IsNorthAmerican() (Boolean, error) {
	return getComProperty[Boolean](r, "IsNorthAmerican")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNORTHAMERICAN_PROPERTY.html */
func (r NotesRegistration) SetIsNorthAmerican(v Boolean) error {
	return putComProperty(r, "IsNorthAmerican", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROAMINGUSER_PROPERTY_REG.html */
func (r NotesRegistration) IsRoamingUser() (Boolean, error) {
	return getComProperty[Boolean](r, "IsRoamingUser")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROAMINGUSER_PROPERTY_REG.html */
func (r NotesRegistration) SetIsRoamingUser(v Boolean) error {
	return putComProperty(r, "IsRoamingUser", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILACLMANAGER_PROPERTY_REG.html */
func (r NotesRegistration) MailACLManager() (String, error) {
	return getComProperty[String](r, "MailACLManager")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILACLMANAGER_PROPERTY_REG.html */
func (r NotesRegistration) SetMailACLManager(v String) error {
	return putComProperty(r, "MailACLManager", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILCREATEFTINDEX_PROPERTY_REG.html */
func (r NotesRegistration) MailCreateFTIndex() (Boolean, error) {
	return getComProperty[Boolean](r, "MailCreateFTIndex")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILCREATEFTINDEX_PROPERTY_REG.html */
func (r NotesRegistration) SetMailCreateFTIndex(v Boolean) error {
	return putComProperty(r, "MailCreateFTIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILINTERNETADDRESS_PROPERTY_REG.html */
func (r NotesRegistration) MailInternetAddress() (String, error) {
	return getComProperty[String](r, "MailInternetAddress")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILINTERNETADDRESS_PROPERTY_REG.html */
func (r NotesRegistration) SetMailInternetAddress(v String) error {
	return putComProperty(r, "MailInternetAddress", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILOWNERACCESS_PROPERTY_REG.html */
func (r NotesRegistration) MailOwnerAccess() (Long, error) {
	return getComProperty[Long](r, "MailOwnerAccess")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILOWNERACCESS_PROPERTY_REG.html */
func (r NotesRegistration) SetMailOwnerAccess(v Long) error {
	return putComProperty(r, "MailOwnerAccess", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTASIZELIMIT_PROPERTY_REG.html */
func (r NotesRegistration) MailQuotaSizeLimit() (Long, error) {
	return getComProperty[Long](r, "MailQuotaSizeLimit")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTASIZELIMIT_PROPERTY_REG.html */
func (r NotesRegistration) SetMailQuotaSizeLimit(v Long) error {
	return putComProperty(r, "MailQuotaSizeLimit", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTAWARNINGTHRESHOLD_PROPERTY_REG.html */
func (r NotesRegistration) MailQuotaWarningThreshold() (Long, error) {
	return getComProperty[Long](r, "MailQuotaWarningThreshold")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTAWARNINGTHRESHOLD_PROPERTY_REG.html */
func (r NotesRegistration) SetMailQuotaWarningThreshold(v Long) error {
	return putComProperty(r, "MailQuotaWarningThreshold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILREPLICASERVERS_PROPERTY_REG.html */
func (r NotesRegistration) MailReplicaServers() ([]String, error) {
	return getComArrayProperty[String](r, "MailReplicaServers")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILREPLICASERVERS_PROPERTY_REG.html */
func (r NotesRegistration) SetMailReplicaServers(v []String) error {
	return putComProperty(r, "MailReplicaServers", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILSYSTEM_PROPERTY_REG.html */
func (r NotesRegistration) MailSystem() (Long, error) {
	return getComProperty[Long](r, "MailSystem")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILSYSTEM_PROPERTY_REG.html */
func (r NotesRegistration) SetMailSystem(v Long) error {
	return putComProperty(r, "MailSystem", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILTEMPLATENAME_PROPERTY_REG.html */
func (r NotesRegistration) MailTemplateName() (String, error) {
	return getComProperty[String](r, "MailTemplateName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILTEMPLATENAME_PROPERTY_REG.html */
func (r NotesRegistration) SetMailTemplateName(v String) error {
	return putComProperty(r, "MailTemplateName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MINPASSWORDLENGTH_PROPERTY.html */
func (r NotesRegistration) MinPasswordLength() (Long, error) {
	return getComProperty[Long](r, "MinPasswordLength")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MINPASSWORDLENGTH_PROPERTY.html */
func (r NotesRegistration) SetMinPasswordLength(v Long) error {
	return putComProperty(r, "MinPasswordLength", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOIDFILE_PROPERTY_REG.html */
func (r NotesRegistration) NoIDFile() (Boolean, error) {
	return getComProperty[Boolean](r, "NoIDFile")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOIDFILE_PROPERTY_REG.html */
func (r NotesRegistration) SetNoIDFile(v Boolean) error {
	return putComProperty(r, "NoIDFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT_PROPERTY.html */
func (r NotesRegistration) OrgUnit() (String, error) {
	return getComProperty[String](r, "OrgUnit")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT_PROPERTY.html */
func (r NotesRegistration) SetOrgUnit(v String) error {
	return putComProperty(r, "OrgUnit", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POLICYNAME_PROPERTY_REG.html */
func (r NotesRegistration) PolicyName() (String, error) {
	return getComProperty[String](r, "PolicyName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POLICYNAME_PROPERTY_REG.html */
func (r NotesRegistration) SetPolicyName(v String) error {
	return putComProperty(r, "PolicyName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONLOG_PROPERTY.html */
func (r NotesRegistration) RegistrationLog() (String, error) {
	return getComProperty[String](r, "RegistrationLog")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONLOG_PROPERTY.html */
func (r NotesRegistration) SetRegistrationLog(v String) error {
	return putComProperty(r, "RegistrationLog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONSERVER_PROPERTY.html */
func (r NotesRegistration) RegistrationServer() (String, error) {
	return getComProperty[String](r, "RegistrationServer")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONSERVER_PROPERTY.html */
func (r NotesRegistration) SetRegistrationServer(v String) error {
	return putComProperty(r, "RegistrationServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPPERIOD_PROPERTY_REG.html */
func (r NotesRegistration) RoamingCleanupPeriod() (Long, error) {
	return getComProperty[Long](r, "RoamingCleanupPeriod")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPPERIOD_PROPERTY_REG.html */
func (r NotesRegistration) SetRoamingCleanupPeriod(v Long) error {
	return putComProperty(r, "RoamingCleanupPeriod", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPSETTING_PROPERTY_REG.html */
func (r NotesRegistration) RoamingCleanupSetting() (Long, error) {
	return getComProperty[Long](r, "RoamingCleanupSetting")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPSETTING_PROPERTY_REG.html */
func (r NotesRegistration) SetRoamingCleanupSetting(v Long) error {
	return putComProperty(r, "RoamingCleanupSetting", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSERVER_PROPERTY_REG.html */
func (r NotesRegistration) RoamingServer() (String, error) {
	return getComProperty[String](r, "RoamingServer")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSERVER_PROPERTY_REG.html */
func (r NotesRegistration) SetRoamingServer(v String) error {
	return putComProperty(r, "RoamingServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSUBDIR_PROPERTY_REG.html */
func (r NotesRegistration) RoamingSubdir() (String, error) {
	return getComProperty[String](r, "RoamingSubdir")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSUBDIR_PROPERTY_REG.html */
func (r NotesRegistration) SetRoamingSubdir(v String) error {
	return putComProperty(r, "RoamingSubdir", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SHORTNAME_PROPERTY_REG.html */
func (r NotesRegistration) ShortName() (String, error) {
	return getComProperty[String](r, "ShortName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SHORTNAME_PROPERTY_REG.html */
func (r NotesRegistration) SetShortName(v String) error {
	return putComProperty(r, "ShortName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDTOADDRESSBOOK_PROPERTY.html */
func (r NotesRegistration) StoreIDInAddressBook() (Boolean, error) {
	return getComProperty[Boolean](r, "StoreIDInAddressBook")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDTOADDRESSBOOK_PROPERTY.html */
func (r NotesRegistration) SetStoreIDInAddressBook(v Boolean) error {
	return putComProperty(r, "StoreIDInAddressBook", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDINMAILFILE_PROPERTY_REG.html */
func (r NotesRegistration) StoreIDInMailfile() (Boolean, error) {
	return getComProperty[Boolean](r, "StoreIDInMailfile")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDINMAILFILE_PROPERTY_REG.html */
func (r NotesRegistration) SetStoreIDInMailfile(v Boolean) error {
	return putComProperty(r, "StoreIDInMailfile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SYNCHINTERNETPASSWORD_PROPERTY_REG.html */
func (r NotesRegistration) SynchInternetPassword() (Boolean, error) {
	return getComProperty[Boolean](r, "SynchInternetPassword")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SYNCHINTERNETPASSWORD_PROPERTY_REG.html */
func (r NotesRegistration) SetSynchInternetPassword(v Boolean) error {
	return putComProperty(r, "SynchInternetPassword", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEADDRESSBOOK_PROPERTY.html */
func (r NotesRegistration) UpdateAddressBook() (Boolean, error) {
	return getComProperty[Boolean](r, "UpdateAddressBook")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEADDRESSBOOK_PROPERTY.html */
func (r NotesRegistration) SetUpdateAddressBook(v Boolean) error {
	return putComProperty(r, "UpdateAddressBook", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_REG.html */
func (r NotesRegistration) UseCertificateAuthority() (Boolean, error) {
	return getComProperty[Boolean](r, "UseCertificateAuthority")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_REG.html */
func (r NotesRegistration) SetUseCertificateAuthority(v Boolean) error {
	return putComProperty(r, "UseCertificateAuthority", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCERTIFIERTOADDRESSBOOK_METHOD.html */
func (r NotesRegistration) AddCertifierToAddressBook(idfile String, certpw String, location String, comment String) error {
	return callComVoidMethod(r, "AddCertifierToAddressBook", idfile, certpw, location, comment)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDSERVERTOADDRESSBOOK_METHOD.html */
func (r NotesRegistration) AddServerToAddressBook(idfile String, server String, domain String, userpw String, network String, adminname String, title String, location String, comment String) error {
	return callComVoidMethod(r, "AddServerToAddressBook", idfile, server, domain, userpw, network, adminname, title, location, comment)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDUSERPROFILE_METHOD.html */
func (r NotesRegistration) AddUserProfile(username String, profilename String) error {
	return callComVoidMethod(r, "AddUserProfile", username, profilename)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDUSERTOADDRESSBOOK_METHOD.html */
func (r NotesRegistration) AddUserToAddressBook(idfile String, fullname String, lastname String, userpw String, firstname String, middle String, mailserver String, mailidpath String, fwdaddress String, location String, comment String) error {
	return callComVoidMethod(r, "AddUserToAddressBook", idfile, fullname, lastname, userpw, firstname, middle, mailserver, mailidpath, fwdaddress, location, comment)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CROSSCERTIFY_METHOD.html */
func (r NotesRegistration) CrossCertify(idfile String, certpw String, comment String) error {
	return callComVoidMethod(r, "CrossCertify", idfile, certpw, comment)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEIDONSERVER_METHOD.html */
func (r NotesRegistration) DeleteIDOnServer(username String, isserverid Boolean) error {
	return callComVoidMethod(r, "DeleteIDOnServer", username, isserverid)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETIDFROMSERVER_METHOD.html */
func (r NotesRegistration) GetIDFromServer(username String, filepath String, isserverid Boolean) error {
	return callComVoidMethod(r, "GetIDFromServer", username, filepath, isserverid)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERINFO_METHOD.html */
func (r NotesRegistration) GetUserInfo(username String, retMailServer String, retMailFile String, retMailDomain String, retMailSystem Integer, retProfile String) error {
	return callComVoidMethod(r, "GetUserInfo", username, retMailServer, retMailFile, retMailDomain, retMailSystem, retProfile)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFY_METHOD.html */
func (r NotesRegistration) Recertify(idfile String, certpw String, comment String) error {
	return callComVoidMethod(r, "Recertify", idfile, certpw, comment)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWCERTIFIER_METHOD.html */
func (r NotesRegistration) RegisterNewCertifier(organization String, idfile String, certpw String, country String) error {
	return callComVoidMethod(r, "RegisterNewCertifier", organization, idfile, certpw, country)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWSERVER_METHOD.html */
func (r NotesRegistration) RegisterNewServer(server String, idfile String, domain String, servpw String, certpw String, location String, comment String, network String, adminname String, title String) error {
	return callComVoidMethod(r, "RegisterNewServer", server, idfile, domain, servpw, certpw, location, comment, network, adminname, title)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWUSER_METHOD.html */
type notesRegistrationRegisterNewUserParams struct {
	firstName   *String
	middle      *String
	certPw      *String
	location    *String
	comment     *String
	mailDbPath  *String
	fwdDomain   *String
	userPw      *String
	altName     *String
	altNameLang *String
	usertype    *Integer
}

type notesRegistrationRegisterNewUserParam func(*notesRegistrationRegisterNewUserParams)

func WithNotesRegistrationRegisterNewUserFirstName(firstName String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.firstName = &firstName
	}
}

func WithNotesRegistrationRegisterNewUserMiddle(middle String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.middle = &middle
	}
}

func WithNotesRegistrationRegisterNewUserCertPw(certPw String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.certPw = &certPw
	}
}

func WithNotesRegistrationRegisterNewUserLocation(location String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.location = &location
	}
}

func WithNotesRegistrationRegisterNewUserComment(comment String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.comment = &comment
	}
}

func WithNotesRegistrationRegisterNewUserMailDbPath(mailDbPath String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.mailDbPath = &mailDbPath
	}
}

func WithNotesRegistrationRegisterNewUserFwdDomain(fwdDomain String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.fwdDomain = &fwdDomain
	}
}

func WithNotesRegistrationRegisterNewUserUserPw(userPw String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.userPw = &userPw
	}
}

func WithNotesRegistrationRegisterNewUserAltName(altName String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.altName = &altName
	}
}

func WithNotesRegistrationRegisterNewUserAltNameLang(altNameLang String) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.altNameLang = &altNameLang
	}
}

func WithNotesRegistrationRegisterNewUserUserType(usertype Integer) notesRegistrationRegisterNewUserParam {
	return func(c *notesRegistrationRegisterNewUserParams) {
		c.usertype = &usertype
	}
}

func (r NotesRegistration) RegisterNewUser(lastName String, idFile String, mailServer String, params ...notesRegistrationRegisterNewUserParam) error {
	paramsStruct := &notesRegistrationRegisterNewUserParams{}
	paramsOrdered := []interface{}{lastName, idFile, mailServer}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.usertype != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.usertype)
		if paramsStruct.middle != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.middle)
			if paramsStruct.certPw != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.certPw)
				if paramsStruct.location != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.location)
					if paramsStruct.comment != nil {
						paramsOrdered = append(paramsOrdered, *paramsStruct.comment)
						if paramsStruct.mailDbPath != nil {
							paramsOrdered = append(paramsOrdered, *paramsStruct.mailDbPath)
							if paramsStruct.fwdDomain != nil {
								paramsOrdered = append(paramsOrdered, *paramsStruct.fwdDomain)
								if paramsStruct.userPw != nil {
									paramsOrdered = append(paramsOrdered, *paramsStruct.userPw)
									if paramsStruct.altName != nil {
										paramsOrdered = append(paramsOrdered, *paramsStruct.altName)
										if paramsStruct.altNameLang != nil {
											paramsOrdered = append(paramsOrdered, *paramsStruct.altNameLang)
											if paramsStruct.usertype != nil {
												paramsOrdered = append(paramsOrdered, *paramsStruct.usertype)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return callComVoidMethod(r, "RegisterNewUser", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SWITCHTOID_METHOD.html */
func (r NotesRegistration) SwitchToID(idfile String, userpw String) error {
	return callComVoidMethod(r, "SwitchToID", idfile, userpw)
}
