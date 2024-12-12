/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREGISTRATION_CLASS.html */
package notesregistration

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesRegistration struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesRegistration {
	return NotesRegistration{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNIT_PROPERTY_NOTESREGISTRATION.html */
func (r NotesRegistration) AltOrgUnit() ([]domino.String, error) {
	vals, err := r.Com().GetArrayProperty("AltOrgUnit")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNIT_PROPERTY_NOTESREGISTRATION.html */
func (r NotesRegistration) SetAltOrgUnit(v []domino.String) error {
	return r.Com().PutProperty("AltOrgUnit", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNITLANG_PROPERTY_NOTESREGISTRATION.html */
func (r NotesRegistration) AltOrgUnitLang() ([]domino.String, error) {
	vals, err := r.Com().GetArrayProperty("AltOrgUnitLang")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALTORGUNITLANG_PROPERTY_NOTESREGISTRATION.html */
func (r NotesRegistration) SetAltOrgUnitLang(v []domino.String) error {
	return r.Com().PutProperty("AltOrgUnitLang", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERIDFILE_PROPERTY.html */
func (r NotesRegistration) CertifierIDFile() (domino.String, error) {
	val, err := r.Com().GetProperty("CertifierIDFile")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERIDFILE_PROPERTY.html */
func (r NotesRegistration) SetCertifierIDFile(v domino.String) error {
	return r.Com().PutProperty("CertifierIDFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERNAME_PROPERTY_REG.html */
func (r NotesRegistration) CertifierName() (domino.String, error) {
	val, err := r.Com().GetProperty("CertifierName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CERTIFIERNAME_PROPERTY_REG.html */
func (r NotesRegistration) SetCertifierName(v domino.String) error {
	return r.Com().PutProperty("CertifierName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMAILDB_PROPERTY.html */
func (r NotesRegistration) CreateMailDb() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("CreateMailDb")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEMAILDB_PROPERTY.html */
func (r NotesRegistration) SetCreateMailDb(v domino.Boolean) error {
	return r.Com().PutProperty("CreateMailDb", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENFORCEUNIQUESHORTNAME_PROPERTY_REG.html */
func (r NotesRegistration) EnforceUniqueShortName() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("EnforceUniqueShortName")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENFORCEUNIQUESHORTNAME_PROPERTY_REG.html */
func (r NotesRegistration) SetEnforceUniqueShortName(v domino.Boolean) error {
	return r.Com().PutProperty("EnforceUniqueShortName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPIRATION_PROPERTY.html */
func (r NotesRegistration) Expiration() (domino.Time, error) {
	val, err := r.Com().GetProperty("Expiration")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXPIRATION_PROPERTY.html */
func (r NotesRegistration) SetExpiration(v domino.Time) error {
	return r.Com().PutProperty("Expiration", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPLIST_PROPERTY_REG.html */
func (r NotesRegistration) GroupList() ([]domino.String, error) {
	vals, err := r.Com().GetArrayProperty("GroupList")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GROUPLIST_PROPERTY_REG.html */
func (r NotesRegistration) SetGroupList(v []domino.String) error {
	return r.Com().PutProperty("GroupList", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IDTYPE_PROPERTY.html */
func (r NotesRegistration) IDType() (domino.Long, error) {
	val, err := r.Com().GetProperty("IDType")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IDTYPE_PROPERTY.html */
func (r NotesRegistration) SetIDType(v domino.Long) error {
	return r.Com().PutProperty("IDType", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNORTHAMERICAN_PROPERTY.html */
func (r NotesRegistration) IsNorthAmerican() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IsNorthAmerican")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNORTHAMERICAN_PROPERTY.html */
func (r NotesRegistration) SetIsNorthAmerican(v domino.Boolean) error {
	return r.Com().PutProperty("IsNorthAmerican", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROAMINGUSER_PROPERTY_REG.html */
func (r NotesRegistration) IsRoamingUser() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IsRoamingUser")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROAMINGUSER_PROPERTY_REG.html */
func (r NotesRegistration) SetIsRoamingUser(v domino.Boolean) error {
	return r.Com().PutProperty("IsRoamingUser", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILACLMANAGER_PROPERTY_REG.html */
func (r NotesRegistration) MailACLManager() (domino.String, error) {
	val, err := r.Com().GetProperty("MailACLManager")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILACLMANAGER_PROPERTY_REG.html */
func (r NotesRegistration) SetMailACLManager(v domino.String) error {
	return r.Com().PutProperty("MailACLManager", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILCREATEFTINDEX_PROPERTY_REG.html */
func (r NotesRegistration) MailCreateFTIndex() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("MailCreateFTIndex")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILCREATEFTINDEX_PROPERTY_REG.html */
func (r NotesRegistration) SetMailCreateFTIndex(v domino.Boolean) error {
	return r.Com().PutProperty("MailCreateFTIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILINTERNETADDRESS_PROPERTY_REG.html */
func (r NotesRegistration) MailInternetAddress() (domino.String, error) {
	val, err := r.Com().GetProperty("MailInternetAddress")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILINTERNETADDRESS_PROPERTY_REG.html */
func (r NotesRegistration) SetMailInternetAddress(v domino.String) error {
	return r.Com().PutProperty("MailInternetAddress", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILOWNERACCESS_PROPERTY_REG.html */
func (r NotesRegistration) MailOwnerAccess() (domino.Long, error) {
	val, err := r.Com().GetProperty("MailOwnerAccess")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILOWNERACCESS_PROPERTY_REG.html */
func (r NotesRegistration) SetMailOwnerAccess(v domino.Long) error {
	return r.Com().PutProperty("MailOwnerAccess", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTASIZELIMIT_PROPERTY_REG.html */
func (r NotesRegistration) MailQuotaSizeLimit() (domino.Long, error) {
	val, err := r.Com().GetProperty("MailQuotaSizeLimit")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTASIZELIMIT_PROPERTY_REG.html */
func (r NotesRegistration) SetMailQuotaSizeLimit(v domino.Long) error {
	return r.Com().PutProperty("MailQuotaSizeLimit", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTAWARNINGTHRESHOLD_PROPERTY_REG.html */
func (r NotesRegistration) MailQuotaWarningThreshold() (domino.Long, error) {
	val, err := r.Com().GetProperty("MailQuotaWarningThreshold")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILQUOTAWARNINGTHRESHOLD_PROPERTY_REG.html */
func (r NotesRegistration) SetMailQuotaWarningThreshold(v domino.Long) error {
	return r.Com().PutProperty("MailQuotaWarningThreshold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILREPLICASERVERS_PROPERTY_REG.html */
func (r NotesRegistration) MailReplicaServers() ([]domino.String, error) {
	vals, err := r.Com().GetArrayProperty("MailReplicaServers")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILREPLICASERVERS_PROPERTY_REG.html */
func (r NotesRegistration) SetMailReplicaServers(v []domino.String) error {
	return r.Com().PutProperty("MailReplicaServers", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILSYSTEM_PROPERTY_REG.html */
func (r NotesRegistration) MailSystem() (domino.Long, error) {
	val, err := r.Com().GetProperty("MailSystem")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILSYSTEM_PROPERTY_REG.html */
func (r NotesRegistration) SetMailSystem(v domino.Long) error {
	return r.Com().PutProperty("MailSystem", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILTEMPLATENAME_PROPERTY_REG.html */
func (r NotesRegistration) MailTemplateName() (domino.String, error) {
	val, err := r.Com().GetProperty("MailTemplateName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAILTEMPLATENAME_PROPERTY_REG.html */
func (r NotesRegistration) SetMailTemplateName(v domino.String) error {
	return r.Com().PutProperty("MailTemplateName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MINPASSWORDLENGTH_PROPERTY.html */
func (r NotesRegistration) MinPasswordLength() (domino.Long, error) {
	val, err := r.Com().GetProperty("MinPasswordLength")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MINPASSWORDLENGTH_PROPERTY.html */
func (r NotesRegistration) SetMinPasswordLength(v domino.Long) error {
	return r.Com().PutProperty("MinPasswordLength", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOIDFILE_PROPERTY_REG.html */
func (r NotesRegistration) NoIDFile() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("NoIDFile")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOIDFILE_PROPERTY_REG.html */
func (r NotesRegistration) SetNoIDFile(v domino.Boolean) error {
	return r.Com().PutProperty("NoIDFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT_PROPERTY.html */
func (r NotesRegistration) OrgUnit() (domino.String, error) {
	val, err := r.Com().GetProperty("OrgUnit")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT_PROPERTY.html */
func (r NotesRegistration) SetOrgUnit(v domino.String) error {
	return r.Com().PutProperty("OrgUnit", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POLICYNAME_PROPERTY_REG.html */
func (r NotesRegistration) PolicyName() (domino.String, error) {
	val, err := r.Com().GetProperty("PolicyName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POLICYNAME_PROPERTY_REG.html */
func (r NotesRegistration) SetPolicyName(v domino.String) error {
	return r.Com().PutProperty("PolicyName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONLOG_PROPERTY.html */
func (r NotesRegistration) RegistrationLog() (domino.String, error) {
	val, err := r.Com().GetProperty("RegistrationLog")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONLOG_PROPERTY.html */
func (r NotesRegistration) SetRegistrationLog(v domino.String) error {
	return r.Com().PutProperty("RegistrationLog", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONSERVER_PROPERTY.html */
func (r NotesRegistration) RegistrationServer() (domino.String, error) {
	val, err := r.Com().GetProperty("RegistrationServer")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTRATIONSERVER_PROPERTY.html */
func (r NotesRegistration) SetRegistrationServer(v domino.String) error {
	return r.Com().PutProperty("RegistrationServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPPERIOD_PROPERTY_REG.html */
func (r NotesRegistration) RoamingCleanupPeriod() (domino.Long, error) {
	val, err := r.Com().GetProperty("RoamingCleanupPeriod")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPPERIOD_PROPERTY_REG.html */
func (r NotesRegistration) SetRoamingCleanupPeriod(v domino.Long) error {
	return r.Com().PutProperty("RoamingCleanupPeriod", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPSETTING_PROPERTY_REG.html */
func (r NotesRegistration) RoamingCleanupSetting() (domino.Long, error) {
	val, err := r.Com().GetProperty("RoamingCleanupSetting")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGCLEANUPSETTING_PROPERTY_REG.html */
func (r NotesRegistration) SetRoamingCleanupSetting(v domino.Long) error {
	return r.Com().PutProperty("RoamingCleanupSetting", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSERVER_PROPERTY_REG.html */
func (r NotesRegistration) RoamingServer() (domino.String, error) {
	val, err := r.Com().GetProperty("RoamingServer")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSERVER_PROPERTY_REG.html */
func (r NotesRegistration) SetRoamingServer(v domino.String) error {
	return r.Com().PutProperty("RoamingServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSUBDIR_PROPERTY_REG.html */
func (r NotesRegistration) RoamingSubdir() (domino.String, error) {
	val, err := r.Com().GetProperty("RoamingSubdir")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROAMINGSUBDIR_PROPERTY_REG.html */
func (r NotesRegistration) SetRoamingSubdir(v domino.String) error {
	return r.Com().PutProperty("RoamingSubdir", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SHORTNAME_PROPERTY_REG.html */
func (r NotesRegistration) ShortName() (domino.String, error) {
	val, err := r.Com().GetProperty("ShortName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SHORTNAME_PROPERTY_REG.html */
func (r NotesRegistration) SetShortName(v domino.String) error {
	return r.Com().PutProperty("ShortName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDTOADDRESSBOOK_PROPERTY.html */
func (r NotesRegistration) StoreIDInAddressBook() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("StoreIDInAddressBook")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDTOADDRESSBOOK_PROPERTY.html */
func (r NotesRegistration) SetStoreIDInAddressBook(v domino.Boolean) error {
	return r.Com().PutProperty("StoreIDInAddressBook", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDINMAILFILE_PROPERTY_REG.html */
func (r NotesRegistration) StoreIDInMailfile() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("StoreIDInMailfile")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STOREIDINMAILFILE_PROPERTY_REG.html */
func (r NotesRegistration) SetStoreIDInMailfile(v domino.Boolean) error {
	return r.Com().PutProperty("StoreIDInMailfile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SYNCHINTERNETPASSWORD_PROPERTY_REG.html */
func (r NotesRegistration) SynchInternetPassword() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("SynchInternetPassword")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SYNCHINTERNETPASSWORD_PROPERTY_REG.html */
func (r NotesRegistration) SetSynchInternetPassword(v domino.Boolean) error {
	return r.Com().PutProperty("SynchInternetPassword", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEADDRESSBOOK_PROPERTY.html */
func (r NotesRegistration) UpdateAddressBook() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("UpdateAddressBook")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEADDRESSBOOK_PROPERTY.html */
func (r NotesRegistration) SetUpdateAddressBook(v domino.Boolean) error {
	return r.Com().PutProperty("UpdateAddressBook", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_REG.html */
func (r NotesRegistration) UseCertificateAuthority() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("UseCertificateAuthority")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USECERTIFICATEAUTHORITY_PROPERTY_REG.html */
func (r NotesRegistration) SetUseCertificateAuthority(v domino.Boolean) error {
	return r.Com().PutProperty("UseCertificateAuthority", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCERTIFIERTOADDRESSBOOK_METHOD.html */
func (r NotesRegistration) AddCertifierToAddressBook(idfile domino.String, certpw domino.String, location domino.String, comment domino.String) error {
	_, err := r.Com().CallMethod("AddCertifierToAddressBook", idfile, certpw, location, comment)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDSERVERTOADDRESSBOOK_METHOD.html */
func (r NotesRegistration) AddServerToAddressBook(idfile domino.String, server domino.String, domain domino.String, userpw domino.String, network domino.String, adminname domino.String, title domino.String, location domino.String, comment domino.String) error {
	_, err := r.Com().CallMethod("AddServerToAddressBook", idfile, server, domain, userpw, network, adminname, title, location, comment)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDUSERPROFILE_METHOD.html */
func (r NotesRegistration) AddUserProfile(username domino.String, profilename domino.String) error {
	_, err := r.Com().CallMethod("AddUserProfile", username, profilename)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDUSERTOADDRESSBOOK_METHOD.html */
func (r NotesRegistration) AddUserToAddressBook(idfile domino.String, fullname domino.String, lastname domino.String, userpw domino.String, firstname domino.String, middle domino.String, mailserver domino.String, mailidpath domino.String, fwdaddress domino.String, location domino.String, comment domino.String) error {
	_, err := r.Com().CallMethod("AddUserToAddressBook", idfile, fullname, lastname, userpw, firstname, middle, mailserver, mailidpath, fwdaddress, location, comment)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CROSSCERTIFY_METHOD.html */
func (r NotesRegistration) CrossCertify(idfile domino.String, certpw domino.String, comment domino.String) error {
	_, err := r.Com().CallMethod("CrossCertify", idfile, certpw, comment)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEIDONSERVER_METHOD.html */
func (r NotesRegistration) DeleteIDOnServer(username domino.String, isserverid domino.Boolean) error {
	_, err := r.Com().CallMethod("DeleteIDOnServer", username, isserverid)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETIDFROMSERVER_METHOD.html */
func (r NotesRegistration) GetIDFromServer(username domino.String, filepath domino.String, isserverid domino.Boolean) error {
	_, err := r.Com().CallMethod("GetIDFromServer", username, filepath, isserverid)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERINFO_METHOD.html */
func (r NotesRegistration) GetUserInfo(username domino.String, retMailServer domino.String, retMailFile domino.String, retMailDomain domino.String, retMailSystem domino.Integer, retProfile domino.String) error {
	_, err := r.Com().CallMethod("GetUserInfo", username, retMailServer, retMailFile, retMailDomain, retMailSystem, retProfile)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RECERTIFY_METHOD.html */
func (r NotesRegistration) Recertify(idfile domino.String, certpw domino.String, comment domino.String) error {
	_, err := r.Com().CallMethod("Recertify", idfile, certpw, comment)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWCERTIFIER_METHOD.html */
func (r NotesRegistration) RegisterNewCertifier(organization domino.String, idfile domino.String, certpw domino.String, country domino.String) error {
	_, err := r.Com().CallMethod("RegisterNewCertifier", organization, idfile, certpw, country)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWSERVER_METHOD.html */
func (r NotesRegistration) RegisterNewServer(server domino.String, idfile domino.String, domain domino.String, servpw domino.String, certpw domino.String, location domino.String, comment domino.String, network domino.String, adminname domino.String, title domino.String) error {
	_, err := r.Com().CallMethod("RegisterNewServer", server, idfile, domain, servpw, certpw, location, comment, network, adminname, title)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REGISTERNEWUSER_METHOD.html */
func (r NotesRegistration) RegisterNewUser(lastname domino.String, idfile domino.String, mailserver domino.String, firstname domino.String, middle domino.String, certpw domino.String, location domino.String, comment domino.String, maildbpath domino.String, fwddomain domino.String, userpw domino.String, usertype domino.Integer, altname domino.String, altnamelang domino.String) error {
	_, err := r.Com().CallMethod("RegisterNewUser", lastname, idfile, mailserver, firstname, middle, certpw, location, comment, maildbpath, fwddomain, userpw, usertype, altname, altnamelang)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SWITCHTOID_METHOD.html */
func (r NotesRegistration) SwitchToID(idfile domino.String, userpw domino.String) error {
	_, err := r.Com().CallMethod("SwitchToID", idfile, userpw)
	return err
}
