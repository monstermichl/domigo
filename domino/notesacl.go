/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACL_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesACL struct {
	NotesStruct
}

func NewNotesACL(dispatchPtr *ole.IDispatch) NotesACL {
	return NotesACL{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func (a NotesACL) AdministrationServer() (String, error) {
	val, err := a.Com().GetProperty("AdministrationServer")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func (a NotesACL) SetAdministrationServer(v String) error {
	return a.Com().PutProperty("AdministrationServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func (a NotesACL) InternetLevel() (Long, error) {
	val, err := a.Com().GetProperty("InternetLevel")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func (a NotesACL) SetInternetLevel(v Long) error {
	return a.Com().PutProperty("InternetLevel", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func (a NotesACL) IsAdminNames() (Boolean, error) {
	val, err := a.Com().GetProperty("IsAdminNames")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func (a NotesACL) SetIsAdminNames(v Boolean) error {
	return a.Com().PutProperty("IsAdminNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func (a NotesACL) IsAdminReaderAuthor() (Boolean, error) {
	val, err := a.Com().GetProperty("IsAdminReaderAuthor")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func (a NotesACL) SetIsAdminReaderAuthor(v Boolean) error {
	return a.Com().PutProperty("IsAdminReaderAuthor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func (a NotesACL) IsExtendedAccess() (Boolean, error) {
	val, err := a.Com().GetProperty("IsExtendedAccess")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func (a NotesACL) SetIsExtendedAccess(v Boolean) error {
	return a.Com().PutProperty("IsExtendedAccess", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ACL.html */
func (a NotesACL) Parent() (NotesDatabase, error) {
	dispatchPtr, err := a.Com().GetObjectProperty("Parent")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROLES_PROPERTY_ACL.html */
func (a NotesACL) Roles() ([]String, error) {
	vals, err := a.Com().GetArrayProperty("Roles")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func (a NotesACL) UniformAccess() (Boolean, error) {
	val, err := a.Com().GetProperty("UniformAccess")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func (a NotesACL) SetUniformAccess(v Boolean) error {
	return a.Com().PutProperty("UniformAccess", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDROLE_METHOD.html */
func (a NotesACL) AddRole(name String) error {
	_, err := a.Com().CallMethod("AddRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEACLENTRY_METHOD.html */
func (a NotesACL) CreateACLEntry(name String, level Long) (NotesACLEntry, error) {
	dispatchPtr, err := a.Com().CallObjectMethod("CreateACLEntry", name, level)
	return NewNotesACLEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEROLE_METHOD.html */
func (a NotesACL) DeleteRole(name String) error {
	_, err := a.Com().CallMethod("DeleteRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD.html */
func (a NotesACL) GetEntry(name String) (NotesACLEntry, error) {
	dispatchPtr, err := a.Com().CallObjectMethod("GetEntry", name)
	return NewNotesACLEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTENTRY_METHOD.html */
func (a NotesACL) GetFirstEntry() (NotesACLEntry, error) {
	dispatchPtr, err := a.Com().CallObjectMethod("GetFirstEntry")
	return NewNotesACLEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_METHOD.html */
func (a NotesACL) GetNextEntry(entry NotesACLEntry) (NotesACLEntry, error) {
	dispatchPtr, err := a.Com().CallObjectMethod("GetNextEntry", entry.Com().Dispatch())
	return NewNotesACLEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEACLENTRY_METHOD_ACL_COM.html */
func (a NotesACL) RemoveACLEntry(name String) error {
	_, err := a.Com().CallMethod("RemoveACLEntry", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEROLE_METHOD.html */
func (a NotesACL) RenameRole(oldName String, newName String) error {
	_, err := a.Com().CallMethod("RenameRole", oldName, newName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_ACL.html */
func (a NotesACL) Save() error {
	_, err := a.Com().CallMethod("Save")
	return err
}
