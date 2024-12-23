/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACLENTRY_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesACLEntry struct {
	NotesStruct
}

func NewNotesACLEntry(dispatchPtr *ole.IDispatch) NotesACLEntry {
	return NotesACLEntry{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORNOCREATE_PROPERTY.html */
func (a NotesACLEntry) CanCreateDocuments() (Boolean, error) {
	val, err := getComProperty(a, "CanCreateDocuments")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORNOCREATE_PROPERTY.html */
func (a NotesACLEntry) SetCanCreateDocuments(v Boolean) error {
	return putComProperty(a, "CanCreateDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATELSORJAVAAGENT_PROPERTY_6760.html */
func (a NotesACLEntry) CanCreateLSOrJavaAgent() (Boolean, error) {
	val, err := getComProperty(a, "CanCreateLSOrJavaAgent")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATELSORJAVAAGENT_PROPERTY_6760.html */
func (a NotesACLEntry) SetCanCreateLSOrJavaAgent(v Boolean) error {
	return putComProperty(a, "CanCreateLSOrJavaAgent", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEAGENT_PROPERTY.html */
func (a NotesACLEntry) CanCreatePersonalAgent() (Boolean, error) {
	val, err := getComProperty(a, "CanCreatePersonalAgent")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEAGENT_PROPERTY.html */
func (a NotesACLEntry) SetCanCreatePersonalAgent(v Boolean) error {
	return putComProperty(a, "CanCreatePersonalAgent", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEFOLDER_PROPERTY.html */
func (a NotesACLEntry) CanCreatePersonalFolder() (Boolean, error) {
	val, err := getComProperty(a, "CanCreatePersonalFolder")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPRIVATEFOLDER_PROPERTY.html */
func (a NotesACLEntry) SetCanCreatePersonalFolder(v Boolean) error {
	return putComProperty(a, "CanCreatePersonalFolder", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATESHAREDFOLDER_PROPERTY_7622.html */
func (a NotesACLEntry) CanCreateSharedFolder() (Boolean, error) {
	val, err := getComProperty(a, "CanCreateSharedFolder")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCREATESHAREDFOLDER_PROPERTY_7622.html */
func (a NotesACLEntry) SetCanCreateSharedFolder(v Boolean) error {
	return putComProperty(a, "CanCreateSharedFolder", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNODELETE_PROPERTY.html */
func (a NotesACLEntry) CanDeleteDocuments() (Boolean, error) {
	val, err := getComProperty(a, "CanDeleteDocuments")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNODELETE_PROPERTY.html */
func (a NotesACLEntry) SetCanDeleteDocuments(v Boolean) error {
	return putComProperty(a, "CanDeleteDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANREPLICATEORCOPYDOCUMENTS_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) CanReplicateOrCopyDocuments() (Boolean, error) {
	val, err := getComProperty(a, "CanReplicateOrCopyDocuments")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANREPLICATEORCOPYDOCUMENTS_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) SetCanReplicateOrCopyDocuments(v Boolean) error {
	return putComProperty(a, "CanReplicateOrCopyDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_88.html */
func (a NotesACLEntry) IsAdminReaderAuthor() (Boolean, error) {
	val, err := getComProperty(a, "IsAdminReaderAuthor")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_88.html */
func (a NotesACLEntry) SetIsAdminReaderAuthor(v Boolean) error {
	return putComProperty(a, "IsAdminReaderAuthor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINSERVER_PROPERTY_2946.html */
func (a NotesACLEntry) IsAdminServer() (Boolean, error) {
	val, err := getComProperty(a, "IsAdminServer")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINSERVER_PROPERTY_2946.html */
func (a NotesACLEntry) SetIsAdminServer(v Boolean) error {
	return putComProperty(a, "IsAdminServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISGROUP_PROPERTY_4731.html */
func (a NotesACLEntry) IsGroup() (Boolean, error) {
	val, err := getComProperty(a, "IsGroup")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISGROUP_PROPERTY_4731.html */
func (a NotesACLEntry) SetIsGroup(v Boolean) error {
	return putComProperty(a, "IsGroup", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPERSON_PROPERTY_651.html */
func (a NotesACLEntry) IsPerson() (Boolean, error) {
	val, err := getComProperty(a, "IsPerson")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPERSON_PROPERTY_651.html */
func (a NotesACLEntry) SetIsPerson(v Boolean) error {
	return putComProperty(a, "IsPerson", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICREADER_PROPERTY.html */
func (a NotesACLEntry) IsPublicReader() (Boolean, error) {
	val, err := getComProperty(a, "IsPublicReader")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICREADER_PROPERTY.html */
func (a NotesACLEntry) SetIsPublicReader(v Boolean) error {
	return putComProperty(a, "IsPublicReader", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICWRITER_PROPERTY.html */
func (a NotesACLEntry) IsPublicWriter() (Boolean, error) {
	val, err := getComProperty(a, "IsPublicWriter")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLICWRITER_PROPERTY.html */
func (a NotesACLEntry) SetIsPublicWriter(v Boolean) error {
	return putComProperty(a, "IsPublicWriter", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSERVER_PROPERTY_1888.html */
func (a NotesACLEntry) IsServer() (Boolean, error) {
	val, err := getComProperty(a, "IsServer")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSERVER_PROPERTY_1888.html */
func (a NotesACLEntry) SetIsServer(v Boolean) error {
	return putComProperty(a, "IsServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY.html */
func (a NotesACLEntry) Level() (Long, error) {
	val, err := getComProperty(a, "Level")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY.html */
func (a NotesACLEntry) SetLevel(v Long) error {
	return putComProperty(a, "Level", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) Name() (String, error) {
	val, err := getComProperty(a, "Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) SetName(v String) error {
	return putComProperty(a, "Name", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEOBJECT_PROPERTY.html */
func (a NotesACLEntry) NameObject() (NotesName, error) {
	dispatchPtr, err := getComObjectProperty(a, "NameObject")
	return NewNotesName(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) Parent() (NotesACL, error) {
	dispatchPtr, err := getComObjectProperty(a, "Parent")
	return NewNotesACL(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROLES_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) Roles() ([]String, error) {
	vals, err := getComArrayProperty(a, "Roles")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERTYPE_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) UserType() (Long, error) {
	val, err := getComProperty(a, "UserType")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERTYPE_PROPERTY_ACLENTRY.html */
func (a NotesACLEntry) SetUserType(v Long) error {
	return putComProperty(a, "UserType", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLEROLE_METHOD.html */
func (a NotesACLEntry) DisableRole(name String) error {
	_, err := callComMethod(a, "DisableRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENABLEROLE_METHOD.html */
func (a NotesACLEntry) EnableRole(name String) error {
	_, err := callComMethod(a, "EnableRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISROLEENABLED_METHOD.html */
func (a NotesACLEntry) IsRoleEnabled(name String) (Boolean, error) {
	val, err := callComMethod(a, "IsRoleEnabled", name)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_ACLENTRY.html */
func (a NotesACLEntry) Remove() error {
	_, err := callComMethod(a, "Remove")
	return err
}
