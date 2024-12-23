/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACL_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
	"github.com/monstermichl/domigo/internal/helpers"
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
	val, err := getComProperty(a, "AdministrationServer")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func (a NotesACL) SetAdministrationServer(v String) error {
	return putComProperty(a, "AdministrationServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func (a NotesACL) InternetLevel() (Long, error) {
	val, err := getComProperty(a, "InternetLevel")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func (a NotesACL) SetInternetLevel(v Long) error {
	return putComProperty(a, "InternetLevel", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func (a NotesACL) IsAdminNames() (Boolean, error) {
	val, err := getComProperty(a, "IsAdminNames")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func (a NotesACL) SetIsAdminNames(v Boolean) error {
	return putComProperty(a, "IsAdminNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func (a NotesACL) IsAdminReaderAuthor() (Boolean, error) {
	val, err := getComProperty(a, "IsAdminReaderAuthor")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func (a NotesACL) SetIsAdminReaderAuthor(v Boolean) error {
	return putComProperty(a, "IsAdminReaderAuthor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func (a NotesACL) IsExtendedAccess() (Boolean, error) {
	val, err := getComProperty(a, "IsExtendedAccess")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func (a NotesACL) SetIsExtendedAccess(v Boolean) error {
	return putComProperty(a, "IsExtendedAccess", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ACL.html */
func (a NotesACL) Parent() (NotesDatabase, error) {
	dispatchPtr, err := getComObjectProperty(a, "Parent")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROLES_PROPERTY_ACL.html */
func (a NotesACL) Roles() ([]String, error) {
	vals, err := getComArrayProperty(a, "Roles")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func (a NotesACL) UniformAccess() (Boolean, error) {
	val, err := getComProperty(a, "UniformAccess")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func (a NotesACL) SetUniformAccess(v Boolean) error {
	return putComProperty(a, "UniformAccess", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDROLE_METHOD.html */
func (a NotesACL) AddRole(name String) error {
	_, err := callComMethod(a, "AddRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEACLENTRY_METHOD.html */
func (a NotesACL) CreateACLEntry(name String, level Long) (NotesACLEntry, error) {
	dispatchPtr, err := callComObjectMethod(a, "CreateACLEntry", name, level)
	return NewNotesACLEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEROLE_METHOD.html */
func (a NotesACL) DeleteRole(name String) error {
	_, err := callComMethod(a, "DeleteRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD.html */
func (a NotesACL) GetEntry(name String) (NotesACLEntry, error) {
	dispatchPtr, err := callComObjectMethod(a, "GetEntry", name)
	return NewNotesACLEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTENTRY_METHOD.html */
func (a NotesACL) GetFirstEntry() (NotesACLEntry, error) {
	dispatchPtr, err := callComObjectMethod(a, "GetFirstEntry")
	return NewNotesACLEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_METHOD.html */
func (a NotesACL) GetNextEntry(entry NotesACLEntry) (NotesACLEntry, error) {
	dispatchPtr, err := callComObjectMethod(a, "GetNextEntry", entry.com().Dispatch())
	return NewNotesACLEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEACLENTRY_METHOD_ACL_COM.html */
func (a NotesACL) RemoveACLEntry(name String) error {
	_, err := callComMethod(a, "RemoveACLEntry", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEROLE_METHOD.html */
func (a NotesACL) RenameRole(oldName String, newName String) error {
	_, err := callComMethod(a, "RenameRole", oldName, newName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_ACL.html */
func (a NotesACL) Save() error {
	_, err := callComMethod(a, "Save")
	return err
}
