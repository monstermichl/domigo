/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACL_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesACL struct {
	NotesStruct
}

func newNotesACL(dispatchPtr *ole.IDispatch) NotesACL {
	return NotesACL{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func (a NotesACL) AdministrationServer() (String, error) {
	return getComProperty[String](a, "AdministrationServer")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func (a NotesACL) SetAdministrationServer(v String) error {
	return putComProperty(a, "AdministrationServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func (a NotesACL) InternetLevel() (Long, error) {
	return getComProperty[Long](a, "InternetLevel")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func (a NotesACL) SetInternetLevel(v Long) error {
	return putComProperty(a, "InternetLevel", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func (a NotesACL) IsAdminNames() (Boolean, error) {
	return getComProperty[Boolean](a, "IsAdminNames")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func (a NotesACL) SetIsAdminNames(v Boolean) error {
	return putComProperty(a, "IsAdminNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func (a NotesACL) IsAdminReaderAuthor() (Boolean, error) {
	return getComProperty[Boolean](a, "IsAdminReaderAuthor")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func (a NotesACL) SetIsAdminReaderAuthor(v Boolean) error {
	return putComProperty(a, "IsAdminReaderAuthor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func (a NotesACL) IsExtendedAccess() (Boolean, error) {
	return getComProperty[Boolean](a, "IsExtendedAccess")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func (a NotesACL) SetIsExtendedAccess(v Boolean) error {
	return putComProperty(a, "IsExtendedAccess", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ACL.html */
func (a NotesACL) Parent() (NotesDatabase, error) {
	return getComObjectProperty(a, newNotesDatabase, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROLES_PROPERTY_ACL.html */
func (a NotesACL) Roles() ([]String, error) {
	return getComArrayProperty[String](a, "Roles")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func (a NotesACL) UniformAccess() (Boolean, error) {
	return getComProperty[Boolean](a, "UniformAccess")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func (a NotesACL) SetUniformAccess(v Boolean) error {
	return putComProperty(a, "UniformAccess", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDROLE_METHOD.html */
func (a NotesACL) AddRole(name String) error {
	return callComVoidMethod(a, "AddRole", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEACLENTRY_METHOD.html */
func (a NotesACL) CreateACLEntry(name String, level Long) (NotesACLEntry, error) {
	return callComObjectMethod(a, newNotesACLEntry, "CreateACLEntry", name, level)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEROLE_METHOD.html */
func (a NotesACL) DeleteRole(name String) error {
	return callComVoidMethod(a, "DeleteRole", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD.html */
func (a NotesACL) GetEntry(name String) (NotesACLEntry, error) {
	return callComObjectMethod(a, newNotesACLEntry, "GetEntry", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTENTRY_METHOD.html */
func (a NotesACL) GetFirstEntry() (NotesACLEntry, error) {
	return callComObjectMethod(a, newNotesACLEntry, "GetFirstEntry")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_METHOD.html */
func (a NotesACL) GetNextEntry(entry NotesACLEntry) (NotesACLEntry, error) {
	return callComObjectMethod(a, newNotesACLEntry, "GetNextEntry", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEACLENTRY_METHOD_ACL_COM.html */
func (a NotesACL) RemoveACLEntry(name String) error {
	return callComVoidMethod(a, "RemoveACLEntry", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEROLE_METHOD.html */
func (a NotesACL) RenameRole(oldName String, newName String) error {
	return callComVoidMethod(a, "RenameRole", oldName, newName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_ACL.html */
func (a NotesACL) Save() error {
	return callComVoidMethod(a, "Save")
}
